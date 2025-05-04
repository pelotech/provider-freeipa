/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"freeipa_automemberadd":                   config.IdentifierFromProvider,
	"freeipa_automemberadd_condition":         config.IdentifierFromProvider,
	"freeipa_dns_record":                      config.IdentifierFromProvider,
	"freeipa_dns_zone":                        config.IdentifierFromProvider,
	"freeipa_group":                           config.IdentifierFromProvider,
	"freeipa_hbac_policy":                     config.IdentifierFromProvider,
	"freeipa_hbac_policy_host_membership":     config.IdentifierFromProvider,
	"freeipa_hbac_policy_service_membership":  config.IdentifierFromProvider,
	"freeipa_hbac_policy_user_membership":     config.IdentifierFromProvider,
	"freeipa_host":                            config.IdentifierFromProvider,
	"freeipa_host_hostgroup_membership":       config.IdentifierFromProvider,
	"freeipa_hostgroup":                       config.IdentifierFromProvider,
	"freeipa_sudo_cmd":                        config.IdentifierFromProvider,
	"freeipa_sudo_cmdgroup":                   config.IdentifierFromProvider,
	"freeipa_sudo_cmdgroup_membership":        config.IdentifierFromProvider,
	"freeipa_sudo_rule":                       config.IdentifierFromProvider,
	"freeipa_sudo_rule_allowcmd_membership":   config.IdentifierFromProvider,
	"freeipa_sudo_rule_denycmd_membership":    config.IdentifierFromProvider,
	"freeipa_sudo_rule_host_membership":       config.IdentifierFromProvider,
	"freeipa_sudo_rule_option":                config.IdentifierFromProvider,
	"freeipa_sudo_rule_runasgroup_membership": config.IdentifierFromProvider,
	"freeipa_sudo_rule_runasuser_membership":  config.IdentifierFromProvider,
	"freeipa_sudo_rule_user_membership":       config.IdentifierFromProvider,
	"freeipa_user":                            config.IdentifierFromProvider,
	"freeipa_user_group_membership":           config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
