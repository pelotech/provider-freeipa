package sudoruledenycmdmembership

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("freeipa_sudo_rule_denycmd_membership", func(r *config.Resource) {
		r.ShortGroup = "sudo_rule_denycmd_membership"
	})
}
