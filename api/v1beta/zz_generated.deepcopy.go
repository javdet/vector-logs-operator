//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Limit) DeepCopyInto(out *Limit) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Limit.
func (in *Limit) DeepCopy() *Limit {
	if in == nil {
		return nil
	}
	out := new(Limit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Requests) DeepCopyInto(out *Requests) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Requests.
func (in *Requests) DeepCopy() *Requests {
	if in == nil {
		return nil
	}
	out := new(Requests)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resources) DeepCopyInto(out *Resources) {
	*out = *in
	out.Limit = in.Limit
	out.Requests = in.Requests
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resources.
func (in *Resources) DeepCopy() *Resources {
	if in == nil {
		return nil
	}
	out := new(Resources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VectorAgent) DeepCopyInto(out *VectorAgent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VectorAgent.
func (in *VectorAgent) DeepCopy() *VectorAgent {
	if in == nil {
		return nil
	}
	out := new(VectorAgent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VectorAgent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VectorAgentList) DeepCopyInto(out *VectorAgentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VectorAgent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VectorAgentList.
func (in *VectorAgentList) DeepCopy() *VectorAgentList {
	if in == nil {
		return nil
	}
	out := new(VectorAgentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VectorAgentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VectorAgentPipeline) DeepCopyInto(out *VectorAgentPipeline) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VectorAgentPipeline.
func (in *VectorAgentPipeline) DeepCopy() *VectorAgentPipeline {
	if in == nil {
		return nil
	}
	out := new(VectorAgentPipeline)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VectorAgentPipeline) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VectorAgentPipelineList) DeepCopyInto(out *VectorAgentPipelineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VectorAgentPipeline, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VectorAgentPipelineList.
func (in *VectorAgentPipelineList) DeepCopy() *VectorAgentPipelineList {
	if in == nil {
		return nil
	}
	out := new(VectorAgentPipelineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VectorAgentPipelineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VectorAgentPipelineSpec) DeepCopyInto(out *VectorAgentPipelineSpec) {
	*out = *in
	if in.Sinks != nil {
		in, out := &in.Sinks, &out.Sinks
		*out = make([]VectorPipelineSinks, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VectorAgentPipelineSpec.
func (in *VectorAgentPipelineSpec) DeepCopy() *VectorAgentPipelineSpec {
	if in == nil {
		return nil
	}
	out := new(VectorAgentPipelineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VectorAgentPipelineStatus) DeepCopyInto(out *VectorAgentPipelineStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VectorAgentPipelineStatus.
func (in *VectorAgentPipelineStatus) DeepCopy() *VectorAgentPipelineStatus {
	if in == nil {
		return nil
	}
	out := new(VectorAgentPipelineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VectorAgentSpec) DeepCopyInto(out *VectorAgentSpec) {
	*out = *in
	out.Resources = in.Resources
	if in.PodAnnotations != nil {
		in, out := &in.PodAnnotations, &out.PodAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VectorAgentSpec.
func (in *VectorAgentSpec) DeepCopy() *VectorAgentSpec {
	if in == nil {
		return nil
	}
	out := new(VectorAgentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VectorAgentStatus) DeepCopyInto(out *VectorAgentStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VectorAgentStatus.
func (in *VectorAgentStatus) DeepCopy() *VectorAgentStatus {
	if in == nil {
		return nil
	}
	out := new(VectorAgentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VectorPipelineSinks) DeepCopyInto(out *VectorPipelineSinks) {
	*out = *in
	in.S3.DeepCopyInto(&out.S3)
	in.Console.DeepCopyInto(&out.Console)
	in.File.DeepCopyInto(&out.File)
	in.Elasticsearch.DeepCopyInto(&out.Elasticsearch)
	in.HTTP.DeepCopyInto(&out.HTTP)
	in.Kafka.DeepCopyInto(&out.Kafka)
	in.Loki.DeepCopyInto(&out.Loki)
	in.Vector.DeepCopyInto(&out.Vector)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VectorPipelineSinks.
func (in *VectorPipelineSinks) DeepCopy() *VectorPipelineSinks {
	if in == nil {
		return nil
	}
	out := new(VectorPipelineSinks)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VectorPipelineSinksConsole) DeepCopyInto(out *VectorPipelineSinksConsole) {
	*out = *in
	if in.Inputs != nil {
		in, out := &in.Inputs, &out.Inputs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VectorPipelineSinksConsole.
func (in *VectorPipelineSinksConsole) DeepCopy() *VectorPipelineSinksConsole {
	if in == nil {
		return nil
	}
	out := new(VectorPipelineSinksConsole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VectorPipelineSinksElasticsearch) DeepCopyInto(out *VectorPipelineSinksElasticsearch) {
	*out = *in
	if in.Inputs != nil {
		in, out := &in.Inputs, &out.Inputs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Secret = in.Secret
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VectorPipelineSinksElasticsearch.
func (in *VectorPipelineSinksElasticsearch) DeepCopy() *VectorPipelineSinksElasticsearch {
	if in == nil {
		return nil
	}
	out := new(VectorPipelineSinksElasticsearch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VectorPipelineSinksElasticsearchSecret) DeepCopyInto(out *VectorPipelineSinksElasticsearchSecret) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VectorPipelineSinksElasticsearchSecret.
func (in *VectorPipelineSinksElasticsearchSecret) DeepCopy() *VectorPipelineSinksElasticsearchSecret {
	if in == nil {
		return nil
	}
	out := new(VectorPipelineSinksElasticsearchSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VectorPipelineSinksFile) DeepCopyInto(out *VectorPipelineSinksFile) {
	*out = *in
	if in.Inputs != nil {
		in, out := &in.Inputs, &out.Inputs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VectorPipelineSinksFile.
func (in *VectorPipelineSinksFile) DeepCopy() *VectorPipelineSinksFile {
	if in == nil {
		return nil
	}
	out := new(VectorPipelineSinksFile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VectorPipelineSinksHTTP) DeepCopyInto(out *VectorPipelineSinksHTTP) {
	*out = *in
	if in.Inputs != nil {
		in, out := &in.Inputs, &out.Inputs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Secret = in.Secret
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VectorPipelineSinksHTTP.
func (in *VectorPipelineSinksHTTP) DeepCopy() *VectorPipelineSinksHTTP {
	if in == nil {
		return nil
	}
	out := new(VectorPipelineSinksHTTP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VectorPipelineSinksHTTPSecret) DeepCopyInto(out *VectorPipelineSinksHTTPSecret) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VectorPipelineSinksHTTPSecret.
func (in *VectorPipelineSinksHTTPSecret) DeepCopy() *VectorPipelineSinksHTTPSecret {
	if in == nil {
		return nil
	}
	out := new(VectorPipelineSinksHTTPSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VectorPipelineSinksKafka) DeepCopyInto(out *VectorPipelineSinksKafka) {
	*out = *in
	if in.Inputs != nil {
		in, out := &in.Inputs, &out.Inputs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Sasl = in.Sasl
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VectorPipelineSinksKafka.
func (in *VectorPipelineSinksKafka) DeepCopy() *VectorPipelineSinksKafka {
	if in == nil {
		return nil
	}
	out := new(VectorPipelineSinksKafka)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VectorPipelineSinksKafkaSasl) DeepCopyInto(out *VectorPipelineSinksKafkaSasl) {
	*out = *in
	out.Secret = in.Secret
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VectorPipelineSinksKafkaSasl.
func (in *VectorPipelineSinksKafkaSasl) DeepCopy() *VectorPipelineSinksKafkaSasl {
	if in == nil {
		return nil
	}
	out := new(VectorPipelineSinksKafkaSasl)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VectorPipelineSinksKafkaSaslSecret) DeepCopyInto(out *VectorPipelineSinksKafkaSaslSecret) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VectorPipelineSinksKafkaSaslSecret.
func (in *VectorPipelineSinksKafkaSaslSecret) DeepCopy() *VectorPipelineSinksKafkaSaslSecret {
	if in == nil {
		return nil
	}
	out := new(VectorPipelineSinksKafkaSaslSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VectorPipelineSinksLoki) DeepCopyInto(out *VectorPipelineSinksLoki) {
	*out = *in
	if in.Inputs != nil {
		in, out := &in.Inputs, &out.Inputs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VectorPipelineSinksLoki.
func (in *VectorPipelineSinksLoki) DeepCopy() *VectorPipelineSinksLoki {
	if in == nil {
		return nil
	}
	out := new(VectorPipelineSinksLoki)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VectorPipelineSinksS3) DeepCopyInto(out *VectorPipelineSinksS3) {
	*out = *in
	if in.Inputs != nil {
		in, out := &in.Inputs, &out.Inputs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Secret = in.Secret
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VectorPipelineSinksS3.
func (in *VectorPipelineSinksS3) DeepCopy() *VectorPipelineSinksS3 {
	if in == nil {
		return nil
	}
	out := new(VectorPipelineSinksS3)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VectorPipelineSinksS3Secret) DeepCopyInto(out *VectorPipelineSinksS3Secret) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VectorPipelineSinksS3Secret.
func (in *VectorPipelineSinksS3Secret) DeepCopy() *VectorPipelineSinksS3Secret {
	if in == nil {
		return nil
	}
	out := new(VectorPipelineSinksS3Secret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VectorPipelineSinksVector) DeepCopyInto(out *VectorPipelineSinksVector) {
	*out = *in
	if in.Inputs != nil {
		in, out := &in.Inputs, &out.Inputs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VectorPipelineSinksVector.
func (in *VectorPipelineSinksVector) DeepCopy() *VectorPipelineSinksVector {
	if in == nil {
		return nil
	}
	out := new(VectorPipelineSinksVector)
	in.DeepCopyInto(out)
	return out
}
