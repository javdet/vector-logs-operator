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
)

type AgentReconciler struct {
	// client.Client
	Log logr.Logger
	util.ReconcilerBase
}

var controllerLog = ctrl.Log.WithName("controller").WithName("VectorAgent")

// +kubebuilder:rbac:groups=logging.vlo.io,resources=vectoragents,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=logging.vlo.io,resources=vectoragents/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=logging.vlo.io,resources=vectoragents/finalizers,verbs=update
// +kubebuilder:rbac:groups=core,resources=namespaces,verbs=get;list;watch
// +kubebuilder:rbac:groups=apps,resources=daemonsets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=nodes,verbs=list;watch
// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=roles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=rolebindings,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=extensions,resources=podsecuritypolicies,verbs=use;
// +kubebuilder:rbac:groups=core,resources=events,verbs=get;list;watch;create;update;delete;patch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the M2Logstash object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.7.0/pkg/reconcile

func (r *AgentReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithName("controller").WithName("agent")
	controllerLog.Info("start reconcile", "request", req.NamespacedName)

	instance := &loggerv1beta.VectorAgent{}
	if err := r.GetClient().Get(ctx, req.NamespacedName, instance); err != nil {
		if apierrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		log.Error(err, "cannot get object VectorAgent")
		return ctrl.Result{}, err
	}

	nameSpaceList := &corev1.NamespaceList{}
	opts := []client.ListOption{
		client.MatchingLabels{nsLabel: "true"},
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
func (r *AgentReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&loggerv1beta.VectorAgent{}).
		WithOptions(controller.Options{MaxConcurrentReconciles: 2}).
		Complete(r)
}

func (r *AgentReconciler) syncPipelineResources(
	ctx context.Context, instance loggerv1beta.VectorAgent, namespaces []string) (error, string) {
	controllerAgentPipelineLog.Info("Sync pipeline resources", "instance", instance.GetName())
	agentPipelineList := &loggerv1beta.VectorAgentPipelineList{}
	var agentPipeline = &loggerv1beta.VectorAgentPipeline{}

	if err := r.GetClient().List(ctx, agentPipelineList); err != nil {
		controllerAgentPipelineLog.Error(err, "Failed to query vectoragentpipeline")
		return err, ""
	}

	for _, pipeline := range agentPipelineList.Items {
		for key, value := range pipeline.Spec.Selector {
			if val, ok := instance.Labels[key]; ok {
				if value == val {
					agentPipeline = &pipeline
					break
				}
			}
		}
	}
	err := r.CreateOrUpdateResource(
		ctx,
		&instance,
		instance.Namespace,
		r.PipelineConfigMapFromCR(&instance, agentPipeline, namespaces),
	)
	if err != nil {
		controllerLog.Error(err, "cannot create configmap")
		return err, ""
	}

	err = r.CreateResourceIfNotExists(ctx, &instance, instance.Namespace, r.serviceAccountFromCR(&instance))
	if err != nil {
		controllerLog.Error(err, "cannot create serviceaccount")
		return err, ""
	}

	err = r.CreateOrUpdateResource(ctx, &instance, instance.Namespace, r.clusterRoleFromCR(&instance))
	if err != nil {
		controllerLog.Error(err, "cannot create role")
		return err, ""
	}

	err = r.CreateOrUpdateResource(ctx, &instance, instance.Namespace, r.clusterRoleBindingFromCR(&instance))
	if err != nil {
		controllerLog.Error(err, "cannot create rolebinding")
		return err, ""
	}

	err = r.CreateOrUpdateResource(ctx, &instance, instance.Namespace, r.daemonSetFromCR(&instance))
	if err != nil {
		controllerLog.Error(err, "cannot create daemonset")
		return err, ""
	}

	err = r.CreateResourceIfNotExists(ctx, &instance, instance.Namespace, r.daemonSetServiceFromCR(&instance))
	if err != nil {
		controllerLog.Error(err, "cannot create configmap")
		return err, ""

	}

	return nil, "Success"
}
