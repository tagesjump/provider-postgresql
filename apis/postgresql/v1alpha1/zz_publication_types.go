// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type PublicationInitParameters struct {

	// Should be ALL TABLES added to the publication. Defaults to 'false'
	// Sets the tables list to publish to ALL tables
	AllTables *bool `json:"allTables,omitempty" tf:"all_tables,omitempty"`

	// Which database to create the publication on. Defaults to provider database.
	// Sets the database to add the publication for
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	// Should all subsequent resources of the publication be dropped. Defaults to 'false'
	// When true, will also drop all the objects that depend on the publication, and in turn all objects that depend on those objects
	DropCascade *bool `json:"dropCascade,omitempty" tf:"drop_cascade,omitempty"`

	// The name of the publication.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Who owns the publication. Defaults to provider user.
	// Sets the owner of the publication
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Which 'publish' options should be turned on. Default to 'insert','update','delete'
	// Sets which DML operations will be published
	PublishParam []*string `json:"publishParam,omitempty" tf:"publish_param,omitempty"`

	// Should be option 'publish_via_partition_root' be turned on. Default to 'false'
	// Sets whether changes in a partitioned table using the identity and schema of the partitioned table
	PublishViaPartitionRootParam *bool `json:"publishViaPartitionRootParam,omitempty" tf:"publish_via_partition_root_param,omitempty"`

	// Which tables add to the publication. By defaults no tables added. Format of table is <schema_name>.<table_name>. If <schema_name> is not specified - default database schema will be used.  Table string must be listed in alphabetical order.
	// Sets the tables list to publish
	// +listType=set
	Tables []*string `json:"tables,omitempty" tf:"tables,omitempty"`
}

type PublicationObservation struct {

	// Should be ALL TABLES added to the publication. Defaults to 'false'
	// Sets the tables list to publish to ALL tables
	AllTables *bool `json:"allTables,omitempty" tf:"all_tables,omitempty"`

	// Which database to create the publication on. Defaults to provider database.
	// Sets the database to add the publication for
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	// Should all subsequent resources of the publication be dropped. Defaults to 'false'
	// When true, will also drop all the objects that depend on the publication, and in turn all objects that depend on those objects
	DropCascade *bool `json:"dropCascade,omitempty" tf:"drop_cascade,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the publication.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Who owns the publication. Defaults to provider user.
	// Sets the owner of the publication
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Which 'publish' options should be turned on. Default to 'insert','update','delete'
	// Sets which DML operations will be published
	PublishParam []*string `json:"publishParam,omitempty" tf:"publish_param,omitempty"`

	// Should be option 'publish_via_partition_root' be turned on. Default to 'false'
	// Sets whether changes in a partitioned table using the identity and schema of the partitioned table
	PublishViaPartitionRootParam *bool `json:"publishViaPartitionRootParam,omitempty" tf:"publish_via_partition_root_param,omitempty"`

	// Which tables add to the publication. By defaults no tables added. Format of table is <schema_name>.<table_name>. If <schema_name> is not specified - default database schema will be used.  Table string must be listed in alphabetical order.
	// Sets the tables list to publish
	// +listType=set
	Tables []*string `json:"tables,omitempty" tf:"tables,omitempty"`
}

type PublicationParameters struct {

	// Should be ALL TABLES added to the publication. Defaults to 'false'
	// Sets the tables list to publish to ALL tables
	// +kubebuilder:validation:Optional
	AllTables *bool `json:"allTables,omitempty" tf:"all_tables,omitempty"`

	// Which database to create the publication on. Defaults to provider database.
	// Sets the database to add the publication for
	// +kubebuilder:validation:Optional
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	// Should all subsequent resources of the publication be dropped. Defaults to 'false'
	// When true, will also drop all the objects that depend on the publication, and in turn all objects that depend on those objects
	// +kubebuilder:validation:Optional
	DropCascade *bool `json:"dropCascade,omitempty" tf:"drop_cascade,omitempty"`

	// The name of the publication.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Who owns the publication. Defaults to provider user.
	// Sets the owner of the publication
	// +kubebuilder:validation:Optional
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Which 'publish' options should be turned on. Default to 'insert','update','delete'
	// Sets which DML operations will be published
	// +kubebuilder:validation:Optional
	PublishParam []*string `json:"publishParam,omitempty" tf:"publish_param,omitempty"`

	// Should be option 'publish_via_partition_root' be turned on. Default to 'false'
	// Sets whether changes in a partitioned table using the identity and schema of the partitioned table
	// +kubebuilder:validation:Optional
	PublishViaPartitionRootParam *bool `json:"publishViaPartitionRootParam,omitempty" tf:"publish_via_partition_root_param,omitempty"`

	// Which tables add to the publication. By defaults no tables added. Format of table is <schema_name>.<table_name>. If <schema_name> is not specified - default database schema will be used.  Table string must be listed in alphabetical order.
	// Sets the tables list to publish
	// +kubebuilder:validation:Optional
	// +listType=set
	Tables []*string `json:"tables,omitempty" tf:"tables,omitempty"`
}

// PublicationSpec defines the desired state of Publication
type PublicationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PublicationParameters `json:"forProvider"`
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
	InitProvider PublicationInitParameters `json:"initProvider,omitempty"`
}

// PublicationStatus defines the observed state of Publication.
type PublicationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PublicationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Publication is the Schema for the Publications API. Creates and manages a publication in a PostgreSQL server database.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,postgresql}
type Publication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   PublicationSpec   `json:"spec"`
	Status PublicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PublicationList contains a list of Publications
type PublicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Publication `json:"items"`
}

// Repository type metadata.
var (
	Publication_Kind             = "Publication"
	Publication_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Publication_Kind}.String()
	Publication_KindAPIVersion   = Publication_Kind + "." + CRDGroupVersion.String()
	Publication_GroupVersionKind = CRDGroupVersion.WithKind(Publication_Kind)
)

func init() {
	SchemeBuilder.Register(&Publication{}, &PublicationList{})
}