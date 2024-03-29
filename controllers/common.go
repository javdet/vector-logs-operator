package controllers

import (
	"bytes"
	"fmt"
	loggerv1beta "github.com/javdet/vector-logs-operator/api/v1beta"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"text/template"
)

const (
	nsLabel      = "vlo.io/logs"
	replicaLabel = "vlo.io/aggregator-replicas"
)

// getLabels get common labels
func getLabels(namespace string) map[string]string {
	return map[string]string{
		"operator":                    "vector",
		"control-plane":               "vector-operator",
		"app.kubernetes.io/instance":  fmt.Sprint("vector-", namespace),
		"app.kubernetes.io/name":      "vector",
		"app.kubernetes.io/component": "Agent",
	}
}

// getAnnotations get common annotations
func getAnnotations() map[string]string {
	return map[string]string{
		"reloader.stakater.com/match": "true",
	}
}

// getSecrets get Secrets by CRD spec
func getSecrets(pipeline []loggerv1beta.VectorPipelineSinks) []corev1.EnvFromSource {
	var secrets []corev1.EnvFromSource

	for _, item := range pipeline {
		if item.Elasticsearch.Secret.Name != "" {
			secrets = append(secrets, getSecretRef(item.Elasticsearch.Secret.Name))
		}
		if item.HTTP.Secret.Name != "" {
			secrets = append(secrets, getSecretRef(item.HTTP.Secret.Name))
		}
		if item.Kafka.Sasl.Secret.Name != "" {
			secrets = append(secrets, getSecretRef(item.Kafka.Sasl.Secret.Name))
		}
		if item.S3.Secret.Name != "" {
			secrets = append(secrets, getSecretRef(item.S3.Secret.Name))
		}
	}
	controllerAgentPipelineLog.Info("Get secrets", "secrets", secrets)
	return secrets
}

// getSecrets get Secrets by CRD spec
func getAggregatorSecrets(pipeline loggerv1beta.VectorAggregatorSpec) []corev1.EnvFromSource {
	var secrets []corev1.EnvFromSource

	for _, item := range pipeline.Sources {
		if item.Kafka.Sasl.Secret.Name != "" {
			secrets = append(secrets, getSecretRef(item.Kafka.Sasl.Secret.Name))
		}
	}

	for _, item := range pipeline.Sinks {
		if item.Elasticsearch.Secret.Name != "" {
			secrets = append(secrets, getSecretRef(item.Elasticsearch.Secret.Name))
		}
		if item.HTTP.Secret.Name != "" {
			secrets = append(secrets, getSecretRef(item.HTTP.Secret.Name))
		}
		if item.Kafka.Sasl.Secret.Name != "" {
			secrets = append(secrets, getSecretRef(item.Kafka.Sasl.Secret.Name))
		}
		if item.S3.Secret.Name != "" {
			secrets = append(secrets, getSecretRef(item.S3.Secret.Name))
		}
	}
	controllerAgentPipelineLog.Info("Get secrets", "secrets", secrets)
	return secrets
}

func getSecretRef(secretName string) corev1.EnvFromSource {
	var secret corev1.EnvFromSource
	var optionalSecret = true

	secret = corev1.EnvFromSource{
		SecretRef: &corev1.SecretEnvSource{
			LocalObjectReference: corev1.LocalObjectReference{
				Name: secretName,
			},
			Optional: &optionalSecret,
		},
	}

	return secret
}

