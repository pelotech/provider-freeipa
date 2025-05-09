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

type RuleRunasuserMembershipInitParameters struct {

	// (String) Unique identifier to differentiate multiple sudo rule runasuser membership resources on the same sudo rule. Manadatory for using runasusers configurations.
	// Unique identifier to differentiate multiple sudo rule runasuser membership resources on the same sudo rule. Manadatory for using runasusers configurations.
	Identifier *string `json:"identifier,omitempty" tf:"identifier,omitempty"`

	// (String) Sudo rule name
	// Sudo rule name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String, Deprecated) deprecated Run As User to add to the sudo rule. Can be an external user (local user of ipa clients)
	// **deprecated** Run As User to add to the sudo rule. Can be an external user (local user of ipa clients)
	Runasuser *string `json:"runasuser,omitempty" tf:"runasuser,omitempty"`

	// (List of String) List of Run As User to add to the sudo rule. Can be an external user (local user of ipa clients)
	// List of Run As User to add to the sudo rule. Can be an external user (local user of ipa clients)
	Runasusers []*string `json:"runasusers,omitempty" tf:"runasusers,omitempty"`
}

type RuleRunasuserMembershipObservation struct {

	// (String) ID of the resource
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Unique identifier to differentiate multiple sudo rule runasuser membership resources on the same sudo rule. Manadatory for using runasusers configurations.
	// Unique identifier to differentiate multiple sudo rule runasuser membership resources on the same sudo rule. Manadatory for using runasusers configurations.
	Identifier *string `json:"identifier,omitempty" tf:"identifier,omitempty"`

	// (String) Sudo rule name
	// Sudo rule name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String, Deprecated) deprecated Run As User to add to the sudo rule. Can be an external user (local user of ipa clients)
	// **deprecated** Run As User to add to the sudo rule. Can be an external user (local user of ipa clients)
	Runasuser *string `json:"runasuser,omitempty" tf:"runasuser,omitempty"`

	// (List of String) List of Run As User to add to the sudo rule. Can be an external user (local user of ipa clients)
	// List of Run As User to add to the sudo rule. Can be an external user (local user of ipa clients)
	Runasusers []*string `json:"runasusers,omitempty" tf:"runasusers,omitempty"`
}

type RuleRunasuserMembershipParameters struct {

	// (String) Unique identifier to differentiate multiple sudo rule runasuser membership resources on the same sudo rule. Manadatory for using runasusers configurations.
	// Unique identifier to differentiate multiple sudo rule runasuser membership resources on the same sudo rule. Manadatory for using runasusers configurations.
	// +kubebuilder:validation:Optional
	Identifier *string `json:"identifier,omitempty" tf:"identifier,omitempty"`

	// (String) Sudo rule name
	// Sudo rule name
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String, Deprecated) deprecated Run As User to add to the sudo rule. Can be an external user (local user of ipa clients)
	// **deprecated** Run As User to add to the sudo rule. Can be an external user (local user of ipa clients)
	// +kubebuilder:validation:Optional
	Runasuser *string `json:"runasuser,omitempty" tf:"runasuser,omitempty"`

	// (List of String) List of Run As User to add to the sudo rule. Can be an external user (local user of ipa clients)
	// List of Run As User to add to the sudo rule. Can be an external user (local user of ipa clients)
	// +kubebuilder:validation:Optional
	Runasusers []*string `json:"runasusers,omitempty" tf:"runasusers,omitempty"`
}

// RuleRunasuserMembershipSpec defines the desired state of RuleRunasuserMembership
type RuleRunasuserMembershipSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RuleRunasuserMembershipParameters `json:"forProvider"`
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
	InitProvider RuleRunasuserMembershipInitParameters `json:"initProvider,omitempty"`
}

// RuleRunasuserMembershipStatus defines the observed state of RuleRunasuserMembership.
type RuleRunasuserMembershipStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RuleRunasuserMembershipObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RuleRunasuserMembership is the Schema for the RuleRunasuserMemberships API. FreeIPA Sudo rule run as user membership resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,freeipa}
type RuleRunasuserMembership struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   RuleRunasuserMembershipSpec   `json:"spec"`
	Status RuleRunasuserMembershipStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RuleRunasuserMembershipList contains a list of RuleRunasuserMemberships
type RuleRunasuserMembershipList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RuleRunasuserMembership `json:"items"`
}

// Repository type metadata.
var (
	RuleRunasuserMembership_Kind             = "RuleRunasuserMembership"
	RuleRunasuserMembership_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RuleRunasuserMembership_Kind}.String()
	RuleRunasuserMembership_KindAPIVersion   = RuleRunasuserMembership_Kind + "." + CRDGroupVersion.String()
	RuleRunasuserMembership_GroupVersionKind = CRDGroupVersion.WithKind(RuleRunasuserMembership_Kind)
)

func init() {
	SchemeBuilder.Register(&RuleRunasuserMembership{}, &RuleRunasuserMembershipList{})
}
