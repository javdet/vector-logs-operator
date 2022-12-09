/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"github.com/go-logr/logr"
	loggerv1beta "github.com/javdet/vector-logs-operator/api/v1beta"
	"github.com/redhat-cop/operator-utils/pkg/util"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

type AggregatorReconciler struct {
	// client.Client
	Log logr.Logger
	util.ReconcilerBase
}

var controllerAggregatorLog = ctrl.Log.WithName("controller").WithName("VectorAggregator")

// +kubebuilder:rbac:groups=logging.vlo.io,resources=vectoraggregators,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=logging.vlo.io,resources=vectoraggregators/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=logging.vlo.io,resources=vectoraggregators/finalizers,verbs=update
// +kubebuilder:rbac:groups=apps,resources=statefulsets,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the M2Logstash object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.7.0/pkg/reconcile

func (r *AggregatorReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	controllerAggregatorLog.Info("start reconcile", "request", req.NamespacedName)

	instance := &loggerv1beta.VectorAggregator{}

	if req.NamespacedName.Namespace == "" {
		controllerAgentPipelineLog.Info("Namespace event", "request", req.NamespacedName)
		namespace := &corev1.Namespace{}
		if err := r.GetClient().Get(ctx, req.NamespacedName, namespace); err != nil {
			if apierrors.IsNotFound(err) {
				return ctrl.Result{}, nil
			}
			controllerAggregatorLog.Error(err, "Cannot get object Namespace")
			return ctrl.Result{}, err
		}

		instanceList := &loggerv1beta.VectorAggregatorList{}
		err := r.GetClient().List(ctx, instanceList)
		if err != nil {
			return ctrl.Result{}, nil
		}
		if len(instanceList.Items) > 0 {
			instance = &instanceList.Items[0]
		} else {
			controllerAggregatorLog.Info("cannot get object AgentPipeline")
			return ctrl.Result{}, nil
		}
	} else {
		controllerAggregatorLog.Info("Not Namespace event")
		if err := r.GetClient().Get(ctx, req.NamespacedName, instance); err != nil {
			if apierrors.IsNotFound(err) {
				return ctrl.Result{}, nil
			}
			controllerAggregatorLog.Error(err, "cannot get object AgentPipeline")
			return ctrl.Result{}, err
		}
	}

	nameSpaceList := &corev1.NamespaceList{}
	opts := []client.ListOption{
		client.MatchingLabels{nsLabel: "true"},
	}
	if err := r.GetClient().List(ctx, nameSpaceList, opts...); err != nil {
		return r.ManageError(ctx, instance, err)
	}
	for _, item := range nameSpaceList.Items {
		err, response := r.syncPipelineResources(ctx, *instance, item.Name)
		if err != nil {
			controllerAgentPipelineLog.Error(err, response)
			return r.ManageError(ctx, instance, err)
		}
	}

	controllerAggregatorLog.Info("finish reconcile", "request", req.NamespacedName)
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *AggregatorReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&loggerv1beta.VectorAggregator{}).
		Watches(
			&source.Kind{Type: &corev1.Namespace{}},
			&handler.EnqueueRequestForObject{},
		).
		WithEventFilter(namespaceFilter()).
		WithOptions(controller.Options{MaxConcurrentReconciles: 2}).
		Complete(r)
}

func (r *AggregatorReconciler) syncPipelineResources(
	ctx context.Context, instance loggerv1beta.VectorAggregator, namespace string) (error, string) {
	controllerAgentPipelineLog.Info("Sync pipeline resources", "instance", instance.GetName())

	err := r.CreateOrUpdateResource(
		ctx,
		&instance,
		instance.Namespace,
		r.PipelineConfigMapFromCR(&instance, namespace),
	)
	if err != nil {
		controllerLog.Error(err, "cannot create configmap")
		return err, ""
	}

	err = r.CreateResourceIfNotExists(ctx, &instance, instance.Namespace, r.serviceAccountFromCR(&instance, namespace))
	if err != nil {
		controllerLog.Error(err, "cannot create serviceaccount")
		return err, ""
	}

	err = r.CreateOrUpdateResource(ctx, &instance, instance.Namespace, r.clusterRoleFromCR(&instance, namespace))
	if err != nil {
		controllerLog.Error(err, "cannot create role")
		return err, ""
	}

	err = r.CreateOrUpdateResource(ctx, &instance, instance.Namespace, r.clusterRoleBindingFromCR(&instance, namespace))
	if err != nil {
		controllerLog.Error(err, "cannot create rolebinding")
		return err, ""
	}

	err = r.CreateOrUpdateResource(ctx, &instance, instance.Namespace, r.statefulSetFromCR(&instance, namespace))
	if err != nil {
		controllerLog.Error(err, "cannot create daemonset")
		return err, ""
	}

	err = r.CreateResourceIfNotExists(ctx, &instance, instance.Namespace, r.statefulSetServiceFromCR(&instance, namespace))
	if err != nil {
		controllerLog.Error(err, "cannot create configmap")
		return err, ""

	}

	return nil, "Success"
}
