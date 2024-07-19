// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

package validation

import (
	"net/netip"

	"github.com/ironcore-dev/machine-controller-manager-provider-metal/pkg/api/v1alpha1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

var fldPath *field.Path

var _ = Describe("Machine", func() {
	invalidIP := netip.Addr{}

	DescribeTable("ValidateProviderSpecAndSecret",
		func(spec *v1alpha1.ProviderSpec, secret *corev1.Secret, fldPath *field.Path, match types.GomegaMatcher) {
			errList := ValidateProviderSpecAndSecret(spec, secret, fldPath)
			Expect(errList).To(match)
		},
		Entry("no secret",
			&v1alpha1.ProviderSpec{},
			nil,
			fldPath,
			ContainElement(field.Required(fldPath.Child("spec.secretRef"), "secretRef is required")),
		),
		Entry("no userData in secret",
			&v1alpha1.ProviderSpec{},
			&corev1.Secret{
				Data: map[string][]byte{
					"userData": nil,
				},
			},
			fldPath,
			ContainElement(field.Required(fldPath.Child("userData"), "userData is required")),
		),
		Entry("no image",
			&v1alpha1.ProviderSpec{
				Image: "",
			},
			&corev1.Secret{},
			fldPath,
			ContainElement(field.Required(fldPath.Child("spec.image"), "image is required")),
		),
		Entry("invalid dns server ip",
			&v1alpha1.ProviderSpec{
				DnsServers: []netip.Addr{invalidIP},
			},
			&corev1.Secret{},
			fldPath,
			ContainElement(field.Invalid(fldPath.Child("spec.dnsServers[0]"), invalidIP, "ip is invalid")),
		),
	)
})
