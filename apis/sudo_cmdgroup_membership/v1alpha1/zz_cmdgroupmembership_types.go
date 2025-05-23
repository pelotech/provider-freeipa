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

type CmdgroupMembershipInitParameters struct {

	// (String) Unique identifier to differentiate multiple sudo command group membership resources on the same sudo command group. Manadatory for using sudocmds configurations.
	// Unique identifier to differentiate multiple sudo command group membership resources on the same sudo command group. Manadatory for using sudocmds configurations.
	Identifier *string `json:"identifier,omitempty" tf:"identifier,omitempty"`

	// (String) Name of the sudo command group
	// Name of the sudo command group
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String, Deprecated) deprecated Sudo command to add as a member
	// **deprecated** Sudo command to add as a member
	Sudocmd *string `json:"sudocmd,omitempty" tf:"sudocmd,omitempty"`

	// (List of String) List of sudo command to add as a member
	// List of sudo command to add as a member
	Sudocmds []*string `json:"sudocmds,omitempty" tf:"sudocmds,omitempty"`
}

type CmdgroupMembershipObservation struct {

	// (String) ID of the resource
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Unique identifier to differentiate multiple sudo command group membership resources on the same sudo command group. Manadatory for using sudocmds configurations.
	// Unique identifier to differentiate multiple sudo command group membership resources on the same sudo command group. Manadatory for using sudocmds configurations.
	Identifier *string `json:"identifier,omitempty" tf:"identifier,omitempty"`

	// (String) Name of the sudo command group
	// Name of the sudo command group
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String, Deprecated) deprecated Sudo command to add as a member
	// **deprecated** Sudo command to add as a member
	Sudocmd *string `json:"sudocmd,omitempty" tf:"sudocmd,omitempty"`

	// (List of String) List of sudo command to add as a member
	// List of sudo command to add as a member
	Sudocmds []*string `json:"sudocmds,omitempty" tf:"sudocmds,omitempty"`
}

type CmdgroupMembershipParameters struct {

	// (String) Unique identifier to differentiate multiple sudo command group membership resources on the same sudo command group. Manadatory for using sudocmds configurations.
	// Unique identifier to differentiate multiple sudo command group membership resources on the same sudo command group. Manadatory for using sudocmds configurations.
	// +kubebuilder:validation:Optional
	Identifier *string `json:"identifier,omitempty" tf:"identifier,omitempty"`

	// (String) Name of the sudo command group
	// Name of the sudo command group
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String, Deprecated) deprecated Sudo command to add as a member
	// **deprecated** Sudo command to add as a member
	// +kubebuilder:validation:Optional
	Sudocmd *string `json:"sudocmd,omitempty" tf:"sudocmd,omitempty"`

	// (List of String) List of sudo command to add as a member
	// List of sudo command to add as a member
	// +kubebuilder:validation:Optional
	Sudocmds []*string `json:"sudocmds,omitempty" tf:"sudocmds,omitempty"`
}

// CmdgroupMembershipSpec defines the desired state of CmdgroupMembership
type CmdgroupMembershipSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CmdgroupMembershipParameters `json:"forProvider"`
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
	InitProvider CmdgroupMembershipInitParameters `json:"initProvider,omitempty"`
}

// CmdgroupMembershipStatus defines the observed state of CmdgroupMembership.
type CmdgroupMembershipStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CmdgroupMembershipObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// CmdgroupMembership is the Schema for the CmdgroupMemberships API. FreeIPA Sudo command group membership resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,freeipa}
type CmdgroupMembership struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   CmdgroupMembershipSpec   `json:"spec"`
	Status CmdgroupMembershipStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CmdgroupMembershipList contains a list of CmdgroupMemberships
type CmdgroupMembershipList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CmdgroupMembership `json:"items"`
}

// Repository type metadata.
var (
	CmdgroupMembership_Kind             = "CmdgroupMembership"
	CmdgroupMembership_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CmdgroupMembership_Kind}.String()
	CmdgroupMembership_KindAPIVersion   = CmdgroupMembership_Kind + "." + CRDGroupVersion.String()
	CmdgroupMembership_GroupVersionKind = CRDGroupVersion.WithKind(CmdgroupMembership_Kind)
)

func init() {
	SchemeBuilder.Register(&CmdgroupMembership{}, &CmdgroupMembershipList{})
}
