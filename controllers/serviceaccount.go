package controllers

import (
	loggerv1beta "github.com/javdet/vector-logs-operator/api/v1beta"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (r *AgentReconciler) serviceAccountFromCR(instance *loggerv1beta.VectorAgent) *corev1.ServiceAccount {
	controllerLog.Info("Create ServiceAccount", "instance", instance)
	return &corev1.ServiceAccount{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ServiceAccount",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        instance.Name,
			Namespace:   instance.Namespace,
			Labels:      getLabels("agent"),
			Annotations: getAnnotations(),
		},
	}
}

func (r *AgentReconciler) clusterRoleFromCR(instance *loggerv1beta.VectorAgent) *rbacv1.ClusterRole {
	controllerLog.Info("Create ClusterRole", "instance", instance)
	return &rbacv1.ClusterRole{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ClusterRole",
			APIVersion: "rbac.authorization.k8s.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        instance.Name,
			Labels:      getLabels("agent"),
			Annotations: getAnnotations(),
		},
		Rules: []rbacv1.PolicyRule{{
			Verbs:     []string{"list", "watch"},
			APIGroups: []string{""},
			Resources: []string{"pods", "namespaces", "nodes"},
		}},
	}
}

func (r *AgentReconciler) clusterRoleBindingFromCR(instance *loggerv1beta.VectorAgent) *rbacv1.ClusterRoleBinding {
	controllerLog.Info("Create cluster ClusterRoleBinding", "instance", instance)
	return &rbacv1.ClusterRoleBinding{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ClusterRoleBinding",
			APIVersion: "rbac.authorization.k8s.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        instance.Name,
			Labels:      getLabels("agent"),
			Annotations: getAnnotations(),
		},
		Subjects: []rbacv1.Subject{{
			Kind:      "ServiceAccount",
			Name:      instance.Name,
			Namespace: instance.Namespace,
		}},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "ClusterRole",
			Name:     instance.Name,
		},
	}
}
