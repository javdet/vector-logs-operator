package controllers

import (
	"fmt"
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

type VectorAggregatorPipeline struct {
	Sources    PipelineSources
	Transforms PipelineTransforms
	Sinks      PipelineSinks
	CRD        loggerv1beta.VectorAggregatorSpec
}

type PipelineSources struct {
	Metrics    PipelineSourcesMetrics
	Kubernetes PipelineSourcesKubernetes
	Kafka      PipelineSourcesKafka
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

type PipelineSourcesKafka struct {
	Topics string
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

func (r *AggregatorReconciler) PipelineConfigMapFromCR(instance *loggerv1beta.VectorAggregator, namespace string) *corev1.ConfigMap {
	controllerAggregatorLog.Info("Get configmap data", "instance", instance)
	data, err := getAggregatorPipelineConfigData(instance, namespace)
	if err != nil {
		return nil
	}

	return &corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        fmt.Sprint(instance.Name, "-", namespace),
			Namespace:   instance.GetNamespace(),
			Labels:      getLabels("agent"),
			Annotations: getAnnotations(),
		},
		Data: data,
	}
}
