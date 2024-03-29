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

package v1beta

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type Limit struct {
	Cpu    string `json:"cpu,omitempty"`
	Memory string `json:"memory,omitempty"`
}

type Requests struct {
	Cpu    string `json:"cpu,omitempty"`
	Memory string `json:"memory,omitempty"`
}

type Resources struct {
	Limit    Limit    `json:"limits,omitempty"`
	Requests Requests `json:"requests,omitempty"`
}

// VectorAgentSpec defines the desired state of VectorAgent
// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
// Important: Run "make" to regenerate code after modifying this file

// Foo is an example field of VectorAgent. Edit agent_types.go to remove/update

type VectorAgentSpec struct {
	Image                 string            `json:"image,omitempty"`
	Tag                   string            `json:"tag,omitempty"`
	MetricsScrapeInterval int               `json:"metricsScrapeInterval,omitempty"`
	InternalLogs          bool              `json:"internalLogs,omitempty"`
	LogLevel              string            `json:"logLevel,omitempty"`
	Resources             Resources         `json:"resources,omitempty"`
	PodAnnotations        map[string]string `json:"podAnnotations,omitempty"`
}

// +patchMergeKey=type
// +patchStrategy=merge
// +listType=map
// +listMapKey=type
// VectorAgentStatus defines the observed state of Agent

type VectorAgentStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// VectorAgent is the Schema for the agents API
type VectorAgent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VectorAgentSpec   `json:"spec,omitempty"`
	Status VectorAgentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VectorAgentList contains a list of VectorAgent
type VectorAgentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VectorAgent `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VectorAgent{}, &VectorAgentList{})
	SchemeBuilder.Register(&VectorAgentPipeline{}, &VectorAgentPipelineList{})
	SchemeBuilder.Register(&VectorAggregator{}, &VectorAggregatorList{})
}
