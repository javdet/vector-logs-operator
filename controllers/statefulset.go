package controllers

import (
	"context"
	"fmt"
	loggerv1beta "github.com/javdet/vector-logs-operator/api/v1beta"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	"strconv"
)

func (r *AggregatorReconciler) statefulSetFromCR(instance *loggerv1beta.VectorAggregator, nameSpace string) *appsv1.StatefulSet {
	controllerAggregatorLog.Info("Create statefulset", "instance", instance)
	replicas := r.getReplicaCount(nameSpace)

	return &appsv1.StatefulSet{
		TypeMeta: metav1.TypeMeta{
			Kind:       "StatefulSet",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        fmt.Sprint(instance.Name, "-", nameSpace),
			Namespace:   instance.GetNamespace(),
			Labels:      getLabels(nameSpace),
			Annotations: getAnnotations(),
		},
		Spec: appsv1.StatefulSetSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: getLabels(nameSpace),
			},
			Replicas: &replicas,
			UpdateStrategy: appsv1.StatefulSetUpdateStrategy{
				Type: appsv1.RollingUpdateStatefulSetStrategyType,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: getLabels(nameSpace),
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Name:  instance.Name,
						Image: fmt.Sprint(instance.Spec.Image, ":", instance.Spec.Tag),
						Args: []string{
							"--config-dir",
							"/etc/vector",
						},
						Ports: getContainerPorts(instance),
						Resources: corev1.ResourceRequirements{
							Limits: corev1.ResourceList{
								"cpu":    resource.MustParse(instance.Spec.Resources.Limit.Cpu),
								"memory": resource.MustParse(instance.Spec.Resources.Limit.Memory),
							},
							Requests: corev1.ResourceList{
								"cpu":    resource.MustParse(instance.Spec.Resources.Requests.Cpu),
								"memory": resource.MustParse(instance.Spec.Resources.Requests.Memory),
							},
						},
						VolumeMounts: []corev1.VolumeMount{{
							Name:      "data",
							MountPath: "/vector/data",
						}, {
							Name:      "config",
							MountPath: "/etc/vector/",
							ReadOnly:  true,
						}},
						EnvFrom: getAggregatorSecrets(instance.Spec),
					}},
					Volumes: []corev1.Volume{{
						Name: "config",
						VolumeSource: corev1.VolumeSource{
							ConfigMap: &corev1.ConfigMapVolumeSource{
								LocalObjectReference: corev1.LocalObjectReference{
									Name: fmt.Sprint(instance.Name, "-", nameSpace),
								},
							},
						},
					}, {
						Name: "data",
						VolumeSource: corev1.VolumeSource{
							HostPath: &corev1.HostPathVolumeSource{
								Path: "/var/lib/vector",
							},
						},
					}},
					ServiceAccountName: fmt.Sprint(instance.Name, "-", nameSpace),
					Tolerations: []corev1.Toleration{{
						Effect:   corev1.TaintEffectNoSchedule,
						Key:      "node-role.kubernetes.io/master",
						Operator: corev1.TolerationOpExists,
					}, {
						Effect:   corev1.TaintEffectNoExecute,
						Operator: corev1.TolerationOpExists,
					}, {
						Effect:   corev1.TaintEffectNoSchedule,
						Operator: corev1.TolerationOpExists,
					}},
				},
			},
		},
	}
}

func (r *AggregatorReconciler) statefulSetServiceFromCR(instance *loggerv1beta.VectorAggregator, namespace string) *corev1.Service {
	controllerLog.Info("Create service", "instance", instance)
	return &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprint(instance.Name, "-", namespace),
			Namespace: instance.Namespace,
			Labels:    getLabels("aggregator"),
		},
		Spec: corev1.ServiceSpec{
			Selector: getLabels("aggregator"),
			Ports: []corev1.ServicePort{{
				Name:       "api",
				Port:       8686,
				Protocol:   corev1.ProtocolTCP,
				TargetPort: intstr.FromInt(8686),
			}, {
				Name:       "metrics",
				Port:       9100,
				Protocol:   corev1.ProtocolTCP,
				TargetPort: intstr.FromInt(9100),
			}},
		},
	}
}

// getReplicaCount get Replicas from namespace label
func (r *AggregatorReconciler) getReplicaCount(namespaceName string) int32 {
	var Count = 1
	namespace := &corev1.Namespace{}

	if err := r.GetClient().Get(context.TODO(), types.NamespacedName{
		Namespace: "",
		Name:      namespaceName,
	}, namespace); err != nil {
		r.Log.Error(err, "failed get namespace object", namespaceName)
	}

	if annotation, ok := namespace.GetAnnotations()[replicaLabel]; ok {
		Count, _ = strconv.Atoi(annotation)
	}

	return int32(Count)
}
