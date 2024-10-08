/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// "cloudflare_access_rule":                                     config.IdentifierFromProvider,
	// "cloudflare_account":                                         config.IdentifierFromProvider,
	// "cloudflare_account_member":                                  config.IdentifierFromProvider,
	// "cloudflare_account_subscription":                            config.IdentifierFromProvider,
	// "cloudflare_address_map":                                     config.IdentifierFromProvider,
	// "cloudflare_api_shield":                                      config.IdentifierFromProvider,
	// "cloudflare_api_shield_operation":                            config.IdentifierFromProvider,
	// "cloudflare_api_shield_operation_schema_validation_settings": config.IdentifierFromProvider,
	// "cloudflare_api_shield_schema":                               config.IdentifierFromProvider,
	// "cloudflare_api_shield_schema_validation_settings":           config.IdentifierFromProvider,
	// "cloudflare_api_token":                                       config.IdentifierFromProvider,
	// "cloudflare_argo_smart_routing":                              config.IdentifierFromProvider,
	// "cloudflare_argo_tiered_caching":                             config.IdentifierFromProvider,
	// "cloudflare_authenticated_origin_pulls":                      config.IdentifierFromProvider,
	// "cloudflare_authenticated_origin_pulls_certificate":          config.IdentifierFromProvider,
	// "cloudflare_bot_management":                                  config.IdentifierFromProvider,
	// "cloudflare_byo_ip_prefix":                                   config.IdentifierFromProvider,
	// "cloudflare_certificate_pack":                                config.IdentifierFromProvider,
	// "cloudflare_custom_hostname":                                 config.IdentifierFromProvider,
	// "cloudflare_custom_hostname_fallback_origin":                 config.IdentifierFromProvider,
	// "cloudflare_custom_ssl":                                      config.IdentifierFromProvider,
	// "cloudflare_d1_database":                                     config.IdentifierFromProvider,
	"cloudflare_record": config.IdentifierFromProvider,
	// "cloudflare_dns_zone_dnssec":                                 config.IdentifierFromProvider,
	// "cloudflare_email_routing_address":                           config.IdentifierFromProvider,
	// "cloudflare_email_routing_catch_all":                         config.IdentifierFromProvider,
	// "cloudflare_email_routing_rule":                              config.IdentifierFromProvider,
	// "cloudflare_email_routing_settings":                          config.IdentifierFromProvider,
	// "cloudflare_filter":                                          config.IdentifierFromProvider,
	// "cloudflare_firewall_rule":                                   config.IdentifierFromProvider,
	// "cloudflare_healthcheck":                                     config.IdentifierFromProvider,
	// "cloudflare_hostname_tls_setting":                            config.IdentifierFromProvider,
	// "cloudflare_hyperdrive_config":                               config.IdentifierFromProvider,
	// "cloudflare_keyless_certificate":                             config.IdentifierFromProvider,
	// "cloudflare_list":                                            config.IdentifierFromProvider,
	// "cloudflare_list_item":                                       config.IdentifierFromProvider,
	// "cloudflare_load_balancer":                                   config.IdentifierFromProvider,
	// "cloudflare_load_balancer_monitor":                           config.IdentifierFromProvider,
	// "cloudflare_load_balancer_pool":                              config.IdentifierFromProvider,
	// "cloudflare_logpull_retention":                               config.IdentifierFromProvider,
	// "cloudflare_logpush_job":                                     config.IdentifierFromProvider,
	// "cloudflare_logpush_ownership_challenge":                     config.IdentifierFromProvider,
	// "cloudflare_magic_firewall_ruleset":                          config.IdentifierFromProvider,
	// "cloudflare_magic_wan_gre_tunnel":                            config.IdentifierFromProvider,
	// "cloudflare_magic_wan_ipsec_tunnel":                          config.IdentifierFromProvider,
	// "cloudflare_magic_wan_static_route":                          config.IdentifierFromProvider,
	// "cloudflare_managed_headers":                                 config.IdentifierFromProvider,
	// "cloudflare_mtls_certificate":                                config.IdentifierFromProvider,
	// "cloudflare_notification_policy":                             config.IdentifierFromProvider,
	// "cloudflare_notification_policy_webhooks":                    config.IdentifierFromProvider,
	// "cloudflare_observatory_scheduled_test":                      config.IdentifierFromProvider,
	// "cloudflare_origin_ca_certificate":                           config.IdentifierFromProvider,
	// "cloudflare_page_rule":                                       config.IdentifierFromProvider,
	// "cloudflare_pages_domain":                                    config.IdentifierFromProvider,
	// "cloudflare_pages_project":                                   config.IdentifierFromProvider,
	// "cloudflare_queue":                                           config.IdentifierFromProvider,
	// "cloudflare_r2_bucket":                                       config.IdentifierFromProvider,
	// "cloudflare_rate_limit":                                      config.IdentifierFromProvider,
	// "cloudflare_regional_hostname":                               config.IdentifierFromProvider,
	// "cloudflare_regional_tiered_cache":                           config.IdentifierFromProvider,
	// "cloudflare_ruleset":                                         config.IdentifierFromProvider,
	// "cloudflare_spectrum_application":                            config.IdentifierFromProvider,
	// "cloudflare_tiered_cache":                                    config.IdentifierFromProvider,
	// "cloudflare_total_tls":                                       config.IdentifierFromProvider,
	// "cloudflare_turnstile_widget":                                config.IdentifierFromProvider,
	// "cloudflare_url_normalization_settings":                      config.IdentifierFromProvider,
	// "cloudflare_user_agent_blocking_rule":                        config.IdentifierFromProvider,
	// "cloudflare_waiting_room":                                    config.IdentifierFromProvider,
	// "cloudflare_waiting_room_event":                              config.IdentifierFromProvider,
	// "cloudflare_waiting_room_rules":                              config.IdentifierFromProvider,
	// "cloudflare_waiting_room_setting":                            config.IdentifierFromProvider,
	// "cloudflare_web_analytics_rule":                              config.IdentifierFromProvider,
	// "cloudflare_web_analytics_site":                              config.IdentifierFromProvider,
	// "cloudflare_web3_hostname":                                   config.IdentifierFromProvider,
	// "cloudflare_workers_cron_trigger":                            config.IdentifierFromProvider,
	// "cloudflare_workers_custom_domain":                           config.IdentifierFromProvider,
	// "cloudflare_workers_for_platforms_dispatch_namespace":        config.IdentifierFromProvider,
	// "cloudflare_workers_kv":                                      config.IdentifierFromProvider,
	// "cloudflare_workers_kv_namespace":                            config.IdentifierFromProvider,
	// "cloudflare_workers_script":                                  config.IdentifierFromProvider,
	// "cloudflare_workers_secret":                                  config.IdentifierFromProvider,
	// "cloudflare_zero_trust_access_application":                   config.IdentifierFromProvider,
	// "cloudflare_zero_trust_access_custom_page":                   config.IdentifierFromProvider,
	// "cloudflare_zero_trust_access_group":                         config.IdentifierFromProvider,
	// "cloudflare_zero_trust_access_identity_provider":             config.IdentifierFromProvider,
	// "cloudflare_zero_trust_access_key_configuration":             config.IdentifierFromProvider,
	// "cloudflare_zero_trust_access_mtls_certificate":              config.IdentifierFromProvider,
	// "cloudflare_zero_trust_access_mtls_hostname_settings":        config.IdentifierFromProvider,
	// "cloudflare_zero_trust_access_policy":                        config.IdentifierFromProvider,
	// "cloudflare_zero_trust_access_service_token":                 config.IdentifierFromProvider,
	// "cloudflare_zero_trust_access_short_lived_certificate":       config.IdentifierFromProvider,
	// "cloudflare_zero_trust_access_tag":                           config.IdentifierFromProvider,
	// "cloudflare_zero_trust_device_certificates":                  config.IdentifierFromProvider,
	// "cloudflare_zero_trust_device_managed_networks":              config.IdentifierFromProvider,
	// "cloudflare_zero_trust_device_posture_integration":           config.IdentifierFromProvider,
	// "cloudflare_zero_trust_device_posture_rule":                  config.IdentifierFromProvider,
	// "cloudflare_zero_trust_device_profiles":                      config.IdentifierFromProvider,
	// "cloudflare_zero_trust_dex_test":                             config.IdentifierFromProvider,
	// "cloudflare_zero_trust_dlp_custom_profile":                   config.IdentifierFromProvider,
	// "cloudflare_zero_trust_dlp_predefined_profile":               config.IdentifierFromProvider,
	// "cloudflare_zero_trust_dns_location":                         config.IdentifierFromProvider,
	// "cloudflare_zero_trust_gateway_policy":                       config.IdentifierFromProvider,
	// "cloudflare_zero_trust_gateway_proxy_endpoint":               config.IdentifierFromProvider,
	// "cloudflare_zero_trust_gateway_settings":                     config.IdentifierFromProvider,
	// "cloudflare_zero_trust_list":                                 config.IdentifierFromProvider,
	// "cloudflare_zero_trust_local_domain_fallback":                config.IdentifierFromProvider,
	// "cloudflare_zero_trust_organization":                         config.IdentifierFromProvider,
	// "cloudflare_zero_trust_tunnel_cloudflared":                   config.IdentifierFromProvider,
	// "cloudflare_zero_trust_tunnel_cloudflared_config":            config.IdentifierFromProvider,
	// "cloudflare_zero_trust_tunnel_cloudflared_route":             config.IdentifierFromProvider,
	// "cloudflare_zero_trust_tunnel_cloudflared_virtual_network":   config.IdentifierFromProvider,
	"cloudflare_zone": config.IdentifierFromProvider,
	// "cloudflare_zone_cache_reserve":                              config.IdentifierFromProvider,
	// "cloudflare_zone_cache_variants":                             config.IdentifierFromProvider,
	// "cloudflare_zone_hold":                                       config.IdentifierFromProvider,
	// "cloudflare_zone_lockdown":                                   config.IdentifierFromProvider,
	// "cloudflare_zone_setting":                                    config.IdentifierFromProvider,
	// "cloudflare_zone_subscription":                               config.IdentifierFromProvider,
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
