// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type PolicyInitParameters struct {

	// Should the specified ROLE have CREATE privileges to the specified SCHEMA.
	// If true, allow the specified ROLEs to CREATE new objects within the schema(s)
	Create *bool `json:"create,omitempty" tf:"create,omitempty"`

	// Should the specified ROLE have CREATE privileges to the specified SCHEMA and the ability to GRANT the CREATE privilege to other ROLEs.
	// If true, allow the specified ROLEs to CREATE new objects within the schema(s) and GRANT the same CREATE privilege to different ROLEs
	CreateWithGrant *bool `json:"createWithGrant,omitempty" tf:"create_with_grant,omitempty"`

	// The ROLE who is receiving the policy.  If this value is empty or not specified it implies the policy is referring to the PUBLIC role.
	// ROLE who will receive this policy (default: PUBLIC)
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// Should the specified ROLE have USAGE privileges to the specified SCHEMA.
	// If true, allow the specified ROLEs to use objects within the schema(s)
	Usage *bool `json:"usage,omitempty" tf:"usage,omitempty"`

	// Should the specified ROLE have USAGE privileges to the specified SCHEMA and the ability to GRANT the USAGE privilege to other ROLEs.
	// If true, allow the specified ROLEs to use objects within the schema(s) and GRANT the same USAGE privilege to different ROLEs
	UsageWithGrant *bool `json:"usageWithGrant,omitempty" tf:"usage_with_grant,omitempty"`
}

type PolicyObservation struct {

	// Should the specified ROLE have CREATE privileges to the specified SCHEMA.
	// If true, allow the specified ROLEs to CREATE new objects within the schema(s)
	Create *bool `json:"create,omitempty" tf:"create,omitempty"`

	// Should the specified ROLE have CREATE privileges to the specified SCHEMA and the ability to GRANT the CREATE privilege to other ROLEs.
	// If true, allow the specified ROLEs to CREATE new objects within the schema(s) and GRANT the same CREATE privilege to different ROLEs
	CreateWithGrant *bool `json:"createWithGrant,omitempty" tf:"create_with_grant,omitempty"`

	// The ROLE who is receiving the policy.  If this value is empty or not specified it implies the policy is referring to the PUBLIC role.
	// ROLE who will receive this policy (default: PUBLIC)
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// Should the specified ROLE have USAGE privileges to the specified SCHEMA.
	// If true, allow the specified ROLEs to use objects within the schema(s)
	Usage *bool `json:"usage,omitempty" tf:"usage,omitempty"`

	// Should the specified ROLE have USAGE privileges to the specified SCHEMA and the ability to GRANT the USAGE privilege to other ROLEs.
	// If true, allow the specified ROLEs to use objects within the schema(s) and GRANT the same USAGE privilege to different ROLEs
	UsageWithGrant *bool `json:"usageWithGrant,omitempty" tf:"usage_with_grant,omitempty"`
}

type PolicyParameters struct {

	// Should the specified ROLE have CREATE privileges to the specified SCHEMA.
	// If true, allow the specified ROLEs to CREATE new objects within the schema(s)
	// +kubebuilder:validation:Optional
	Create *bool `json:"create,omitempty" tf:"create,omitempty"`

	// Should the specified ROLE have CREATE privileges to the specified SCHEMA and the ability to GRANT the CREATE privilege to other ROLEs.
	// If true, allow the specified ROLEs to CREATE new objects within the schema(s) and GRANT the same CREATE privilege to different ROLEs
	// +kubebuilder:validation:Optional
	CreateWithGrant *bool `json:"createWithGrant,omitempty" tf:"create_with_grant,omitempty"`

	// The ROLE who is receiving the policy.  If this value is empty or not specified it implies the policy is referring to the PUBLIC role.
	// ROLE who will receive this policy (default: PUBLIC)
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// Should the specified ROLE have USAGE privileges to the specified SCHEMA.
	// If true, allow the specified ROLEs to use objects within the schema(s)
	// +kubebuilder:validation:Optional
	Usage *bool `json:"usage,omitempty" tf:"usage,omitempty"`

	// Should the specified ROLE have USAGE privileges to the specified SCHEMA and the ability to GRANT the USAGE privilege to other ROLEs.
	// If true, allow the specified ROLEs to use objects within the schema(s) and GRANT the same USAGE privilege to different ROLEs
	// +kubebuilder:validation:Optional
	UsageWithGrant *bool `json:"usageWithGrant,omitempty" tf:"usage_with_grant,omitempty"`
}

