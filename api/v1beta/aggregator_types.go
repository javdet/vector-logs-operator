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

// VectorAgentSpec defines the desired state of VectorAgent
// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
// Important: Run "make" to regenerate code after modifying this file

// Foo is an example field of VectorAggregator. Edit agent_types.go to remove/update

type VectorAggregatorSpec struct {
	Image                 string                    `json:"image,omitempty"`
	Tag                   string                    `json:"tag,omitempty"`
	MetricsScrapeInterval int                       `json:"metricsScrapeInterval,omitempty"`
	InternalLogs          bool                      `json:"internalLogs,omitempty"`
	LogLevel              string                    `json:"logLevel,omitempty"`
	Resources             Resources                 `json:"resources,omitempty"`
	Sources               []VectorAggregatorSources `json:"sources"`
	Transforms            string                    `json:"transforms,omitempty"`
	Sinks                 []VectorAggregatorSinks   `json:"sinks"`
}

type VectorAggregatorSources struct {
	Kafka  VectorPipelineSourcesKafka  `json:"kafka,omitempty"`
	Vector VectorPipelineSourcesVector `json:"vector,omitempty"`
	AMQP   VectorAggregatorSourcesAMQP `json:"amqp,omitempty"`
	S3     VectorAggregatorSourcesS3   `json:"s3,omitempty"`
}

type VectorPipelineSourcesKafka struct {
	Name             string                       `json:"name,omitempty"`
	BootstrapServers string                       `json:"bootstrapServers"`
	KeyField         string                       `json:"keyField,omitempty"`
	Topics           []string                     `json:"topics,omitempty"`
	Decoding         string                       `json:"decoding,omitempty"`
	GroupID          string                       `json:"groupID"`
	Sasl             VectorPipelineSinksKafkaSasl `json:"sasl,omitempty"`
	AutoOffsetReset  string                       `json:"autoOffsetReset"`
}

type VectorPipelineSourcesVector struct {
	Name    string `json:"name,omitempty"`
	Host    string `json:"host"`
	Port    int32  `json:"port"`
	Version string `json:"version,omitempty"`
}

type VectorAggregatorSourcesAMQP struct {
	Name       string `json:"name,omitempty"`
	Connection string `json:"connection"`
	GroupID    string `json:"groupID"`
	OffsetKey  string `json:"offsetKey,omitempty"`
}

type VectorAggregatorSourcesS3 struct {
	Name        string `json:"name,omitempty"`
	Region      string `json:"region"`
	SQS         string `json:"SQS"`
	Compression string `json:"compression,omitempty"`
	Endpoint    string `json:"endpoint,omitempty"`
}

type VectorAggregatorSinks struct {
	S3            VectorPipelineSinksS3            `json:"s3,omitempty"`
	Console       VectorPipelineSinksConsole       `json:"console,omitempty"`
	File          VectorPipelineSinksFile          `json:"file,omitempty"`
	Elasticsearch VectorPipelineSinksElasticsearch `json:"elasticsearch,omitempty"`
	HTTP          VectorPipelineSinksHTTP          `json:"HTTP,omitempty"`
	Kafka         VectorPipelineSinksKafka         `json:"kafka,omitempty"`
	Loki          VectorPipelineSinksLoki          `json:"loki,omitempty"`
}

// +patchMergeKey=type
// +patchStrategy=merge
// +listType=map
// +listMapKey=type

// VectorAggregatorStatus defines the observed state of Agent
type VectorAggregatorStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// VectorAggregator is the Schema for the agents API
type VectorAggregator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VectorAggregatorSpec   `json:"spec,omitempty"`
	Status VectorAggregatorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VectorAggregatorList contains a list of VectorAggregator
type VectorAggregatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VectorAggregator `json:"items"`
}
