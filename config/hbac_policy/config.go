package hbac_policy

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("freeipa_hbac_policy", func(r *config.Resource) {
		r.ShortGroup = "hbac_policy"
	})
}
