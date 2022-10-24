package controllers

import (
	"bytes"
	"fmt"
	loggerv1beta "github.com/javdet/vector-logs-operator/api/v1beta"
	"html/template"
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
	Metrics PipelineSourcesMetrics
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

func (r *AgentReconciler) PipelineConfigMapFromCR(instance *loggerv1beta.VectorAgent) *corev1.ConfigMap {
	data, err := r.getPipelineConfigData(instance)
	if err != nil {
		return nil
	}

	return &corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        fmt.Sprint(instance.Name),
			Namespace:   instance.GetNamespace(),
			Labels:      getLabels("agent"),
			Annotations: getAnnotations(),
		},
		Data: data,
	}
}

func (r *AgentReconciler) getPipelineConfigData(instance *loggerv1beta.VectorAgent) (map[string]string, error) {
	log := r.Log.WithName("creator").WithName("configmap")
	controllerLog.Info("Get configmap data", "instance", instance)

	var data = make(map[string]string)
	var vectorTpl bytes.Buffer

	templ, err := template.ParseFiles("templates/vector-agent.yaml")
	if err != nil {
		log.Error(err, "failed parse config template", "template", "vector.yaml")
		return nil, err
	}
	pipeline := VectorAgentPipeline{
		Sources: PipelineSources{
			Metrics: PipelineSourcesMetrics{
				Namespace: instance.Name,
			},
		},
		Sinks: PipelineSinks{
			Prometheus: PipelineSinksPrometheus{
				Namespace: instance.Name,
			},
		},
		CRD: loggerv1beta.VectorAgentPipelineSpec{
			Sinks: []loggerv1beta.VectorPipelineSinks{{
				S3: loggerv1beta.VectorPipelineSinksS3{
					Bucket: "",
				},
				Console: loggerv1beta.VectorPipelineSinksConsole{
					Target: "",
				},
				File: loggerv1beta.VectorPipelineSinksFile{
					Path: "",
				},
				Elasticsearch: loggerv1beta.VectorPipelineSinksElasticsearch{
					Endpoint: "",
				},
				HTTP: loggerv1beta.VectorPipelineSinksHTTP{
					URI: "",
				},
				Kafka: loggerv1beta.VectorPipelineSinksKafka{
					Topic: "",
				},
				Loki: loggerv1beta.VectorPipelineSinksLoki{
					Endpoint: "",
				},
				Vector: loggerv1beta.VectorPipelineSinksVector{
					Address: "",
				},
			}},
		},
	}

	if err := templ.Execute(&vectorTpl, pipeline); err != nil {
		log.Error(err, "failed generate config file", "template", "vector.yaml")
		return nil, err
	}
	data["vector.yaml"] = vectorTpl.String()

	return data, nil
}

func (r *AgentPipelineReconciler) PipelineConfigMapFromCR(
	instance *loggerv1beta.VectorAgentPipeline, vector string, namespace string, namespaces []string) *corev1.ConfigMap {
	data, err := r.getPipelineConfigData(instance, vector, namespaces)
	if err != nil {
		return nil
	}
	return &corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        vector,
			Namespace:   namespace,
			Labels:      getLabels("agent"),
			Annotations: getAnnotations(),
		},
		Data: data,
	}
}

func (r *AgentPipelineReconciler) getPipelineConfigData(
	instance *loggerv1beta.VectorAgentPipeline, vector string, namespaces []string) (map[string]string, error) {
	var data = make(map[string]string)
	var vectorTpl bytes.Buffer

	templ, err := template.ParseFiles("templates/vector-agent.yaml")
	if err != nil {
		return nil, err
	}
	pipeline := VectorAgentPipeline{
		Sources: PipelineSources{
			Metrics: PipelineSourcesMetrics{
				Namespace: vector,
			},
		},
		Sinks: PipelineSinks{
			Prometheus: PipelineSinksPrometheus{
				Namespace: vector,
			},
		},
		Transforms: PipelineTransforms{
			Filter: PipelineTransformsFilter{
				Namespaces: namespaces,
			},
		},
		CRD: instance.Spec,
	}

	if err := templ.Execute(&vectorTpl, pipeline); err != nil {
		return nil, err
	}
	data["vector.yaml"] = vectorTpl.String()

	return data, nil
}
