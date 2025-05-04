/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/pelotech/provider-freeipa/config/automemberadd"
	"github.com/pelotech/provider-freeipa/config/automemberadd_condition"
	"github.com/pelotech/provider-freeipa/config/dns_record"
	"github.com/pelotech/provider-freeipa/config/dns_zone"
	"github.com/pelotech/provider-freeipa/config/group"
	"github.com/pelotech/provider-freeipa/config/hbac_policy"
	"github.com/pelotech/provider-freeipa/config/hbac_policy_host_membership"
	"github.com/pelotech/provider-freeipa/config/hbac_policy_service_membership"
	"github.com/pelotech/provider-freeipa/config/hbac_policy_user_membership"
	"github.com/pelotech/provider-freeipa/config/host"
	"github.com/pelotech/provider-freeipa/config/host_hostgroup_membership"
	"github.com/pelotech/provider-freeipa/config/hostgroup"
	"github.com/pelotech/provider-freeipa/config/sudo_cmd"
	"github.com/pelotech/provider-freeipa/config/sudo_cmdgroup"
	"github.com/pelotech/provider-freeipa/config/sudo_cmdgroup_membership"
	"github.com/pelotech/provider-freeipa/config/sudo_rule"
	"github.com/pelotech/provider-freeipa/config/sudo_rule_allowcmd_membership"
	"github.com/pelotech/provider-freeipa/config/sudo_rule_denycmd_membership"
	"github.com/pelotech/provider-freeipa/config/sudo_rule_host_membership"
	"github.com/pelotech/provider-freeipa/config/sudo_rule_option"
	"github.com/pelotech/provider-freeipa/config/sudo_rule_runasgroup_membership"
	"github.com/pelotech/provider-freeipa/config/sudo_rule_runasuser_membership"
	"github.com/pelotech/provider-freeipa/config/sudo_rule_user_membership"
	"github.com/pelotech/provider-freeipa/config/user"
	"github.com/pelotech/provider-freeipa/config/user_group_membership"
)

const (
	resourcePrefix = "freeipa"
	modulePath     = "github.com/pelotech/provider-freeipa"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("pelotech"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		automemberadd.Configure,
		automemberadd_condition.Configure,
		dns_record.Configure,
		dns_zone.Configure,
		group.Configure,
		hbac_policy.Configure,
		hbac_policy_host_membership.Configure,
		hbac_policy_service_membership.Configure,
		hbac_policy_user_membership.Configure,
		host.Configure,
		host_hostgroup_membership.Configure,
		hostgroup.Configure,
		sudo_cmd.Configure,
		sudo_cmdgroup.Configure,
		sudo_cmdgroup_membership.Configure,
		sudo_rule.Configure,
		sudo_rule_allowcmd_membership.Configure,
		sudo_rule_denycmd_membership.Configure,
		sudo_rule_host_membership.Configure,
		sudo_rule_option.Configure,
		sudo_rule_runasgroup_membership.Configure,
		sudo_rule_runasuser_membership.Configure,
		sudo_rule_user_membership.Configure,
		user.Configure,
		user_group_membership.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
