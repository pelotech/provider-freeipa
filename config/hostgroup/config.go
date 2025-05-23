package hostgroup

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("freeipa_hostgroup", func(r *config.Resource) {
		r.ShortGroup = "hostgroup"
	})
}
