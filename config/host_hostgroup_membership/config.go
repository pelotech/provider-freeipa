package host_hostgroup_membership

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("freeipa_host_hostgroup_membership", func(r *config.Resource) {
		r.ShortGroup = "host_hostgroup_membership"
	})
}