// namespaceFilter event filter
func namespaceFilter() predicate.Predicate {
	var response bool

	return predicate.Funcs{
		CreateFunc: func(event event.CreateEvent) bool {
			response = false
			if _, ok := event.Object.(*corev1.Namespace); ok {
				if val, ok := event.Object.GetLabels()[nsLabel]; ok {
					if val == "true" {
						response = true
					}
				}
				if _, ok := event.Object.GetLabels()[replicaLabel]; ok {
					response = true
				}
			}
			if _, ok := event.Object.(*loggerv1beta.VectorAgentPipeline); ok {
				response = true
			}
			if _, ok := event.Object.(*loggerv1beta.VectorAggregator); ok {
				response = true
			}
			return response
		},
		UpdateFunc: func(updateEvent event.UpdateEvent) bool {
			response = false

			_, oldObject := updateEvent.ObjectOld.(*corev1.Namespace)
			_, newObject := updateEvent.ObjectNew.(*corev1.Namespace)
			if oldObject && newObject {
				oldValue, _ := updateEvent.ObjectOld.GetLabels()[nsLabel]
				newValue, _ := updateEvent.ObjectNew.GetLabels()[nsLabel]

				old := oldValue == "true"
				new := newValue == "true"

				response = old != new

				if oldValue == "true" && newValue == "true" {
					response = true
				}
			}

			_, oldObject = updateEvent.ObjectOld.(*loggerv1beta.VectorAgentPipeline)
			_, newObject = updateEvent.ObjectNew.(*loggerv1beta.VectorAgentPipeline)
			if oldObject && newObject {
				response = true
			}

			_, oldObject = updateEvent.ObjectOld.(*loggerv1beta.VectorAggregator)
			_, newObject = updateEvent.ObjectNew.(*loggerv1beta.VectorAggregator)
			if oldObject && newObject {
				response = true
			}

			return response
		},

		DeleteFunc: func(deleteEvent event.DeleteEvent) bool {
			response = false
			_, ok := deleteEvent.Object.(*corev1.Namespace)
			if ok {
				response = true
			}
			_, ok = deleteEvent.Object.(*loggerv1beta.VectorAgentPipeline)
			if ok {
				response = true
			}

			_, ok = deleteEvent.Object.(*loggerv1beta.VectorAggregator)
			if ok {
				response = true
			}
			return response
		},
	}
}

// getPipelineConfigData rendering te
func getPipelineConfigData(
	agent *loggerv1beta.VectorAgent, agentPipeline *loggerv1beta.VectorAgentPipeline, namespaces []string) (map[string]string, error) {
	var data = make(map[string]string)
	var vectorTpl bytes.Buffer

	pipeline := VectorAgentPipeline{
		Sources: PipelineSources{
			Kubernetes: PipelineSourcesKubernetes{
				PodAnnotations: agent.Spec.PodAnnotations,
			},
		},
		Sinks: PipelineSinks{
			Prometheus: PipelineSinksPrometheus{
				Namespace: agent.Name,
			},
		},
		Transforms: PipelineTransforms{
			Filter: PipelineTransformsFilter{
				Namespaces: namespaces,
			},
		},
		CRD: agentPipeline.Spec,
	}

	templateGeneral, err := template.ParseFiles("templates/vector-agent.yaml")
	if err != nil {
		controllerLog.Error(err, "failed parse config template", "template", "vector.yaml")
		return nil, err
	}

	if err := templateGeneral.Execute(&vectorTpl, pipeline); err != nil {
		controllerLog.Error(err, "failed generate config file", "template", "vector.yaml")
		return nil, err
	}
	data["vector.yaml"] = vectorTpl.String()

	return data, nil
}

func getContainerPorts(sources *loggerv1beta.VectorAggregator) []corev1.ContainerPort {
	var ports = []corev1.ContainerPort{{
		ContainerPort: 8686,
		Name:          "api",
		Protocol:      corev1.ProtocolTCP,
	}, {
		ContainerPort: 9100,
		Name:          "metrics",
		Protocol:      corev1.ProtocolTCP,
	}}

	for _, item := range sources.Spec.Sources {
		if item.Vector.Port != 0 {
			ports = append(ports, corev1.ContainerPort{
				ContainerPort: item.Vector.Port,
				Protocol:      corev1.ProtocolTCP,
				Name:          item.Vector.Name,
			})
		}
	}
	return ports
}

func getAggregatorPipelineConfigData(aggregator *loggerv1beta.VectorAggregator, namespace string) (map[string]string, error) {
	var data = make(map[string]string)
	var vectorTpl bytes.Buffer

	pipeline := VectorAggregatorPipeline{
		Sources: PipelineSources{
			Kafka: PipelineSourcesKafka{
				Topics: namespace,
			},
		},
		Sinks: PipelineSinks{
			Prometheus: PipelineSinksPrometheus{
				Namespace: aggregator.Name,
			},
		},
		CRD: aggregator.Spec,
	}
	templateGeneral, err := template.ParseFiles("templates/vector-aggregator.yaml")
	if err != nil {
		controllerAggregatorLog.Error(err, "failed parse config template", "template", "vector.yaml")
		return nil, err
	}

	if err := templateGeneral.Execute(&vectorTpl, pipeline); err != nil {
		controllerAggregatorLog.Error(err, "failed generate config file", "template", "vector.yaml")
		return nil, err
	}
	data["vector.yaml"] = vectorTpl.String()

	return data, nil
}
