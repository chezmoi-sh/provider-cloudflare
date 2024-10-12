//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSSEC) DeepCopyInto(out *DNSSEC) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSSEC.
func (in *DNSSEC) DeepCopy() *DNSSEC {
	if in == nil {
		return nil
	}
	out := new(DNSSEC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DNSSEC) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSSECInitParameters) DeepCopyInto(out *DNSSECInitParameters) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ModifiedOn != nil {
		in, out := &in.ModifiedOn, &out.ModifiedOn
		*out = new(string)
		**out = **in
	}
	if in.ZoneID != nil {
		in, out := &in.ZoneID, &out.ZoneID
		*out = new(string)
		**out = **in
	}
	if in.ZoneIDSelector != nil {
		in, out := &in.ZoneIDSelector, &out.ZoneIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSSECInitParameters.
func (in *DNSSECInitParameters) DeepCopy() *DNSSECInitParameters {
	if in == nil {
		return nil
	}
	out := new(DNSSECInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSSECList) DeepCopyInto(out *DNSSECList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DNSSEC, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSSECList.
func (in *DNSSECList) DeepCopy() *DNSSECList {
	if in == nil {
		return nil
	}
	out := new(DNSSECList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DNSSECList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSSECObservation) DeepCopyInto(out *DNSSECObservation) {
	*out = *in
	if in.Algorithm != nil {
		in, out := &in.Algorithm, &out.Algorithm
		*out = new(string)
		**out = **in
	}
	if in.Digest != nil {
		in, out := &in.Digest, &out.Digest
		*out = new(string)
		**out = **in
	}
	if in.DigestAlgorithm != nil {
		in, out := &in.DigestAlgorithm, &out.DigestAlgorithm
		*out = new(string)
		**out = **in
	}
	if in.DigestType != nil {
		in, out := &in.DigestType, &out.DigestType
		*out = new(string)
		**out = **in
	}
	if in.Ds != nil {
		in, out := &in.Ds, &out.Ds
		*out = new(string)
		**out = **in
	}
	if in.Flags != nil {
		in, out := &in.Flags, &out.Flags
		*out = new(float64)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.KeyTag != nil {
		in, out := &in.KeyTag, &out.KeyTag
		*out = new(float64)
		**out = **in
	}
	if in.KeyType != nil {
		in, out := &in.KeyType, &out.KeyType
		*out = new(string)
		**out = **in
	}
	if in.ModifiedOn != nil {
		in, out := &in.ModifiedOn, &out.ModifiedOn
		*out = new(string)
		**out = **in
	}
	if in.PublicKey != nil {
		in, out := &in.PublicKey, &out.PublicKey
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.ZoneID != nil {
		in, out := &in.ZoneID, &out.ZoneID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSSECObservation.
func (in *DNSSECObservation) DeepCopy() *DNSSECObservation {
	if in == nil {
		return nil
	}
	out := new(DNSSECObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSSECParameters) DeepCopyInto(out *DNSSECParameters) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ModifiedOn != nil {
		in, out := &in.ModifiedOn, &out.ModifiedOn
		*out = new(string)
		**out = **in
	}
	if in.ZoneID != nil {
		in, out := &in.ZoneID, &out.ZoneID
		*out = new(string)
		**out = **in
	}
	if in.ZoneIDSelector != nil {
		in, out := &in.ZoneIDSelector, &out.ZoneIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSSECParameters.
func (in *DNSSECParameters) DeepCopy() *DNSSECParameters {
	if in == nil {
		return nil
	}
	out := new(DNSSECParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSSECSpec) DeepCopyInto(out *DNSSECSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSSECSpec.
func (in *DNSSECSpec) DeepCopy() *DNSSECSpec {
	if in == nil {
		return nil
	}
	out := new(DNSSECSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSSECStatus) DeepCopyInto(out *DNSSECStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSSECStatus.
func (in *DNSSECStatus) DeepCopy() *DNSSECStatus {
	if in == nil {
		return nil
	}
	out := new(DNSSECStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataInitParameters) DeepCopyInto(out *DataInitParameters) {
	*out = *in
	if in.Algorithm != nil {
		in, out := &in.Algorithm, &out.Algorithm
		*out = new(float64)
		**out = **in
	}
	if in.Altitude != nil {
		in, out := &in.Altitude, &out.Altitude
		*out = new(float64)
		**out = **in
	}
	if in.Certificate != nil {
		in, out := &in.Certificate, &out.Certificate
		*out = new(string)
		**out = **in
	}
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = new(string)
		**out = **in
	}
	if in.Digest != nil {
		in, out := &in.Digest, &out.Digest
		*out = new(string)
		**out = **in
	}
	if in.DigestType != nil {
		in, out := &in.DigestType, &out.DigestType
		*out = new(float64)
		**out = **in
	}
	if in.Fingerprint != nil {
		in, out := &in.Fingerprint, &out.Fingerprint
		*out = new(string)
		**out = **in
	}
	if in.Flags != nil {
		in, out := &in.Flags, &out.Flags
		*out = new(string)
		**out = **in
	}
	if in.KeyTag != nil {
		in, out := &in.KeyTag, &out.KeyTag
		*out = new(float64)
		**out = **in
	}
	if in.LatDegrees != nil {
		in, out := &in.LatDegrees, &out.LatDegrees
		*out = new(float64)
		**out = **in
	}
	if in.LatDirection != nil {
		in, out := &in.LatDirection, &out.LatDirection
		*out = new(string)
		**out = **in
	}
	if in.LatMinutes != nil {
		in, out := &in.LatMinutes, &out.LatMinutes
		*out = new(float64)
		**out = **in
	}
	if in.LatSeconds != nil {
		in, out := &in.LatSeconds, &out.LatSeconds
		*out = new(float64)
		**out = **in
	}
	if in.LongDegrees != nil {
		in, out := &in.LongDegrees, &out.LongDegrees
		*out = new(float64)
		**out = **in
	}
	if in.LongDirection != nil {
		in, out := &in.LongDirection, &out.LongDirection
		*out = new(string)
		**out = **in
	}
	if in.LongMinutes != nil {
		in, out := &in.LongMinutes, &out.LongMinutes
		*out = new(float64)
		**out = **in
	}
	if in.LongSeconds != nil {
		in, out := &in.LongSeconds, &out.LongSeconds
		*out = new(float64)
		**out = **in
	}
	if in.MatchingType != nil {
		in, out := &in.MatchingType, &out.MatchingType
		*out = new(float64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Order != nil {
		in, out := &in.Order, &out.Order
		*out = new(float64)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(float64)
		**out = **in
	}
	if in.PrecisionHorz != nil {
		in, out := &in.PrecisionHorz, &out.PrecisionHorz
		*out = new(float64)
		**out = **in
	}
	if in.PrecisionVert != nil {
		in, out := &in.PrecisionVert, &out.PrecisionVert
		*out = new(float64)
		**out = **in
	}
	if in.Preference != nil {
		in, out := &in.Preference, &out.Preference
		*out = new(float64)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(float64)
		**out = **in
	}
	if in.Proto != nil {
		in, out := &in.Proto, &out.Proto
		*out = new(string)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(float64)
		**out = **in
	}
	if in.PublicKey != nil {
		in, out := &in.PublicKey, &out.PublicKey
		*out = new(string)
		**out = **in
	}
	if in.Regex != nil {
		in, out := &in.Regex, &out.Regex
		*out = new(string)
		**out = **in
	}
	if in.Replacement != nil {
		in, out := &in.Replacement, &out.Replacement
		*out = new(string)
		**out = **in
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(float64)
		**out = **in
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(string)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(float64)
		**out = **in
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = new(string)
		**out = **in
	}
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(float64)
		**out = **in
	}
	if in.Usage != nil {
		in, out := &in.Usage, &out.Usage
		*out = new(float64)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataInitParameters.
func (in *DataInitParameters) DeepCopy() *DataInitParameters {
	if in == nil {
		return nil
	}
	out := new(DataInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataObservation) DeepCopyInto(out *DataObservation) {
	*out = *in
	if in.Algorithm != nil {
		in, out := &in.Algorithm, &out.Algorithm
		*out = new(float64)
		**out = **in
	}
	if in.Altitude != nil {
		in, out := &in.Altitude, &out.Altitude
		*out = new(float64)
		**out = **in
	}
	if in.Certificate != nil {
		in, out := &in.Certificate, &out.Certificate
		*out = new(string)
		**out = **in
	}
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = new(string)
		**out = **in
	}
	if in.Digest != nil {
		in, out := &in.Digest, &out.Digest
		*out = new(string)
		**out = **in
	}
	if in.DigestType != nil {
		in, out := &in.DigestType, &out.DigestType
		*out = new(float64)
		**out = **in
	}
	if in.Fingerprint != nil {
		in, out := &in.Fingerprint, &out.Fingerprint
		*out = new(string)
		**out = **in
	}
	if in.Flags != nil {
		in, out := &in.Flags, &out.Flags
		*out = new(string)
		**out = **in
	}
	if in.KeyTag != nil {
		in, out := &in.KeyTag, &out.KeyTag
		*out = new(float64)
		**out = **in
	}
	if in.LatDegrees != nil {
		in, out := &in.LatDegrees, &out.LatDegrees
		*out = new(float64)
		**out = **in
	}
	if in.LatDirection != nil {
		in, out := &in.LatDirection, &out.LatDirection
		*out = new(string)
		**out = **in
	}
	if in.LatMinutes != nil {
		in, out := &in.LatMinutes, &out.LatMinutes
		*out = new(float64)
		**out = **in
	}
	if in.LatSeconds != nil {
		in, out := &in.LatSeconds, &out.LatSeconds
		*out = new(float64)
		**out = **in
	}
	if in.LongDegrees != nil {
		in, out := &in.LongDegrees, &out.LongDegrees
		*out = new(float64)
		**out = **in
	}
	if in.LongDirection != nil {
		in, out := &in.LongDirection, &out.LongDirection
		*out = new(string)
		**out = **in
	}
	if in.LongMinutes != nil {
		in, out := &in.LongMinutes, &out.LongMinutes
		*out = new(float64)
		**out = **in
	}
	if in.LongSeconds != nil {
		in, out := &in.LongSeconds, &out.LongSeconds
		*out = new(float64)
		**out = **in
	}
	if in.MatchingType != nil {
		in, out := &in.MatchingType, &out.MatchingType
		*out = new(float64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Order != nil {
		in, out := &in.Order, &out.Order
		*out = new(float64)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(float64)
		**out = **in
	}
	if in.PrecisionHorz != nil {
		in, out := &in.PrecisionHorz, &out.PrecisionHorz
		*out = new(float64)
		**out = **in
	}
	if in.PrecisionVert != nil {
		in, out := &in.PrecisionVert, &out.PrecisionVert
		*out = new(float64)
		**out = **in
	}
	if in.Preference != nil {
		in, out := &in.Preference, &out.Preference
		*out = new(float64)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(float64)
		**out = **in
	}
	if in.Proto != nil {
		in, out := &in.Proto, &out.Proto
		*out = new(string)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(float64)
		**out = **in
	}
	if in.PublicKey != nil {
		in, out := &in.PublicKey, &out.PublicKey
		*out = new(string)
		**out = **in
	}
	if in.Regex != nil {
		in, out := &in.Regex, &out.Regex
		*out = new(string)
		**out = **in
	}
	if in.Replacement != nil {
		in, out := &in.Replacement, &out.Replacement
		*out = new(string)
		**out = **in
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(float64)
		**out = **in
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(string)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(float64)
		**out = **in
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = new(string)
		**out = **in
	}
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(float64)
		**out = **in
	}
	if in.Usage != nil {
		in, out := &in.Usage, &out.Usage
		*out = new(float64)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataObservation.
func (in *DataObservation) DeepCopy() *DataObservation {
	if in == nil {
		return nil
	}
	out := new(DataObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataParameters) DeepCopyInto(out *DataParameters) {
	*out = *in
	if in.Algorithm != nil {
		in, out := &in.Algorithm, &out.Algorithm
		*out = new(float64)
		**out = **in
	}
	if in.Altitude != nil {
		in, out := &in.Altitude, &out.Altitude
		*out = new(float64)
		**out = **in
	}
	if in.Certificate != nil {
		in, out := &in.Certificate, &out.Certificate
		*out = new(string)
		**out = **in
	}
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = new(string)
		**out = **in
	}
	if in.Digest != nil {
		in, out := &in.Digest, &out.Digest
		*out = new(string)
		**out = **in
	}
	if in.DigestType != nil {
		in, out := &in.DigestType, &out.DigestType
		*out = new(float64)
		**out = **in
	}
	if in.Fingerprint != nil {
		in, out := &in.Fingerprint, &out.Fingerprint
		*out = new(string)
		**out = **in
	}
	if in.Flags != nil {
		in, out := &in.Flags, &out.Flags
		*out = new(string)
		**out = **in
	}
	if in.KeyTag != nil {
		in, out := &in.KeyTag, &out.KeyTag
		*out = new(float64)
		**out = **in
	}
	if in.LatDegrees != nil {
		in, out := &in.LatDegrees, &out.LatDegrees
		*out = new(float64)
		**out = **in
	}
	if in.LatDirection != nil {
		in, out := &in.LatDirection, &out.LatDirection
		*out = new(string)
		**out = **in
	}
	if in.LatMinutes != nil {
		in, out := &in.LatMinutes, &out.LatMinutes
		*out = new(float64)
		**out = **in
	}
	if in.LatSeconds != nil {
		in, out := &in.LatSeconds, &out.LatSeconds
		*out = new(float64)
		**out = **in
	}
	if in.LongDegrees != nil {
		in, out := &in.LongDegrees, &out.LongDegrees
		*out = new(float64)
		**out = **in
	}
	if in.LongDirection != nil {
		in, out := &in.LongDirection, &out.LongDirection
		*out = new(string)
		**out = **in
	}
	if in.LongMinutes != nil {
		in, out := &in.LongMinutes, &out.LongMinutes
		*out = new(float64)
		**out = **in
	}
	if in.LongSeconds != nil {
		in, out := &in.LongSeconds, &out.LongSeconds
		*out = new(float64)
		**out = **in
	}
	if in.MatchingType != nil {
		in, out := &in.MatchingType, &out.MatchingType
		*out = new(float64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Order != nil {
		in, out := &in.Order, &out.Order
		*out = new(float64)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(float64)
		**out = **in
	}
	if in.PrecisionHorz != nil {
		in, out := &in.PrecisionHorz, &out.PrecisionHorz
		*out = new(float64)
		**out = **in
	}
	if in.PrecisionVert != nil {
		in, out := &in.PrecisionVert, &out.PrecisionVert
		*out = new(float64)
		**out = **in
	}
	if in.Preference != nil {
		in, out := &in.Preference, &out.Preference
		*out = new(float64)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(float64)
		**out = **in
	}
	if in.Proto != nil {
		in, out := &in.Proto, &out.Proto
		*out = new(string)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(float64)
		**out = **in
	}
	if in.PublicKey != nil {
		in, out := &in.PublicKey, &out.PublicKey
		*out = new(string)
		**out = **in
	}
	if in.Regex != nil {
		in, out := &in.Regex, &out.Regex
		*out = new(string)
		**out = **in
	}
	if in.Replacement != nil {
		in, out := &in.Replacement, &out.Replacement
		*out = new(string)
		**out = **in
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(float64)
		**out = **in
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(string)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(float64)
		**out = **in
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = new(string)
		**out = **in
	}
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(float64)
		**out = **in
	}
	if in.Usage != nil {
		in, out := &in.Usage, &out.Usage
		*out = new(float64)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataParameters.
func (in *DataParameters) DeepCopy() *DataParameters {
	if in == nil {
		return nil
	}
	out := new(DataParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Record) DeepCopyInto(out *Record) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Record.
func (in *Record) DeepCopy() *Record {
	if in == nil {
		return nil
	}
	out := new(Record)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Record) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecordInitParameters) DeepCopyInto(out *RecordInitParameters) {
	*out = *in
	if in.AllowOverwrite != nil {
		in, out := &in.AllowOverwrite, &out.AllowOverwrite
		*out = new(bool)
		**out = **in
	}
	if in.Comment != nil {
		in, out := &in.Comment, &out.Comment
		*out = new(string)
		**out = **in
	}
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = new(string)
		**out = **in
	}
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make([]DataInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(float64)
		**out = **in
	}
	if in.Proxied != nil {
		in, out := &in.Proxied, &out.Proxied
		*out = new(bool)
		**out = **in
	}
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(float64)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	if in.ZoneID != nil {
		in, out := &in.ZoneID, &out.ZoneID
		*out = new(string)
		**out = **in
	}
	if in.ZoneIDSelector != nil {
		in, out := &in.ZoneIDSelector, &out.ZoneIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecordInitParameters.
func (in *RecordInitParameters) DeepCopy() *RecordInitParameters {
	if in == nil {
		return nil
	}
	out := new(RecordInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecordList) DeepCopyInto(out *RecordList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Record, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecordList.
func (in *RecordList) DeepCopy() *RecordList {
	if in == nil {
		return nil
	}
	out := new(RecordList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RecordList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecordObservation) DeepCopyInto(out *RecordObservation) {
	*out = *in
	if in.AllowOverwrite != nil {
		in, out := &in.AllowOverwrite, &out.AllowOverwrite
		*out = new(bool)
		**out = **in
	}
	if in.Comment != nil {
		in, out := &in.Comment, &out.Comment
		*out = new(string)
		**out = **in
	}
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = new(string)
		**out = **in
	}
	if in.CreatedOn != nil {
		in, out := &in.CreatedOn, &out.CreatedOn
		*out = new(string)
		**out = **in
	}
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make([]DataObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Hostname != nil {
		in, out := &in.Hostname, &out.Hostname
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.ModifiedOn != nil {
		in, out := &in.ModifiedOn, &out.ModifiedOn
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(float64)
		**out = **in
	}
	if in.Proxiable != nil {
		in, out := &in.Proxiable, &out.Proxiable
		*out = new(bool)
		**out = **in
	}
	if in.Proxied != nil {
		in, out := &in.Proxied, &out.Proxied
		*out = new(bool)
		**out = **in
	}
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(float64)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	if in.ZoneID != nil {
		in, out := &in.ZoneID, &out.ZoneID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecordObservation.
func (in *RecordObservation) DeepCopy() *RecordObservation {
	if in == nil {
		return nil
	}
	out := new(RecordObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecordParameters) DeepCopyInto(out *RecordParameters) {
	*out = *in
	if in.AllowOverwrite != nil {
		in, out := &in.AllowOverwrite, &out.AllowOverwrite
		*out = new(bool)
		**out = **in
	}
	if in.Comment != nil {
		in, out := &in.Comment, &out.Comment
		*out = new(string)
		**out = **in
	}
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = new(string)
		**out = **in
	}
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make([]DataParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(float64)
		**out = **in
	}
	if in.Proxied != nil {
		in, out := &in.Proxied, &out.Proxied
		*out = new(bool)
		**out = **in
	}
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(float64)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	if in.ZoneID != nil {
		in, out := &in.ZoneID, &out.ZoneID
		*out = new(string)
		**out = **in
	}
	if in.ZoneIDSelector != nil {
		in, out := &in.ZoneIDSelector, &out.ZoneIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecordParameters.
func (in *RecordParameters) DeepCopy() *RecordParameters {
	if in == nil {
		return nil
	}
	out := new(RecordParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecordSpec) DeepCopyInto(out *RecordSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecordSpec.
func (in *RecordSpec) DeepCopy() *RecordSpec {
	if in == nil {
		return nil
	}
	out := new(RecordSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecordStatus) DeepCopyInto(out *RecordStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecordStatus.
func (in *RecordStatus) DeepCopy() *RecordStatus {
	if in == nil {
		return nil
	}
	out := new(RecordStatus)
	in.DeepCopyInto(out)
	return out
}
