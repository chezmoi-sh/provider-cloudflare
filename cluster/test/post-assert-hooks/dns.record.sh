#!/usr/bin/env bash
set -o allexport -o errexit -o nounset -o pipefail

echo "Running post-assert-hooks/dns.record.sh"

# Retrieve the data from the data-source ($UPTEST_DATASOURCE_PATH)
echo "Loading data from ${UPTEST_DATASOURCE_PATH}..."
UPTEST_DATASOURCE_PATH="$(realpath "$(dirname "${BASH_SOURCE[0]}")/../../../${UPTEST_DATASOURCE_PATH}")"
cloudflare_zone_name=$(${YQ} '.cloudflare_zone' "${UPTEST_DATASOURCE_PATH}")

# Check if the record has been created
echo "Checking if the record has been created..."
timeout 30s bash -c "until dog -1 e2e.uptest.${cloudflare_zone_name} TXT @1.1.1.1; do sleep 1; done"

answer=$(dog -1 e2e.uptest."${cloudflare_zone_name}" TXT @1.1.1.1)
if [[ ${answer} != '"Hello, World from Uptest!"' ]]; then
	echo "DNS record e2e.uptest.${cloudflare_zone_name} is not set to 'Hello, World from Uptest!'"
	exit 1
fi
echo "DNS record e2e.uptest.${cloudflare_zone_name} verified"
