package dnsrecord

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("freeipa_dns_record", func(r *config.Resource) {
		r.ShortGroup = "dns_record"
	})
}
