// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	account "github.com/chezmoi-sh/provider-cloudflare/internal/controller/account/account"
	apitoken "github.com/chezmoi-sh/provider-cloudflare/internal/controller/account/apitoken"
	dnssec "github.com/chezmoi-sh/provider-cloudflare/internal/controller/dns/dnssec"
	record "github.com/chezmoi-sh/provider-cloudflare/internal/controller/dns/record"
	providerconfig "github.com/chezmoi-sh/provider-cloudflare/internal/controller/providerconfig"
	zone "github.com/chezmoi-sh/provider-cloudflare/internal/controller/zone/zone"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		account.Setup,
		apitoken.Setup,
		dnssec.Setup,
		record.Setup,
		providerconfig.Setup,
		zone.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
