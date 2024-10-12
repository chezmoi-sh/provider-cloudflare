/*
Copyright 2021 Upbound Inc.
*/
package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/pkg/terraform"

	"github.com/chezmoi-sh/provider-cloudflare/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal cloudflare credentials as JSON"

	// configuration keys
	keyApiKey                  = "api_key"
	keyApiToken                = "api_token"
	keyApiUserServiceKey       = "api_user_service_key"
	keyBaseURL                 = "base_url"
	keyEmail                   = "email"
	keyUserAgentOperatorSuffix = "user_agent_operator_suffix"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
//
//nolint:gocyclo
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds := map[string]string{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		// Set credentials in Terraform provider configuration.
		ps.Configuration = map[string]any{}
		if apiKey, ok := creds[keyApiKey]; ok {
			ps.Configuration[keyApiKey] = apiKey
		}
		if apiToken, ok := creds[keyApiToken]; ok {
			ps.Configuration[keyApiToken] = apiToken
		}
		if apiUserServiceKey, ok := creds[keyApiUserServiceKey]; ok {
			ps.Configuration[keyApiUserServiceKey] = apiUserServiceKey
		}
		if baseURL, ok := creds[keyBaseURL]; ok {
			ps.Configuration[keyBaseURL] = baseURL
		}
		if email, ok := creds[keyEmail]; ok {
			ps.Configuration[keyEmail] = email
		}
		if userAgentOperatorSuffix, ok := creds[keyUserAgentOperatorSuffix]; ok {
			ps.Configuration[keyUserAgentOperatorSuffix] = userAgentOperatorSuffix
		}
		return ps, nil
	}
}
