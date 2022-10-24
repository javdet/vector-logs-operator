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

type AgentPipelineReconciler struct {
	// client.Client
	Log logr.Logger
	util.ReconcilerBase
}

var controllerAgentPipelineLog = ctrl.Log.WithName("controller").WithName("VectorAgentPipeline")

// +kubebuilder:rbac:groups=logging.vlo.io,resources=vectoragentpipelines,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=logging.vlo.io,resources=vectoragentpipelines/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=logging.vlo.io,resources=vectoragentpipelines/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the M2Logstash object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.7.0/pkg/reconcile

func (r *AgentPipelineReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	controllerAgentPipelineLog.Info("start reconcile", "request", req.NamespacedName)

	instance := &loggerv1beta.VectorAgentPipeline{}
	if req.NamespacedName.Namespace == "" {
		namespace := &corev1.Namespace{}
		if err := r.GetClient().Get(ctx, req.NamespacedName, namespace); err != nil {
			if apierrors.IsNotFound(err) {
				return ctrl.Result{}, nil
			}
			controllerLog.Error(err, "Cannot get object Namespace")
			return ctrl.Result{}, err
		}

		instanceList := &loggerv1beta.VectorAgentPipelineList{}
		err := r.GetClient().List(ctx, instanceList)
		if err != nil {
			return ctrl.Result{}, nil
		}
		instance = &instanceList.Items[0]

	} else {
		if err := r.GetClient().Get(ctx, req.NamespacedName, instance); err != nil {
			if apierrors.IsNotFound(err) {
				return ctrl.Result{}, nil
			}
			controllerAgentPipelineLog.Error(err, "cannot get object M2LogstashPipeline")
			return ctrl.Result{}, err
		}
	}

	nameSpaceList := &corev1.NamespaceList{}
	opts := []client.ListOption{
		client.MatchingLabels{"vlo.io/logs": "true"},
	}
	if err := r.GetClient().List(ctx, nameSpaceList, opts...); err != nil {
		return r.ManageError(ctx, instance, err)
	}
	var namespaces []string
	for _, item := range nameSpaceList.Items {
		namespaces = append(namespaces, item.Name)
	}
	err, response := r.syncPipelineResources(ctx, *instance, namespaces)
	if err != nil {
		controllerAgentPipelineLog.Error(err, response)
		return r.ManageError(ctx, instance, err)
	}

	controllerLog.Info("finish reconcile", "request", req.NamespacedName)
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *AgentPipelineReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&loggerv1beta.VectorAgentPipeline{}).
		Watches(
			&source.Kind{Type: &corev1.Namespace{}},
			&handler.EnqueueRequestForObject{},
		).
		WithEventFilter(namespaceFilter()).
		WithOptions(controller.Options{MaxConcurrentReconciles: 2}).
		Complete(r)
}

func (r *AgentPipelineReconciler) syncPipelineResources(
	ctx context.Context, instance loggerv1beta.VectorAgentPipeline, namespaces []string) (error, string) {
	controllerAgentPipelineLog.Info("Sync instance resources", "instance", instance.GetName())
	agentList := &loggerv1beta.VectorAgentList{}

	var matchingLabels client.MatchingLabels
	matchingLabels = instance.Spec.Selector

	opts := []client.ListOption{
		matchingLabels,
	}
	if err := r.GetClient().List(ctx, agentList, opts...); err != nil {
		controllerAgentPipelineLog.Error(err, "Failed to query vectoragent")
		return err, ""
	}

	for _, agent := range agentList.Items {
		err := r.CreateOrUpdateResource(
			ctx,
			&instance,
			agent.Namespace,
			r.PipelineConfigMapFromCR(&instance, agent.Name, agent.Namespace, namespaces),
		)
		if err != nil {
			return err, "Cannot update pipeline "
		}

		err = r.CreateOrUpdateResource(
			ctx,
			&instance,
			agent.Namespace,
			r.daemonSetFromCR(&instance, &agent),
		)
	}

	return nil, "Success"
}