type SchemaInitParameters struct {

	// The DATABASE in which where this schema will be created. (Default: The database used by your provider configuration)
	// The database name to alter schema
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	// When true, will also drop all the objects that are contained in the schema. (Default: false)
	// When true, will also drop all the objects that are contained in the schema
	DropCascade *bool `json:"dropCascade,omitempty" tf:"drop_cascade,omitempty"`

	// When true, use the existing schema if it exists. (Default: true)
	// When true, use the existing schema if it exists
	IfNotExists *bool `json:"ifNotExists,omitempty" tf:"if_not_exists,omitempty"`

	// The name of the schema. Must be unique in the PostgreSQL
	// database instance where it is configured.
	// The name of the schema
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ROLE who owns the schema.
	// The ROLE name who owns the schema
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Can be specified multiple times for each policy.  Each
	// policy block supports fields documented below.
	Policy []PolicyInitParameters `json:"policy,omitempty" tf:"policy,omitempty"`
}

type SchemaObservation struct {

	// The DATABASE in which where this schema will be created. (Default: The database used by your provider configuration)
	// The database name to alter schema
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	// When true, will also drop all the objects that are contained in the schema. (Default: false)
	// When true, will also drop all the objects that are contained in the schema
	DropCascade *bool `json:"dropCascade,omitempty" tf:"drop_cascade,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// When true, use the existing schema if it exists. (Default: true)
	// When true, use the existing schema if it exists
	IfNotExists *bool `json:"ifNotExists,omitempty" tf:"if_not_exists,omitempty"`

	// The name of the schema. Must be unique in the PostgreSQL
	// database instance where it is configured.
	// The name of the schema
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ROLE who owns the schema.
	// The ROLE name who owns the schema
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Can be specified multiple times for each policy.  Each
	// policy block supports fields documented below.
	Policy []PolicyObservation `json:"policy,omitempty" tf:"policy,omitempty"`
}

type SchemaParameters struct {

	// The DATABASE in which where this schema will be created. (Default: The database used by your provider configuration)
	// The database name to alter schema
	// +kubebuilder:validation:Optional
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	// When true, will also drop all the objects that are contained in the schema. (Default: false)
	// When true, will also drop all the objects that are contained in the schema
	// +kubebuilder:validation:Optional
	DropCascade *bool `json:"dropCascade,omitempty" tf:"drop_cascade,omitempty"`

	// When true, use the existing schema if it exists. (Default: true)
	// When true, use the existing schema if it exists
	// +kubebuilder:validation:Optional
	IfNotExists *bool `json:"ifNotExists,omitempty" tf:"if_not_exists,omitempty"`

	// The name of the schema. Must be unique in the PostgreSQL
	// database instance where it is configured.
	// The name of the schema
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ROLE who owns the schema.
	// The ROLE name who owns the schema
	// +kubebuilder:validation:Optional
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Can be specified multiple times for each policy.  Each
	// policy block supports fields documented below.
	// +kubebuilder:validation:Optional
	Policy []PolicyParameters `json:"policy,omitempty" tf:"policy,omitempty"`
}

// SchemaSpec defines the desired state of Schema
type SchemaSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SchemaParameters `json:"forProvider"`
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
	InitProvider SchemaInitParameters `json:"initProvider,omitempty"`
}

// SchemaStatus defines the observed state of Schema.
type SchemaStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SchemaObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Schema is the Schema for the Schemas API. Creates and manages a schema within a PostgreSQL database.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,postgresql}
type Schema struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   SchemaSpec   `json:"spec"`
	Status SchemaStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SchemaList contains a list of Schemas
type SchemaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Schema `json:"items"`
}

// Repository type metadata.
var (
	Schema_Kind             = "Schema"
	Schema_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Schema_Kind}.String()
	Schema_KindAPIVersion   = Schema_Kind + "." + CRDGroupVersion.String()
	Schema_GroupVersionKind = CRDGroupVersion.WithKind(Schema_Kind)
)

func init() {
	SchemeBuilder.Register(&Schema{}, &SchemaList{})
}