/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/pelotech/provider-freeipa/config/automemberadd"
	"github.com/pelotech/provider-freeipa/config/automemberaddcondition"
	"github.com/pelotech/provider-freeipa/config/dnsrecord"
	"github.com/pelotech/provider-freeipa/config/dnszone"
	"github.com/pelotech/provider-freeipa/config/group"
	"github.com/pelotech/provider-freeipa/config/hbacpolicy"
	"github.com/pelotech/provider-freeipa/config/hbacpolicyhostmembership"
	"github.com/pelotech/provider-freeipa/config/hbacpolicyservicemembership"
	"github.com/pelotech/provider-freeipa/config/hbacpolicyusermembership"
	"github.com/pelotech/provider-freeipa/config/host"
	"github.com/pelotech/provider-freeipa/config/hostgroup"
	"github.com/pelotech/provider-freeipa/config/hosthostgroupmembership"
	"github.com/pelotech/provider-freeipa/config/sudocmd"
	"github.com/pelotech/provider-freeipa/config/sudocmdgroup"
	"github.com/pelotech/provider-freeipa/config/sudocmdgroupmembership"
	"github.com/pelotech/provider-freeipa/config/sudorule"
	"github.com/pelotech/provider-freeipa/config/sudoruleallowcmdmembership"
	"github.com/pelotech/provider-freeipa/config/sudoruledenycmdmembership"
	"github.com/pelotech/provider-freeipa/config/sudorulehostmembership"
	"github.com/pelotech/provider-freeipa/config/sudoruleoption"
	"github.com/pelotech/provider-freeipa/config/sudorulerunasgroupmembership"
	"github.com/pelotech/provider-freeipa/config/sudorulerunasusermembership"
	"github.com/pelotech/provider-freeipa/config/sudoruleusermembership"
	"github.com/pelotech/provider-freeipa/config/user"
	"github.com/pelotech/provider-freeipa/config/usergroupmembership"
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
		automemberaddcondition.Configure,
		dnsrecord.Configure,
		dnszone.Configure,
		group.Configure,
		hbacpolicy.Configure,
		hbacpolicyhostmembership.Configure,
		hbacpolicyservicemembership.Configure,
		hbacpolicyusermembership.Configure,
		host.Configure,
		hosthostgroupmembership.Configure,
		hostgroup.Configure,
		sudocmd.Configure,
		sudocmdgroup.Configure,
		sudocmdgroupmembership.Configure,
		sudorule.Configure,
		sudoruleallowcmdmembership.Configure,
		sudoruledenycmdmembership.Configure,
		sudorulehostmembership.Configure,
		sudoruleoption.Configure,
		sudorulerunasgroupmembership.Configure,
		sudorulerunasusermembership.Configure,
		sudoruleusermembership.Configure,
		user.Configure,
		usergroupmembership.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
