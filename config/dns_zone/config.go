package dns_zone

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("freeipa_dns_zone", func(r *config.Resource) {
		r.ShortGroup = "dns_zone"
	})
}
