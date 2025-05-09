package sudocmdgroup

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("freeipa_sudo_cmdgroup", func(r *config.Resource) {
		r.ShortGroup = "sudo_cmdgroup"
	})
}
