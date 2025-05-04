package usergroupmembership

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("freeipa_user_group_membership", func(r *config.Resource) {
		r.ShortGroup = "user_group_membership"
	})
}
