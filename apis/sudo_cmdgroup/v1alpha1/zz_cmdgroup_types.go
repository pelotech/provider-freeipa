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

type CmdgroupInitParameters struct {

	// (String) Sudo command group description
	// Sudo command group description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) Name of the sudo command group
	// Name of the sudo command group
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type CmdgroupObservation struct {

	// (String) Sudo command group description
	// Sudo command group description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) ID of the resource
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Name of the sudo command group
	// Name of the sudo command group
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type CmdgroupParameters struct {

	// (String) Sudo command group description
	// Sudo command group description
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) Name of the sudo command group
	// Name of the sudo command group
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

// CmdgroupSpec defines the desired state of Cmdgroup
type CmdgroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CmdgroupParameters `json:"forProvider"`
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
	InitProvider CmdgroupInitParameters `json:"initProvider,omitempty"`
}

// CmdgroupStatus defines the observed state of Cmdgroup.
type CmdgroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CmdgroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Cmdgroup is the Schema for the Cmdgroups API. FreeIPA Sudo command group resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,freeipa}
type Cmdgroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   CmdgroupSpec   `json:"spec"`
	Status CmdgroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CmdgroupList contains a list of Cmdgroups
type CmdgroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cmdgroup `json:"items"`
}

// Repository type metadata.
var (
	Cmdgroup_Kind             = "Cmdgroup"
	Cmdgroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Cmdgroup_Kind}.String()
	Cmdgroup_KindAPIVersion   = Cmdgroup_Kind + "." + CRDGroupVersion.String()
	Cmdgroup_GroupVersionKind = CRDGroupVersion.WithKind(Cmdgroup_Kind)
)

func init() {
	SchemeBuilder.Register(&Cmdgroup{}, &CmdgroupList{})
}
