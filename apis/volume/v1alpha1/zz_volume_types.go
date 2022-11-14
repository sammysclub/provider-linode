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

type VolumeObservation struct {

	// The full filesystem path for the Volume based on the Volume's label. The path is "/dev/disk/by-id/scsi-0Linode_Volume_" + the Volume label
	// The full filesystem path for the Volume based on the Volume's label. Path is /dev/disk/by-id/scsi-0Linode_Volume_ + Volume label.
	FilesystemPath *string `json:"filesystemPath,omitempty" tf:"filesystem_path,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The status of the Linode Volume. (creating, active, resizing, contact_support)
	// The status of the volume, indicating the current readiness state.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type VolumeParameters struct {

	// The label of the Linode Volume
	// The label of the Linode Volume.
	// +kubebuilder:validation:Required
	Label *string `json:"label" tf:"label,omitempty"`

	// The ID of a Linode Instance where the Volume should be attached.
	// The Linode ID where the Volume should be attached.
	// +kubebuilder:validation:Optional
	LinodeID *float64 `json:"linodeId,omitempty" tf:"linode_id,omitempty"`

	// The region where this volume will be deployed.  Examples are "us-east", "us-west", "ap-south", etc. See all regions here. This field is optional for cloned volumes. Changing .
	// The region where this volume will be deployed.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Size of the Volume in GB.
	// Size of the Volume in GB
	// +kubebuilder:validation:Optional
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// The ID of a Linode Volume to clone. NOTE: Cloned volumes must be in the same region as the source volume.
	// The ID of a volume to clone.
	// +kubebuilder:validation:Optional
	SourceVolumeID *float64 `json:"sourceVolumeId,omitempty" tf:"source_volume_id,omitempty"`

	// A list of tags applied to this object. Tags are for organizational purposes only.
	// An array of tags applied to this object. Tags are for organizational purposes only.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// VolumeSpec defines the desired state of Volume
type VolumeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VolumeParameters `json:"forProvider"`
}

// VolumeStatus defines the observed state of Volume.
type VolumeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VolumeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Volume is the Schema for the Volumes API. Manages a Linode Volume.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,linode}
type Volume struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VolumeSpec   `json:"spec"`
	Status            VolumeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VolumeList contains a list of Volumes
type VolumeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Volume `json:"items"`
}

// Repository type metadata.
var (
	Volume_Kind             = "Volume"
	Volume_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Volume_Kind}.String()
	Volume_KindAPIVersion   = Volume_Kind + "." + CRDGroupVersion.String()
	Volume_GroupVersionKind = CRDGroupVersion.WithKind(Volume_Kind)
)

func init() {
	SchemeBuilder.Register(&Volume{}, &VolumeList{})
}
