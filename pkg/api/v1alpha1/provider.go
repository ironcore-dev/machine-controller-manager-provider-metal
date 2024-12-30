// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"net/netip"

	corev1 "k8s.io/api/core/v1"
)

const (
	// V1Alpha1 is the API version
	V1Alpha1 = "mcm.gardener.cloud/v1alpha1"
	// ProviderName is the provider name
	ProviderName = "metal"
)

// ProviderSpec is the spec to be used while parsing the calls
type ProviderSpec struct {
	// Image is the URL pointing to an OCI registry containing the operating system image which should be used to boot the Machine
	Image string `json:"image,omitempty"`
	// Ignition contains the ignition configuration which should be run on first boot of a Machine.
	Ignition string `json:"ignition,omitempty"`
	// By default, if ignition is set it will be merged it with our template
	// If IgnitionOverride is set to true allows to fully override
	IgnitionOverride bool `json:"ignitionOverride,omitempty"`
	// IgnitionSecretKey is optional key field used to identify the ignition content in the Secret
	// If the key is empty, the DefaultIgnitionKey will be used as fallback.
	IgnitionSecretKey string `json:"ignitionSecretKey,omitempty"`
	// Labels are used to tag resources which the MCM creates, so they can be identified later.
	Labels map[string]string `json:"labels,omitempty"`
	// DnsServers is a list of DNS resolvers which should be configured on the host.
	DnsServers []netip.Addr `json:"dnsServers,omitempty"`
	// ServerLabels are passed to the ServerClaim to find a server with certain properties
	ServerLabels map[string]string `json:"serverLabels,omitempty"`
	// MedaData is a key-value map of additional data which should be passed to the Machine.
	MetaData map[string]any `json:"metaData,omitempty"`
	// AddressesFromNetworks is a list of LocalObjectReferences to Network resources that should be used to assign IP addresses to the worker nodes.
	AddressesFromNetworks []AddressesFromNetworks `json:"addressesFromNetworks,omitempty"`
}

type AddressesFromNetworks struct {
	// Key is the name of metadata key for the network.
	Key string `json:"key"`
	// SubnetRef is a reference to the IP subnet.
	SubnetRef corev1.LocalObjectReference `json:"subnetRef"`
}
