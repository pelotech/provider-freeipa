// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ConditionInitParameters struct {

	// (String) Automember rule condition description
	// Automember rule condition description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (List of String) Regex expression for values that should be excluded.
	// Regex expression for values that should be excluded.
	Exclusiveregex []*string `json:"exclusiveregex,omitempty" tf:"exclusiveregex,omitempty"`

	// (List of String) Regex expression for values that should be included.
	// Regex expression for values that should be included.
	Inclusiveregex []*string `json:"inclusiveregex,omitempty" tf:"inclusiveregex,omitempty"`

	// (String) Automember rule condition key
	// Automember rule condition key
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// (String) Automember rule condition name
	// Automember rule condition name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Automember rule condition type
	// Automember rule condition type
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ConditionObservation struct {

	// (String) Automember rule condition description
	// Automember rule condition description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (List of String) Regex expression for values that should be excluded.
	// Regex expression for values that should be excluded.
	Exclusiveregex []*string `json:"exclusiveregex,omitempty" tf:"exclusiveregex,omitempty"`

	// (String) ID of the resource
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (List of String) Regex expression for values that should be included.
	// Regex expression for values that should be included.
	Inclusiveregex []*string `json:"inclusiveregex,omitempty" tf:"inclusiveregex,omitempty"`

	// (String) Automember rule condition key
	// Automember rule condition key
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// (String) Automember rule condition name
	// Automember rule condition name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Automember rule condition type
	// Automember rule condition type
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ConditionParameters struct {

	// (String) Automember rule condition description
	// Automember rule condition description
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (List of String) Regex expression for values that should be excluded.
	// Regex expression for values that should be excluded.
	// +kubebuilder:validation:Optional
	Exclusiveregex []*string `json:"exclusiveregex,omitempty" tf:"exclusiveregex,omitempty"`

	// (List of String) Regex expression for values that should be included.
	// Regex expression for values that should be included.
	// +kubebuilder:validation:Optional
	Inclusiveregex []*string `json:"inclusiveregex,omitempty" tf:"inclusiveregex,omitempty"`

	// (String) Automember rule condition key
	// Automember rule condition key
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// (String) Automember rule condition name
	// Automember rule condition name
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Automember rule condition type
	// Automember rule condition type
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// ConditionSpec defines the desired state of Condition
type ConditionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConditionParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ConditionInitParameters `json:"initProvider,omitempty"`
}

// ConditionStatus defines the observed state of Condition.
type ConditionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConditionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Condition is the Schema for the Conditions API. FreeIPA Automember conditionresource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,freeipa}
type Condition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.description) || (has(self.initProvider) && has(self.initProvider.description))",message="spec.forProvider.description is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.key) || (has(self.initProvider) && has(self.initProvider.key))",message="spec.forProvider.key is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   ConditionSpec   `json:"spec"`
	Status ConditionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConditionList contains a list of Conditions
type ConditionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Condition `json:"items"`
}

// Repository type metadata.
var (
	Condition_Kind             = "Condition"
	Condition_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Condition_Kind}.String()
	Condition_KindAPIVersion   = Condition_Kind + "." + CRDGroupVersion.String()
	Condition_GroupVersionKind = CRDGroupVersion.WithKind(Condition_Kind)
)

func init() {
	SchemeBuilder.Register(&Condition{}, &ConditionList{})
}
