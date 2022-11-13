package controllers

import (
	loggerv1beta "github.com/javdet/vector-logs-operator/api/v1beta"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type VectorAgentPipeline struct {
	Sources    PipelineSources
	Transforms PipelineTransforms
	Sinks      PipelineSinks
	CRD        loggerv1beta.VectorAgentPipelineSpec
}

type PipelineSources struct {
	Metrics    PipelineSourcesMetrics
	Kubernetes PipelineSourcesKubernetes
}

type PipelineSourcesMetrics struct {
	Namespace string
}

type PipelineTransforms struct {
	Filter PipelineTransformsFilter
}

type PipelineTransformsFilter struct {
	Namespaces []string
}

type PipelineSinks struct {
	Prometheus PipelineSinksPrometheus
}

type PipelineSinksPrometheus struct {
	Namespace string
}

type PipelineSourcesKubernetes struct {
	PodAnnotations map[string]string
}

func (r *AgentReconciler) PipelineConfigMapFromCR(
	instance *loggerv1beta.VectorAgent, agentPipeline *loggerv1beta.VectorAgentPipeline, namespaces []string) *corev1.ConfigMap {
	controllerLog.Info("Get configmap data", "instance", instance)
	data, err := getPipelineConfigData(instance, agentPipeline, namespaces)
	if err != nil {
		return nil
	}

	return &corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        instance.Name,
			Namespace:   instance.GetNamespace(),
			Labels:      getLabels("agent"),
			Annotations: getAnnotations(),
		},
		Data: data,
	}
}

func (r *AgentPipelineReconciler) PipelineConfigMapFromCR(
	instance *loggerv1beta.VectorAgent, agentPipeline *loggerv1beta.VectorAgentPipeline, namespaces []string) *corev1.ConfigMap {
	controllerAgentPipelineLog.Info("Get configmap data", "instance", agentPipeline)
	data, err := getPipelineConfigData(instance, agentPipeline, namespaces)
	if err != nil {
		return nil
	}

	return &corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        instance.Name,
			Namespace:   instance.GetNamespace(),
			Labels:      getLabels("agent"),
			Annotations: getAnnotations(),
		},
		Data: data,
	}
}
