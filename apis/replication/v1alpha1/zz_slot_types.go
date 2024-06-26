// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SlotInitParameters struct {

	// Which database to create the replication slot on. Defaults to provider database.
	// Sets the database to add the replication slot to
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	// The name of the replication slot.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Sets the output plugin.
	// Sets the output plugin to use
	Plugin *string `json:"plugin,omitempty" tf:"plugin,omitempty"`
}

type SlotObservation struct {

	// Which database to create the replication slot on. Defaults to provider database.
	// Sets the database to add the replication slot to
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the replication slot.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Sets the output plugin.
	// Sets the output plugin to use
	Plugin *string `json:"plugin,omitempty" tf:"plugin,omitempty"`
}

type SlotParameters struct {

	// Which database to create the replication slot on. Defaults to provider database.
	// Sets the database to add the replication slot to
	// +kubebuilder:validation:Optional
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	// The name of the replication slot.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Sets the output plugin.
	// Sets the output plugin to use
	// +kubebuilder:validation:Optional
	Plugin *string `json:"plugin,omitempty" tf:"plugin,omitempty"`
}

// SlotSpec defines the desired state of Slot
type SlotSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SlotParameters `json:"forProvider"`
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
	InitProvider SlotInitParameters `json:"initProvider,omitempty"`
}

// SlotStatus defines the observed state of Slot.
type SlotStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SlotObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Slot is the Schema for the Slots API. Creates and manages a replication slot on a PostgreSQL server.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,postgresql}
type Slot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.plugin) || (has(self.initProvider) && has(self.initProvider.plugin))",message="spec.forProvider.plugin is a required parameter"
	Spec   SlotSpec   `json:"spec"`
	Status SlotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SlotList contains a list of Slots
type SlotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Slot `json:"items"`
}

// Repository type metadata.
var (
	Slot_Kind             = "Slot"
	Slot_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Slot_Kind}.String()
	Slot_KindAPIVersion   = Slot_Kind + "." + CRDGroupVersion.String()
	Slot_GroupVersionKind = CRDGroupVersion.WithKind(Slot_Kind)
)

func init() {
	SchemeBuilder.Register(&Slot{}, &SlotList{})
}
