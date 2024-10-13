# Provider Cloudflare

`provider-cloudflare` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
Cloudflare API.

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/chezmoi-sh/provider-cloudflare):

```
up ctp provider install chezmoi-sh/provider-cloudflare:v0.1.0
```

Alternatively, you can use declarative installation:

```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-cloudflare
spec:
  package: chezmoi-sh/provider-cloudflare:v0.1.0
EOF
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/chezmoi-sh/provider-cloudflare).

## Resources Reference and Status

Table below shows the resources that are currently supported by the provider, their relation with the
Cloudflare Terraform provider and their status.

> [!NOTE]
> Possible statuses are:
>
> - :white_check_mark:: Supported and automatically tested through E2E
> - :passport_control:: Supported but manually tested
> - :warning:: Supported but not tested
> - :black_square_button:: Unsupported / Not implemented
> - :no_entry_sign:: Deprecated

| Status                | API Group                        | Kind     | Terraform Equivalent                                       |
| --------------------- | -------------------------------- | -------- | ---------------------------------------------------------- |
| :white_check_mark:    | account.cloudflare.crossplane.io | APIToken | cloudflare_api_token                                       |
| :passport_control:    | account.cloudflare.crossplane.io | Account  | cloudflare_account                                         |
| :white_check_mark:    | dns.cloudflare.crossplane.io     | DNSSEC   | cloudflare_zone_dnssec                                     |
| :white_check_mark:    | dns.cloudflare.crossplane.io     | Record   | cloudflare_record                                          |
| :passport_control:    | zone.cloudflare.crossplane.io    | Zone     | cloudflare_zone                                            |
| :black_square_button: |                                  |          | cloudflare_access_rule                                     |
| :black_square_button: |                                  |          | cloudflare_account_member                                  |
| :black_square_button: |                                  |          | cloudflare_address_map                                     |
| :black_square_button: |                                  |          | cloudflare_api_shield_operation_schema_validation_settings |
| :black_square_button: |                                  |          | cloudflare_api_shield_operation                            |
| :black_square_button: |                                  |          | cloudflare_api_shield_schema_validation_settings           |
| :black_square_button: |                                  |          | cloudflare_api_shield_schema                               |
| :black_square_button: |                                  |          | cloudflare_api_shield                                      |
| :black_square_button: |                                  |          | cloudflare_argo                                            |
| :black_square_button: |                                  |          | cloudflare_authenticated_origin_pulls_certificate          |
| :black_square_button: |                                  |          | cloudflare_authenticated_origin_pulls                      |
| :black_square_button: |                                  |          | cloudflare_bot_management                                  |
| :black_square_button: |                                  |          | cloudflare_byo_ip_prefix                                   |
| :black_square_button: |                                  |          | cloudflare_certificate_pack                                |
| :black_square_button: |                                  |          | cloudflare_cloud_connector_rules                           |
| :black_square_button: |                                  |          | cloudflare_custom_hostname_fallback_origin                 |
| :black_square_button: |                                  |          | cloudflare_custom_hostname                                 |
| :black_square_button: |                                  |          | cloudflare_custom_pages                                    |
| :black_square_button: |                                  |          | cloudflare_custom_ssl                                      |
| :black_square_button: |                                  |          | cloudflare_d1_database                                     |
| :black_square_button: |                                  |          | cloudflare_email_routing_address                           |
| :black_square_button: |                                  |          | cloudflare_email_routing_catch_all                         |
| :black_square_button: |                                  |          | cloudflare_email_routing_rule                              |
| :black_square_button: |                                  |          | cloudflare_email_routing_settings                          |
| :black_square_button: |                                  |          | cloudflare_healthcheck                                     |
| :black_square_button: |                                  |          | cloudflare_hostname_tls_setting_ciphers                    |
| :black_square_button: |                                  |          | cloudflare_hostname_tls_setting                            |
| :black_square_button: |                                  |          | cloudflare_hyperdrive_config                               |
| :black_square_button: |                                  |          | cloudflare_infrastructure_access_target                    |
| :black_square_button: |                                  |          | cloudflare_keyless_certificate                             |
| :black_square_button: |                                  |          | cloudflare_list_item                                       |
| :black_square_button: |                                  |          | cloudflare_list                                            |
| :black_square_button: |                                  |          | cloudflare_load_balancer_monitor                           |
| :black_square_button: |                                  |          | cloudflare_load_balancer_pool                              |
| :black_square_button: |                                  |          | cloudflare_load_balancer                                   |
| :black_square_button: |                                  |          | cloudflare_logpull_retention                               |
| :black_square_button: |                                  |          | cloudflare_logpush_job                                     |
| :black_square_button: |                                  |          | cloudflare_logpush_ownership_challenge                     |
| :black_square_button: |                                  |          | cloudflare_magic_firewall_ruleset                          |
| :black_square_button: |                                  |          | cloudflare_magic_wan_gre_tunnel                            |
| :black_square_button: |                                  |          | cloudflare_magic_wan_ipsec_tunnel                          |
| :black_square_button: |                                  |          | cloudflare_magic_wan_static_route                          |
| :black_square_button: |                                  |          | cloudflare_managed_headers                                 |
| :black_square_button: |                                  |          | cloudflare_mtls_certificate                                |
| :black_square_button: |                                  |          | cloudflare_notification_policy_webhooks                    |
| :black_square_button: |                                  |          | cloudflare_notification_policy                             |
| :black_square_button: |                                  |          | cloudflare_observatory_scheduled_test                      |
| :black_square_button: |                                  |          | cloudflare_origin_ca_certificate                           |
| :black_square_button: |                                  |          | cloudflare_page_rule                                       |
| :black_square_button: |                                  |          | cloudflare_pages_domain                                    |
| :black_square_button: |                                  |          | cloudflare_pages_project                                   |
| :black_square_button: |                                  |          | cloudflare_queue                                           |
| :black_square_button: |                                  |          | cloudflare_r2_bucket                                       |
| :black_square_button: |                                  |          | cloudflare_regional_hostname                               |
| :black_square_button: |                                  |          | cloudflare_regional_tiered_cache                           |
| :black_square_button: |                                  |          | cloudflare_risk_behavior                                   |
| :black_square_button: |                                  |          | cloudflare_ruleset                                         |
| :black_square_button: |                                  |          | cloudflare_spectrum_application                            |
| :black_square_button: |                                  |          | cloudflare_tiered_cache                                    |
| :black_square_button: |                                  |          | cloudflare_total_tls                                       |
| :black_square_button: |                                  |          | cloudflare_turnstile_widget                                |
| :black_square_button: |                                  |          | cloudflare_url_normalization_settings                      |
| :black_square_button: |                                  |          | cloudflare_user_agent_blocking_rule                        |
| :black_square_button: |                                  |          | cloudflare_waiting_room_event                              |
| :black_square_button: |                                  |          | cloudflare_waiting_room_rules                              |
| :black_square_button: |                                  |          | cloudflare_waiting_room_settings                           |
| :black_square_button: |                                  |          | cloudflare_waiting_room                                    |
| :black_square_button: |                                  |          | cloudflare_web3_hostname                                   |
| :black_square_button: |                                  |          | cloudflare_web_analytics_rule                              |
| :black_square_button: |                                  |          | cloudflare_web_analytics_site                              |
| :black_square_button: |                                  |          | cloudflare_workers_cron_trigger                            |
| :black_square_button: |                                  |          | cloudflare_workers_domain                                  |
| :black_square_button: |                                  |          | cloudflare_workers_for_platforms_dispatch_namespace        |
| :black_square_button: |                                  |          | cloudflare_workers_for_platforms_namespace                 |
| :black_square_button: |                                  |          | cloudflare_workers_kv_namespace                            |
| :black_square_button: |                                  |          | cloudflare_workers_kv                                      |
| :black_square_button: |                                  |          | cloudflare_workers_route                                   |
| :black_square_button: |                                  |          | cloudflare_workers_script                                  |
| :black_square_button: |                                  |          | cloudflare_workers_secret                                  |
| :black_square_button: |                                  |          | cloudflare_zero_trust_access_application                   |
| :black_square_button: |                                  |          | cloudflare_zero_trust_access_custom_page                   |
| :black_square_button: |                                  |          | cloudflare_zero_trust_access_group                         |
| :black_square_button: |                                  |          | cloudflare_zero_trust_access_identity_provider             |
| :black_square_button: |                                  |          | cloudflare_zero_trust_access_mtls_certificate              |
| :black_square_button: |                                  |          | cloudflare_zero_trust_access_mtls_hostname_settings        |
| :black_square_button: |                                  |          | cloudflare_zero_trust_access_organization                  |
| :black_square_button: |                                  |          | cloudflare_zero_trust_access_policy                        |
| :black_square_button: |                                  |          | cloudflare_zero_trust_access_service_token                 |
| :black_square_button: |                                  |          | cloudflare_zero_trust_access_short_lived_certificate       |
| :black_square_button: |                                  |          | cloudflare_zero_trust_access_tag                           |
| :black_square_button: |                                  |          | cloudflare_zero_trust_device_certificates                  |
| :black_square_button: |                                  |          | cloudflare_zero_trust_device_managed_networks              |
| :black_square_button: |                                  |          | cloudflare_zero_trust_device_posture_integration           |
| :black_square_button: |                                  |          | cloudflare_zero_trust_device_posture_rule                  |
| :black_square_button: |                                  |          | cloudflare_zero_trust_device_profiles                      |
| :black_square_button: |                                  |          | cloudflare_zero_trust_dex_test                             |
| :black_square_button: |                                  |          | cloudflare_zero_trust_dlp_profile                          |
| :black_square_button: |                                  |          | cloudflare_zero_trust_dns_location                         |
| :black_square_button: |                                  |          | cloudflare_zero_trust_gateway_certificate                  |
| :black_square_button: |                                  |          | cloudflare_zero_trust_gateway_policy                       |
| :black_square_button: |                                  |          | cloudflare_zero_trust_gateway_proxy_endpoint               |
| :black_square_button: |                                  |          | cloudflare_zero_trust_gateway_settings                     |
| :black_square_button: |                                  |          | cloudflare_zero_trust_key_access_key_configuration         |
| :black_square_button: |                                  |          | cloudflare_zero_trust_list                                 |
| :black_square_button: |                                  |          | cloudflare_zero_trust_local_fallback_domain                |
| :black_square_button: |                                  |          | cloudflare_zero_trust_risk_behavior                        |
| :black_square_button: |                                  |          | cloudflare_zero_trust_risk_score_integration               |
| :black_square_button: |                                  |          | cloudflare_zero_trust_split_tunnel                         |
| :black_square_button: |                                  |          | cloudflare_zero_trust_tunnel_cloudflared_config            |
| :black_square_button: |                                  |          | cloudflare_zero_trust_tunnel_cloudflared                   |
| :black_square_button: |                                  |          | cloudflare_zero_trust_tunnel_route                         |
| :black_square_button: |                                  |          | cloudflare_zero_trust_tunnel_virtual_network               |
| :black_square_button: |                                  |          | cloudflare_zone_cache_reserve                              |
| :black_square_button: |                                  |          | cloudflare_zone_cache_variants                             |
| :black_square_button: |                                  |          | cloudflare_zone_hold                                       |
| :black_square_button: |                                  |          | cloudflare_zone_lockdown                                   |
| :black_square_button: |                                  |          | cloudflare_zone_settings_override                          |
| :no_entry_sign:       |                                  |          | cloudflare_access_application                              |
| :no_entry_sign:       |                                  |          | cloudflare_access_ca_certificate                           |
| :no_entry_sign:       |                                  |          | cloudflare_access_custom_page                              |
| :no_entry_sign:       |                                  |          | cloudflare_access_group                                    |
| :no_entry_sign:       |                                  |          | cloudflare_access_identity_provider                        |
| :no_entry_sign:       |                                  |          | cloudflare_access_keys_configuration                       |
| :no_entry_sign:       |                                  |          | cloudflare_access_mutual_tls_certificate                   |
| :no_entry_sign:       |                                  |          | cloudflare_access_mutual_tls_hostname_settings             |
| :no_entry_sign:       |                                  |          | cloudflare_access_organization                             |
| :no_entry_sign:       |                                  |          | cloudflare_access_policy                                   |
| :no_entry_sign:       |                                  |          | cloudflare_access_service_token                            |
| :no_entry_sign:       |                                  |          | cloudflare_access_tag                                      |
| :no_entry_sign:       |                                  |          | cloudflare_device_dex_test                                 |
| :no_entry_sign:       |                                  |          | cloudflare_device_managed_networks                         |
| :no_entry_sign:       |                                  |          | cloudflare_device_policy_certificates                      |
| :no_entry_sign:       |                                  |          | cloudflare_device_posture_integration                      |
| :no_entry_sign:       |                                  |          | cloudflare_device_posture_rule                             |
| :no_entry_sign:       |                                  |          | cloudflare_device_settings_policy                          |
| :no_entry_sign:       |                                  |          | cloudflare_dlp_profile                                     |
| :no_entry_sign:       |                                  |          | cloudflare_fallback_domain                                 |
| :no_entry_sign:       |                                  |          | cloudflare_filter                                          |
| :no_entry_sign:       |                                  |          | cloudflare_firewall_rule                                   |
| :no_entry_sign:       |                                  |          | cloudflare_gre_tunnel                                      |
| :no_entry_sign:       |                                  |          | cloudflare_ipsec_tunnel                                    |
| :no_entry_sign:       |                                  |          | cloudflare_rate_limit                                      |
| :no_entry_sign:       |                                  |          | cloudflare_split_tunnel                                    |
| :no_entry_sign:       |                                  |          | cloudflare_static_route                                    |
| :no_entry_sign:       |                                  |          | cloudflare_teams_account                                   |
| :no_entry_sign:       |                                  |          | cloudflare_teams_list                                      |
| :no_entry_sign:       |                                  |          | cloudflare_teams_location                                  |
| :no_entry_sign:       |                                  |          | cloudflare_teams_proxy_endpoint                            |
| :no_entry_sign:       |                                  |          | cloudflare_teams_rule                                      |
| :no_entry_sign:       |                                  |          | cloudflare_tunnel_config                                   |
| :no_entry_sign:       |                                  |          | cloudflare_tunnel_route                                    |
| :no_entry_sign:       |                                  |          | cloudflare_tunnel_virtual_network                          |
| :no_entry_sign:       |                                  |          | cloudflare_tunnel                                          |
| :no_entry_sign:       |                                  |          | cloudflare_worker_cron_trigger                             |
| :no_entry_sign:       |                                  |          | cloudflare_worker_domain                                   |
| :no_entry_sign:       |                                  |          | cloudflare_worker_route                                    |
| :no_entry_sign:       |                                  |          | cloudflare_worker_script                                   |
| :no_entry_sign:       |                                  |          | cloudflare_worker_secret                                   |

## Contributing

In order to provide a better experience for the developers, this provider uses a [Nix Flake](https://nixos.wiki/wiki/Flakes)
to manage the development environment.
If you are familiar with [direnv](https://direnv.net/) and you have [Nix](https://nixos.org/download.html) installed,
all you need is to allow the `.envrc` file to be loaded by direnv and everything will be set up for you.

> [!NOTE]
> For other users, I will create a Devcontainer for this project in the future.

Here is a step-by-step guide to get you started:

```console
# Clone the repository and go to the project directory
git clone https://github.com/chezmoi-sh/provider-cloudflare
cd provider-cloudflare

