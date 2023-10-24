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

type IPObservation struct {

	// The resulting IPv4 address.
	// The resulting IPv4 address.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// If true, the instance will be rebooted to update network interfaces.
	// If true, the instance will be rebooted to update network interfaces. This functionality is not affected by the `skip_implicit_reboots` provider argument.
	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`

	// The default gateway for this address
	// The default gateway for this address
	Gateway *string `json:"gateway,omitempty" tf:"gateway,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the Linode to allocate an IPv4 address for.
	// The ID of the Linode to allocate an IPv4 address for.
	LinodeID *float64 `json:"linodeId,omitempty" tf:"linode_id,omitempty"`

	// The number of bits set in the subnet mask.
	// The number of bits set in the subnet mask.
	Prefix *float64 `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Whether the IPv4 address is public or private. Defaults to true.
	// Whether the IPv4 address is public or private.
	Public *bool `json:"public,omitempty" tf:"public,omitempty"`

	// The reverse DNS assigned to this address.
	// The reverse DNS assigned to this address.
	Rdns *string `json:"rdns,omitempty" tf:"rdns,omitempty"`

	// The region this IP resides in.
	// The region this IP resides in.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The mask that separates host bits from network bits for this address.
	// The mask that separates host bits from network bits for this address.
	SubnetMask *string `json:"subnetMask,omitempty" tf:"subnet_mask,omitempty"`

	// The type of IP address. (ipv4, ipv6, ipv6/pool, ipv6/range)
	// The type of IP address.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IPParameters struct {

	// If true, the instance will be rebooted to update network interfaces.
	// If true, the instance will be rebooted to update network interfaces. This functionality is not affected by the `skip_implicit_reboots` provider argument.
	// +kubebuilder:validation:Optional
	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`

	// The ID of the Linode to allocate an IPv4 address for.
	// The ID of the Linode to allocate an IPv4 address for.
	// +crossplane:generate:reference:type=Instance
	// +kubebuilder:validation:Optional
	LinodeID *float64 `json:"linodeId,omitempty" tf:"linode_id,omitempty"`

	// Reference to a Instance to populate linodeId.
	// +kubebuilder:validation:Optional
	LinodeIDRef *v1.Reference `json:"linodeIdRef,omitempty" tf:"-"`

	// Selector for a Instance to populate linodeId.
	// +kubebuilder:validation:Optional
	LinodeIDSelector *v1.Selector `json:"linodeIdSelector,omitempty" tf:"-"`

	// Whether the IPv4 address is public or private. Defaults to true.
	// Whether the IPv4 address is public or private.
	// +kubebuilder:validation:Optional
	Public *bool `json:"public,omitempty" tf:"public,omitempty"`

	// The reverse DNS assigned to this address.
	// The reverse DNS assigned to this address.
	// +crossplane:generate:reference:type=github.com/linode/provider-linode/apis/rdns/v1alpha1.RDNS
	// +kubebuilder:validation:Optional
	Rdns *string `json:"rdns,omitempty" tf:"rdns,omitempty"`

	// Reference to a RDNS in rdns to populate rdns.
	// +kubebuilder:validation:Optional
	RdnsRef *v1.Reference `json:"rdnsRef,omitempty" tf:"-"`

	// Selector for a RDNS in rdns to populate rdns.
	// +kubebuilder:validation:Optional
	RdnsSelector *v1.Selector `json:"rdnsSelector,omitempty" tf:"-"`
}

// IPSpec defines the desired state of IP
type IPSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IPParameters `json:"forProvider"`
}

// IPStatus defines the observed state of IP.
type IPStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IPObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IP is the Schema for the IPs API. Manages a Linode instance IP.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,linode}
type IP struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IPSpec   `json:"spec"`
	Status            IPStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IPList contains a list of IPs
type IPList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IP `json:"items"`
}

// Repository type metadata.
var (
	IP_Kind             = "IP"
	IP_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IP_Kind}.String()
	IP_KindAPIVersion   = IP_Kind + "." + CRDGroupVersion.String()
	IP_GroupVersionKind = CRDGroupVersion.WithKind(IP_Kind)
)

func init() {
	SchemeBuilder.Register(&IP{}, &IPList{})
}
