#!/usr/bin/env bash
set -o allexport -o errexit -o nounset -o pipefail

echo "Running post-assert-hooks/account.api-token.sh"

# Check if the API token is set
echo "Checking if the API token is set..."
api_token=$(${KUBECTL} -n default get secret cloudflare-api-token-example -o json | ${YQ} '.data["attribute.value"] | @base64d')
if [[ -z ${api_token} ]]; then
	echo "API token not found in secret cloudflare-api-token-example"
	exit 1
fi

# Validate the API token
echo "Validating the API token..."
token_validation=$(curl -sSl -X GET "https://api.cloudflare.com/client/v4/user/tokens/verify" -H "Authorization: Bearer ${api_token}")
echo "${token_validation}" | ${YQ} --prettyPrint | sed 's/^/  /'

# Check if the token verification was successful
if ! echo "${token_validation}" | ${YQ} -e '.success' | grep -q "true"; then
	echo "API token verification failed"
	exit 1
fi
echo "API token verification succeeded"
