package controllers

import (
	"fmt"
	loggerv1beta "github.com/javdet/vector-logs-operator/api/v1beta"
	corev1 "k8s.io/api/core/v1"
)

// getLabels get common labels
func getLabels(namespace string) map[string]string {
	return map[string]string{
		"operator":                   "vector",
		"control-plane":              "vector-operator",
		"app.kubernetes.io/instance": fmt.Sprint("vector-", namespace),
		"app.kubernetes.io/name":     "vector",
	}
}

// getAnnotations get common annotations
func getAnnotations() map[string]string {
	return map[string]string{
		"reloader.stakater.com/match": "true",
	}
}

func getSecrets(pipeline []loggerv1beta.VectorPipelineSinks) []corev1.EnvFromSource {
	var secrets = []corev1.EnvFromSource{}

	for _, item := range pipeline {
		if item.Elasticsearch.Secret != "" {
			secrets = append(secrets, getSecretRef(item.Elasticsearch.Secret))
		}
		if item.HTTP.Secret != "" {
			secrets = append(secrets, getSecretRef(item.HTTP.Secret))
		}
		if item.Kafka.Sasl.Secret != "" {
			secrets = append(secrets, getSecretRef(item.Kafka.Sasl.Secret))
		}
		if item.S3.Secret != "" {
			secrets = append(secrets, getSecretRef(item.S3.Secret))
		}
	}

	return secrets
}

func getSecretRef(secretName string) corev1.EnvFromSource {
	var secret = corev1.EnvFromSource{}
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
