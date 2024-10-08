package dns

import "github.com/crossplane/upjet/pkg/config"

// Configure adds configurations for dns group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_record", func(r *config.Resource) {
		r.ShortGroup = "dns"
		r.References["zone_id"] = config.Reference{
			Type: "github.com/chezmoi-sh/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
	})
}
