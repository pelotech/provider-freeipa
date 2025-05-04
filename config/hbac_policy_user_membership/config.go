package hbac_policy_user_membership

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("freeipa_hbac_policy_user_membership", func(r *config.Resource) {
		r.ShortGroup = "hbac_policy_user_membership"
	})
}
