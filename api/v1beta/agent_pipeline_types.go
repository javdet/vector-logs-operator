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

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// VectorAgentPipelineSpec defines the desired state of VectorAgentPipeline
type VectorAgentPipelineSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of VectorAgentPipeline. Edit agent_pipeline_types.go to remove/update
	Transforms string                `json:"transforms,omitempty"`
	Sinks      []VectorPipelineSinks `json:"sinks,omitempty"`
	Selector   map[string]string     `json:"selector"`
}

type VectorPipelineSinks struct {
	S3            VectorPipelineSinksS3            `json:"s3,omitempty"`
	Console       VectorPipelineSinksConsole       `json:"console,omitempty"`
	File          VectorPipelineSinksFile          `json:"file,omitempty"`
	Elasticsearch VectorPipelineSinksElasticsearch `json:"elasticsearch,omitempty"`
	HTTP          VectorPipelineSinksHTTP          `json:"HTTP,omitempty"`
	Kafka         VectorPipelineSinksKafka         `json:"kafka,omitempty"`
	Loki          VectorPipelineSinksLoki          `json:"loki,omitempty"`
	Vector        VectorPipelineSinksVector        `json:"vector,omitempty"`
}

type VectorPipelineSinksS3 struct {
	Name                 string                      `json:"name,omitempty"`
	Inputs               []string                    `json:"inputs,omitempty"`
	Bucket               string                      `json:"bucket"`
	Region               string                      `json:"region"`
	ACL                  string                      `json:"acl"`
	Compression          string                      `json:"compression,omitempty"`
	ContentType          string                      `json:"contentType,omitempty"`
	Encoding             string                      `json:"encoding,omitempty"`
	Endpoint             string                      `json:"endpoint,omitempty"`
	KeyPrefix            string                      `json:"keyPrefix,omitempty"`
	ServerSideEncryption string                      `json:"serverSideEncryption"`
	Secret               VectorPipelineSinksS3Secret `json:"secret"`
}

type VectorPipelineSinksS3Secret struct {
	Name string `json:"name"`
}

type VectorPipelineSinksConsole struct {
	Name     string   `json:"name,omitempty"`
	Inputs   []string `json:"inputs,omitempty"`
	Target   string   `json:"target,omitempty"`
	Encoding string   `json:"encoding,omitempty"`
}

type VectorPipelineSinksFile struct {
	Name        string   `json:"name,omitempty"`
	Inputs      []string `json:"inputs,omitempty"`
	Compression string   `json:"compression,omitempty"`
	Path        string   `json:"path"`
	Encoding    string   `json:"encoding,omitempty"`
}

type VectorPipelineSinksElasticsearch struct {
	Name        string                                 `json:"name,omitempty"`
	Inputs      []string                               `json:"inputs,omitempty"`
	Compression string                                 `json:"compression,omitempty"`
	Endpoints   []string                               `json:"endpoints,omitempty"`
	Pipeline    string                                 `json:"pipeline,omitempty"`
	Index       string                                 `json:"index,omitempty"`
	Mode        string                                 `json:"mode,omitempty"`
	Secret      VectorPipelineSinksElasticsearchSecret `json:"secret,omitempty"`
	IDKey       string                                 `json:"idKey,omitempty"`
	TLSCA       string                                 `json:"tlsCA,omitempty"`
}

type VectorPipelineSinksElasticsearchSecret struct {
	Name        string `json:"name"`
	UsernameKey string `json:"usernameKey"`
	PasswordKey string `json:"passwordKey"`
}

type VectorPipelineSinksHTTP struct {
	Name        string                        `json:"name,omitempty"`
	Inputs      []string                      `json:"inputs,omitempty"`
	Compression string                        `json:"compression,omitempty"`
	URI         string                        `json:"uri"`
	Encoding    string                        `json:"encoding,omitempty"`
	Secret      VectorPipelineSinksHTTPSecret `json:"secret,omitempty"`
	Method      string                        `json:"method,omitempty"`
	TLSCA       string                        `json:"tlsCA,omitempty"`
}

type VectorPipelineSinksHTTPSecret struct {
	Name        string `json:"name"`
	UsernameKey string `json:"usernameKey"`
	PasswordKey string `json:"passwordKey"`
}

type VectorPipelineSinksKafka struct {
	Name             string                       `json:"name,omitempty"`
	Inputs           []string                     `json:"inputs,omitempty"`
	BootstrapServers string                       `json:"bootstrapServers"`
	KeyField         string                       `json:"keyField,omitempty"`
	Topic            string                       `json:"topic"`
	Compression      string                       `json:"compression,omitempty"`
	Encoding         string                       `json:"encoding,omitempty"`
	Sasl             VectorPipelineSinksKafkaSasl `json:"sasl,omitempty"`
}

type VectorPipelineSinksKafkaSasl struct {
	Mechanism string                             `json:"mechanism"`
	Secret    VectorPipelineSinksKafkaSaslSecret `json:"secret"`
}

type VectorPipelineSinksKafkaSaslSecret struct {
	Name        string `json:"name"`
	UsernameKey string `json:"usernameKey"`
	PasswordKey string `json:"passwordKey"`
}

type VectorPipelineSinksLoki struct {
	Name         string            `json:"name,omitempty"`
	Inputs       []string          `json:"inputs,omitempty"`
	Endpoint     string            `json:"endpoint"`
	Labels       map[string]string `json:"labels,omitempty"`
	Compression  string            `json:"compression,omitempty"`
	Encoding     string            `json:"encoding,omitempty"`
	ExceptFields []string          `json:"exceptFields,omitempty"`
	TenantId     string            `json:"tenantId,omitempty"`
}

type VectorPipelineSinksVector struct {
	Name        string   `json:"name,omitempty"`
	Inputs      []string `json:"inputs,omitempty"`
	Address     string   `json:"address"`
	Compression string   `json:"compression,omitempty"`
}

// +patchMergeKey=type
// +patchStrategy=merge
// +listType=map
// +listMapKey=type
// VectorAgentStatus defines the observed state of Agent

type VectorAgentPipelineStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// VectorAgentPipeline is the Schema for the agents API
type VectorAgentPipeline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VectorAgentPipelineSpec   `json:"spec,omitempty"`
	Status VectorAgentPipelineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VectorAgentPipelineList contains a list of VectorAgentPipeline
type VectorAgentPipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VectorAgentPipeline `json:"items"`
}
