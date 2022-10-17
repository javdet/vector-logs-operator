package controllers

import (
	loggerv1beta "github.com/javdet/vector-logs-operator/api/v1beta"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func (r *AgentReconciler) daemonSetFromCR(instance *loggerv1beta.VectorAgent) *appsv1.DaemonSet {
	var secrets = []loggerv1beta.VectorPipelineSinks{}

	return &appsv1.DaemonSet{
		TypeMeta: metav1.TypeMeta{
			Kind:       "DaemonSet",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        instance.Spec.Name,
			Namespace:   instance.Namespace,
			Labels:      getLabels("agent"),
			Annotations: getAnnotations(),
		},
		Spec: appsv1.DaemonSetSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: getLabels("agent"),
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: getLabels("agent"),
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Name:  instance.Spec.Name,
						Image: instance.Spec.Image,
						Args: []string{
							"--config-dir",
							"/etc/vector",
						},
						WorkingDir: "",
						Ports: []corev1.ContainerPort{{
							ContainerPort: 8686,
							Name:          "api",
							Protocol:      corev1.ProtocolTCP,
						}, {
							ContainerPort: 9100,
							Name:          "metrics",
							Protocol:      corev1.ProtocolTCP,
						}},
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
						Env: []corev1.EnvVar{{
							Name:  "PROCFS_ROOT",
							Value: "/host/proc",
						}, {
							Name:  "SYSFS_ROOT",
							Value: "/host/sys",
						}},
						EnvFrom: getSecrets(secrets),
						VolumeMounts: []corev1.VolumeMount{{
							Name:      "data",
							MountPath: "/vector/data",
						}, {
							Name:      "config",
							MountPath: "/etc/vector/",
							ReadOnly:  true,
						}, {
							Name:      "var-log",
							MountPath: "/var/log/",
							ReadOnly:  true,
						}, {
							Name:      "var-lib",
							MountPath: "/var/lib",
							ReadOnly:  true,
						}, {
							Name:      "procfs",
							MountPath: "/host/proc",
							ReadOnly:  true,
						}, {
							Name:      "sysfs",
							MountPath: "/host/sys",
							ReadOnly:  true,
						}},
					}},
					Volumes: []corev1.Volume{{
						Name: "config",
						VolumeSource: corev1.VolumeSource{
							ConfigMap: &corev1.ConfigMapVolumeSource{
								LocalObjectReference: corev1.LocalObjectReference{Name: instance.Name},
							},
						},
					}, {
						Name: "data",
						VolumeSource: corev1.VolumeSource{
							HostPath: &corev1.HostPathVolumeSource{
								Path: "/var/lib/vector",
							},
						},
					}, {
						Name: "var-log",
						VolumeSource: corev1.VolumeSource{
							HostPath: &corev1.HostPathVolumeSource{
								Path: "/var/log/",
							},
						},
					}, {
						Name: "var-lib",
						VolumeSource: corev1.VolumeSource{
							HostPath: &corev1.HostPathVolumeSource{
								Path: "/var/lib/",
							},
						},
					}, {
						Name: "procfs",
						VolumeSource: corev1.VolumeSource{
							HostPath: &corev1.HostPathVolumeSource{
								Path: "/proc",
							},
						},
					}, {
						Name: "sysfs",
						VolumeSource: corev1.VolumeSource{
							HostPath: &corev1.HostPathVolumeSource{
								Path: "/sys",
							},
						},
					}},
					ServiceAccountName: instance.Name,
				},
			},
		},
	}
}

func (r *AgentReconciler) daemonSetServiceFromCR(instance *loggerv1beta.VectorAgent) *corev1.Service {
	return &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      instance.Spec.Name,
			Namespace: instance.Namespace,
			Labels:    getLabels("agent"),
		},
		Spec: corev1.ServiceSpec{
			Selector: getLabels("agent"),
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

func (r *AgentPipelineReconciler) daemonSetFromCR(
	instance *loggerv1beta.VectorAgentPipeline, agent *loggerv1beta.VectorAgent) *appsv1.DaemonSet {

	return &appsv1.DaemonSet{
		TypeMeta: metav1.TypeMeta{
			Kind:       "DaemonSet",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        agent.Spec.Name,
			Namespace:   agent.Namespace,
			Labels:      getLabels("agent"),
			Annotations: getAnnotations(),
		},
		Spec: appsv1.DaemonSetSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: getLabels("agent"),
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: getLabels("agent"),
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Name:  agent.Spec.Name,
						Image: agent.Spec.Image,
						Args: []string{
							"--config-dir",
							"/etc/vector",
						},
						WorkingDir: "",
						Ports: []corev1.ContainerPort{{
							ContainerPort: 8686,
							Name:          "api",
							Protocol:      corev1.ProtocolTCP,
						}, {
							ContainerPort: 9100,
							Name:          "metrics",
							Protocol:      corev1.ProtocolTCP,
						}},
						Resources: corev1.ResourceRequirements{
							Limits: corev1.ResourceList{
								"cpu":    resource.MustParse(agent.Spec.Resources.Limit.Cpu),
								"memory": resource.MustParse(agent.Spec.Resources.Limit.Memory),
							},
							Requests: corev1.ResourceList{
								"cpu":    resource.MustParse(agent.Spec.Resources.Requests.Cpu),
								"memory": resource.MustParse(agent.Spec.Resources.Requests.Memory),
							},
						},
						Env: []corev1.EnvVar{{
							Name:  "PROCFS_ROOT",
							Value: "/host/proc",
						}, {
							Name:  "SYSFS_ROOT",
							Value: "/host/sys",
						}},
						EnvFrom: getSecrets(instance.Spec.Sinks),
						VolumeMounts: []corev1.VolumeMount{{
							Name:      "data",
							MountPath: "/vector/data",
						}, {
							Name:      "config",
							MountPath: "/etc/vector/",
							ReadOnly:  true,
						}, {
							Name:      "var-log",
							MountPath: "/var/log/",
							ReadOnly:  true,
						}, {
							Name:      "var-lib",
							MountPath: "/var/lib",
							ReadOnly:  true,
						}, {
							Name:      "procfs",
							MountPath: "/host/proc",
							ReadOnly:  true,
						}, {
							Name:      "sysfs",
							MountPath: "/host/sys",
							ReadOnly:  true,
						}},
					}},
					Volumes: []corev1.Volume{{
						Name: "config",
						VolumeSource: corev1.VolumeSource{
							ConfigMap: &corev1.ConfigMapVolumeSource{
								LocalObjectReference: corev1.LocalObjectReference{Name: instance.Name},
							},
						},
					}, {
						Name: "data",
						VolumeSource: corev1.VolumeSource{
							HostPath: &corev1.HostPathVolumeSource{
								Path: "/var/lib/vector",
							},
						},
					}, {
						Name: "var-log",
						VolumeSource: corev1.VolumeSource{
							HostPath: &corev1.HostPathVolumeSource{
								Path: "/var/log/",
							},
						},
					}, {
						Name: "var-lib",
						VolumeSource: corev1.VolumeSource{
							HostPath: &corev1.HostPathVolumeSource{
								Path: "/var/lib/",
							},
						},
					}, {
						Name: "procfs",
						VolumeSource: corev1.VolumeSource{
							HostPath: &corev1.HostPathVolumeSource{
								Path: "/proc",
							},
						},
					}, {
						Name: "sysfs",
						VolumeSource: corev1.VolumeSource{
							HostPath: &corev1.HostPathVolumeSource{
								Path: "/sys",
							},
						},
					}},
					ServiceAccountName: instance.Name,
				},
			},
		},
	}
}