# Allow the .envrc file to be loaded (will be slow the first time)
direnv allow

# Install the dependencies
make submodules
```

Now, everything is set up and you can start developing.

### How add a new resource

Like other Crossplane providers, this provider uses the Upjet code generation tools to generate the code
for the resources, based on the Terraform provider.  
However, because the Terraform provider for Cloudflare is a massive project (150+ resources), I decided to
use an intermediate step to generate the code for the resources.

This intermediate step is a [Go program](./scripts/generate_provider_config.go) that takes an
[inventory of resources](./provider-cloudflare.resources.yaml) and generates the Go code for all resources.

> [!IMPORTANT]
> As this program is not perfect, if a resource needs some customization, you will need to update the program
> to add the necessary logic.

So, here is a step-by-step guide to add a new resource:

1. Add the resource to the [inventory](./provider-cloudflare.resources.yaml) file, inside the `supported` section.  
   _NOTE: Everything we need to know about the structure of the resource is in the inventory file it-self, at
   the beginning of the file._
2. Run `make provider-cloudflare.generate` to generate the Go files containing the resource configuration.
3. Run `make generate` to generate the provider code.
4. Add examples and tests for the new resource, inside the `examples` directory.
5. Run `make e2e` to test the new resource.
6. Update the [Resources Reference and Status](#resources-reference-and-status) table in the README file.
7. Create a pull request.

> [!IMPORTANT]
> This repository uses the [Gitmoji](https://gitmoji.dev/) convention for commit messages and uses
> [Trunk](https://trunk.io) to format and lint the code.

> [!NOTE]
> I've tried to automate as much as possible without leaving the dev experience provided by Upjet and its
> documentation; apart from the Go code generation part, everything else follows the
> [official documentation](https://github.com/crossplane/upjet/blob/main/docs/generating-a-provider.md).

## Roadmap

- [ ] Publish the provider to the Crossplane marketplace through the Github Actions workflow
- [ ] Make it publicly available in the Crossplane marketplace
- [ ] Add Devcontainer for people who don't use Nix
- [ ] Configure Renovate to improve how it updates the dependencies (grouping with logical changes, updating
      dependencies that are cross-referenced by the CI and the Makefile)
- [ ] Add some checks to the CI to ensure that all changes follow the guidelines

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/chezmoi-sh/provider-cloudflare/issues).

## Why another Cloudflare provider?

This provider is mainly inspired by the one created by [cdloh](https://github.com/cdloh/provider-cloudflare),
which is much more complete and feature-rich than this one _(for now at least)_.  
The main reason for creating this provider was to learn how to create a Crossplane provider and to have a provider
that is more aligned with the way resources are displayed inside the Cloudflare dashboard. For example, the DNSSEC
resource should be a `DNS` resource and not a `zone` resource.  
The other reason is that I wanted to have a provider that is easier to maintain through code generation, which will
simplify the migration from v4 to v5 of Cloudflare's official terraform provider.

## License

This provider is released under the Apache 2.0 license. See the [LICENSE](./LICENSE) file for more details.
