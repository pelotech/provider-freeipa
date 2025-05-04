// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	automemberadd "github.com/pelotech/provider-freeipa/internal/controller/automemberadd/automemberadd"
	condition "github.com/pelotech/provider-freeipa/internal/controller/automemberadd_condition/condition"
	record "github.com/pelotech/provider-freeipa/internal/controller/dns_record/record"
	zone "github.com/pelotech/provider-freeipa/internal/controller/dns_zone/zone"
	group "github.com/pelotech/provider-freeipa/internal/controller/group/group"
	policy "github.com/pelotech/provider-freeipa/internal/controller/hbac_policy/policy"
	policyhostmembership "github.com/pelotech/provider-freeipa/internal/controller/hbac_policy_host_membership/policyhostmembership"
	policyservicemembership "github.com/pelotech/provider-freeipa/internal/controller/hbac_policy_service_membership/policyservicemembership"
	policyusermembership "github.com/pelotech/provider-freeipa/internal/controller/hbac_policy_user_membership/policyusermembership"
	host "github.com/pelotech/provider-freeipa/internal/controller/host/host"
	hostgroupmembership "github.com/pelotech/provider-freeipa/internal/controller/host_hostgroup_membership/hostgroupmembership"
	hostgroup "github.com/pelotech/provider-freeipa/internal/controller/hostgroup/hostgroup"
	providerconfig "github.com/pelotech/provider-freeipa/internal/controller/providerconfig"
	cmd "github.com/pelotech/provider-freeipa/internal/controller/sudo_cmd/cmd"
	cmdgroup "github.com/pelotech/provider-freeipa/internal/controller/sudo_cmdgroup/cmdgroup"
	cmdgroupmembership "github.com/pelotech/provider-freeipa/internal/controller/sudo_cmdgroup_membership/cmdgroupmembership"
	rule "github.com/pelotech/provider-freeipa/internal/controller/sudo_rule/rule"
	ruleallowcmdmembership "github.com/pelotech/provider-freeipa/internal/controller/sudo_rule_allowcmd_membership/ruleallowcmdmembership"
	ruledenycmdmembership "github.com/pelotech/provider-freeipa/internal/controller/sudo_rule_denycmd_membership/ruledenycmdmembership"
	rulehostmembership "github.com/pelotech/provider-freeipa/internal/controller/sudo_rule_host_membership/rulehostmembership"
	ruleoption "github.com/pelotech/provider-freeipa/internal/controller/sudo_rule_option/ruleoption"
	rulerunasgroupmembership "github.com/pelotech/provider-freeipa/internal/controller/sudo_rule_runasgroup_membership/rulerunasgroupmembership"
	rulerunasusermembership "github.com/pelotech/provider-freeipa/internal/controller/sudo_rule_runasuser_membership/rulerunasusermembership"
	ruleusermembership "github.com/pelotech/provider-freeipa/internal/controller/user/ruleusermembership"
	user "github.com/pelotech/provider-freeipa/internal/controller/user/user"
	groupmembership "github.com/pelotech/provider-freeipa/internal/controller/user_group_membership/groupmembership"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		automemberadd.Setup,
		condition.Setup,
		record.Setup,
		zone.Setup,
		group.Setup,
		policy.Setup,
		policyhostmembership.Setup,
		policyservicemembership.Setup,
		policyusermembership.Setup,
		host.Setup,
		hostgroupmembership.Setup,
		hostgroup.Setup,
		providerconfig.Setup,
		cmd.Setup,
		cmdgroup.Setup,
		cmdgroupmembership.Setup,
		rule.Setup,
		ruleallowcmdmembership.Setup,
		ruledenycmdmembership.Setup,
		rulehostmembership.Setup,
		ruleoption.Setup,
		rulerunasgroupmembership.Setup,
		rulerunasusermembership.Setup,
		ruleusermembership.Setup,
		user.Setup,
		groupmembership.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
