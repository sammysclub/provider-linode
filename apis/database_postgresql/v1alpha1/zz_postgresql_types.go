/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type PostgresqlObservation struct {

	// When this Managed Database was created.
	// When this Managed Database was created.
	Created *string `json:"created,omitempty" tf:"created,omitempty"`

	// The Managed Database engine. (e.g. postgresql)
	// The Managed Database engine.
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// The primary host for the Managed Database.
	// The primary host for the Managed Database.
	HostPrimary *string `json:"hostPrimary,omitempty" tf:"host_primary,omitempty"`

	// The secondary/private network host for the Managed Database.
	// The secondary host for the Managed Database.
	HostSecondary *string `json:"hostSecondary,omitempty" tf:"host_secondary,omitempty"`

	// The ID of the Managed Database.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The access port for this Managed Database.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The operating status of the Managed Database.
	// The operating status of the Managed Database.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// When this Managed Database was last updated.
	// When this Managed Database was last updated.
	Updated *string `json:"updated,omitempty" tf:"updated,omitempty"`

	// The Managed Database engine version. (e.g. 13.2)
	// The Managed Database engine version.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type PostgresqlParameters struct {

	// A list of IP addresses that can access the Managed Database. Each item can be a single IP address or a range in CIDR format. Use linode_database_access_controls to manage your allow list separately.
	// A list of IP addresses that can access the Managed Database. Each item can be a single IP address or a range in CIDR format.
	// +kubebuilder:validation:Optional
	AllowList []*string `json:"allowList,omitempty" tf:"allow_list,omitempty"`

	// The number of Linode Instance nodes deployed to the Managed Database. (default 1)
	// The number of Linode Instance nodes deployed to the Managed Database. Defaults to 1.
	// +kubebuilder:validation:Optional
	ClusterSize *float64 `json:"clusterSize,omitempty" tf:"cluster_size,omitempty"`

	// Whether the Managed Databases is encrypted. (default false)
	// Whether the Managed Databases is encrypted.
	// +kubebuilder:validation:Optional
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// The Managed Database engine in engine/version format. (e.g. postgresql/13.2)
	// The Managed Database engine in engine/version format. (e.g. mongodb/4.4.10)
	// +kubebuilder:validation:Required
	EngineID *string `json:"engineId" tf:"engine_id,omitempty"`

	// A unique, user-defined string referring to the Managed Database.
	// A unique, user-defined string referring to the Managed Database.
	// +kubebuilder:validation:Required
	Label *string `json:"label" tf:"label,omitempty"`

	// The region to use for the Managed Database.
	// The region to use for the Managed Database.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// The synchronization level of the replicating server. (on, local, remote_write, remote_apply, off; default off)
	// The synchronization level of the replicating server.Must be `local` or `off` for the `asynch` replication type. Must be `on`, `remote_write`, or `remote_apply` for the `semi_synch` replication type.
	// +kubebuilder:validation:Optional
	ReplicationCommitType *string `json:"replicationCommitType,omitempty" tf:"replication_commit_type,omitempty"`

	// The replication method used for the Managed Database. (none, asynch, semi_synch; default none)
	// The replication method used for the Managed Database. Must be `none` for a single node cluster. Must be `asynch` or `semi_synch` for a high availability cluster.
	// +kubebuilder:validation:Optional
	ReplicationType *string `json:"replicationType,omitempty" tf:"replication_type,omitempty"`

	// Whether to require SSL credentials to establish a connection to the Managed Database. (default false)
	// Whether to require SSL credentials to establish a connection to the Managed Database.
	// +kubebuilder:validation:Optional
	SSLConnection *bool `json:"sslConnection,omitempty" tf:"ssl_connection,omitempty"`

	// The Linode Instance type used for the nodes of the  Managed Database instance.
	// The Linode Instance type used by the Managed Database for its nodes.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// Configuration settings for automated patch update maintenance for the Managed Database.
	// +kubebuilder:validation:Optional
	Updates []UpdatesParameters `json:"updates,omitempty" tf:"updates,omitempty"`
}

type UpdatesObservation struct {
}

type UpdatesParameters struct {

	// The day to perform maintenance. (monday, tuesday, ...)
	// The day to perform maintenance.
	// +kubebuilder:validation:Required
	DayOfWeek *string `json:"dayOfWeek" tf:"day_of_week,omitempty"`

	// The maximum maintenance window time in hours. (1..3)
	// The maximum maintenance window time in hours.
	// +kubebuilder:validation:Required
	Duration *float64 `json:"duration" tf:"duration,omitempty"`

	// Whether maintenance occurs on a weekly or monthly basis. (weekly, monthly)
	// Whether maintenance occurs on a weekly or monthly basis.
	// +kubebuilder:validation:Required
	Frequency *string `json:"frequency" tf:"frequency,omitempty"`

	// The hour to begin maintenance based in UTC time. (0..23)
	// The hour to begin maintenance based in UTC time.
	// +kubebuilder:validation:Required
	HourOfDay *float64 `json:"hourOfDay" tf:"hour_of_day,omitempty"`

	// The week of the month to perform monthly frequency updates. Required for monthly frequency updates. (1..4)
	// The week of the month to perform monthly frequency updates. Required for monthly frequency updates.
	// +kubebuilder:validation:Optional
	WeekOfMonth *float64 `json:"weekOfMonth,omitempty" tf:"week_of_month,omitempty"`
}

// PostgresqlSpec defines the desired state of Postgresql
type PostgresqlSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PostgresqlParameters `json:"forProvider"`
}

// PostgresqlStatus defines the observed state of Postgresql.
type PostgresqlStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PostgresqlObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Postgresql is the Schema for the Postgresqls API. Manages a Linode PostgreSQL Database.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,linode}
type Postgresql struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PostgresqlSpec   `json:"spec"`
	Status            PostgresqlStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PostgresqlList contains a list of Postgresqls
type PostgresqlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Postgresql `json:"items"`
}

// Repository type metadata.
var (
	Postgresql_Kind             = "Postgresql"
	Postgresql_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Postgresql_Kind}.String()
	Postgresql_KindAPIVersion   = Postgresql_Kind + "." + CRDGroupVersion.String()
	Postgresql_GroupVersionKind = CRDGroupVersion.WithKind(Postgresql_Kind)
)

func init() {
	SchemeBuilder.Register(&Postgresql{}, &PostgresqlList{})
}
