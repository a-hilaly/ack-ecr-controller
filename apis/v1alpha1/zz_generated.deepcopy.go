// +build !ignore_autogenerated

// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionConfiguration) DeepCopyInto(out *EncryptionConfiguration) {
	*out = *in
	if in.EncryptionType != nil {
		in, out := &in.EncryptionType, &out.EncryptionType
		*out = new(string)
		**out = **in
	}
	if in.KMSKey != nil {
		in, out := &in.KMSKey, &out.KMSKey
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionConfiguration.
func (in *EncryptionConfiguration) DeepCopy() *EncryptionConfiguration {
	if in == nil {
		return nil
	}
	out := new(EncryptionConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Image) DeepCopyInto(out *Image) {
	*out = *in
	if in.RegistryID != nil {
		in, out := &in.RegistryID, &out.RegistryID
		*out = new(string)
		**out = **in
	}
	if in.RepositoryName != nil {
		in, out := &in.RepositoryName, &out.RepositoryName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Image.
func (in *Image) DeepCopy() *Image {
	if in == nil {
		return nil
	}
	out := new(Image)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageDetail) DeepCopyInto(out *ImageDetail) {
	*out = *in
	if in.RegistryID != nil {
		in, out := &in.RegistryID, &out.RegistryID
		*out = new(string)
		**out = **in
	}
	if in.RepositoryName != nil {
		in, out := &in.RepositoryName, &out.RepositoryName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageDetail.
func (in *ImageDetail) DeepCopy() *ImageDetail {
	if in == nil {
		return nil
	}
	out := new(ImageDetail)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageScanFinding) DeepCopyInto(out *ImageScanFinding) {
	*out = *in
	if in.URI != nil {
		in, out := &in.URI, &out.URI
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageScanFinding.
func (in *ImageScanFinding) DeepCopy() *ImageScanFinding {
	if in == nil {
		return nil
	}
	out := new(ImageScanFinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageScanningConfiguration) DeepCopyInto(out *ImageScanningConfiguration) {
	*out = *in
	if in.ScanOnPush != nil {
		in, out := &in.ScanOnPush, &out.ScanOnPush
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageScanningConfiguration.
func (in *ImageScanningConfiguration) DeepCopy() *ImageScanningConfiguration {
	if in == nil {
		return nil
	}
	out := new(ImageScanningConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationDestination) DeepCopyInto(out *ReplicationDestination) {
	*out = *in
	if in.RegistryID != nil {
		in, out := &in.RegistryID, &out.RegistryID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationDestination.
func (in *ReplicationDestination) DeepCopy() *ReplicationDestination {
	if in == nil {
		return nil
	}
	out := new(ReplicationDestination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Repository) DeepCopyInto(out *Repository) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Repository.
func (in *Repository) DeepCopy() *Repository {
	if in == nil {
		return nil
	}
	out := new(Repository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Repository) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryList) DeepCopyInto(out *RepositoryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Repository, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryList.
func (in *RepositoryList) DeepCopy() *RepositoryList {
	if in == nil {
		return nil
	}
	out := new(RepositoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RepositoryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositorySpec) DeepCopyInto(out *RepositorySpec) {
	*out = *in
	if in.EncryptionConfiguration != nil {
		in, out := &in.EncryptionConfiguration, &out.EncryptionConfiguration
		*out = new(EncryptionConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.ImageTagMutability != nil {
		in, out := &in.ImageTagMutability, &out.ImageTagMutability
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ScanConfig != nil {
		in, out := &in.ScanConfig, &out.ScanConfig
		*out = new(ImageScanningConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*Tag, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Tag)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositorySpec.
func (in *RepositorySpec) DeepCopy() *RepositorySpec {
	if in == nil {
		return nil
	}
	out := new(RepositorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryStatus) DeepCopyInto(out *RepositoryStatus) {
	*out = *in
	if in.ACKResourceMetadata != nil {
		in, out := &in.ACKResourceMetadata, &out.ACKResourceMetadata
		*out = new(corev1alpha1.ResourceMetadata)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]*corev1alpha1.Condition, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(corev1alpha1.Condition)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = (*in).DeepCopy()
	}
	if in.RegistryID != nil {
		in, out := &in.RegistryID, &out.RegistryID
		*out = new(string)
		**out = **in
	}
	if in.RepositoryURI != nil {
		in, out := &in.RepositoryURI, &out.RepositoryURI
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryStatus.
func (in *RepositoryStatus) DeepCopy() *RepositoryStatus {
	if in == nil {
		return nil
	}
	out := new(RepositoryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Repository_SDK) DeepCopyInto(out *Repository_SDK) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = (*in).DeepCopy()
	}
	if in.EncryptionConfiguration != nil {
		in, out := &in.EncryptionConfiguration, &out.EncryptionConfiguration
		*out = new(EncryptionConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.ImageScanningConfiguration != nil {
		in, out := &in.ImageScanningConfiguration, &out.ImageScanningConfiguration
		*out = new(ImageScanningConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.ImageTagMutability != nil {
		in, out := &in.ImageTagMutability, &out.ImageTagMutability
		*out = new(string)
		**out = **in
	}
	if in.RegistryID != nil {
		in, out := &in.RegistryID, &out.RegistryID
		*out = new(string)
		**out = **in
	}
	if in.RepositoryARN != nil {
		in, out := &in.RepositoryARN, &out.RepositoryARN
		*out = new(string)
		**out = **in
	}
	if in.RepositoryName != nil {
		in, out := &in.RepositoryName, &out.RepositoryName
		*out = new(string)
		**out = **in
	}
	if in.RepositoryURI != nil {
		in, out := &in.RepositoryURI, &out.RepositoryURI
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Repository_SDK.
func (in *Repository_SDK) DeepCopy() *Repository_SDK {
	if in == nil {
		return nil
	}
	out := new(Repository_SDK)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tag) DeepCopyInto(out *Tag) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tag.
func (in *Tag) DeepCopy() *Tag {
	if in == nil {
		return nil
	}
	out := new(Tag)
	in.DeepCopyInto(out)
	return out
}