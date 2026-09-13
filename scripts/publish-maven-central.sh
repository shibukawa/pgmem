#!/usr/bin/env bash
# Uploads a bundle made by scripts/build-maven-bundle.sh to the Central
# Publisher Portal and waits for the outcome. CENTRAL_USERNAME and
# CENTRAL_PASSWORD are a Portal user token. --validate-only stops after the
# Portal's checks (namespace, POM, signatures) and drops the deployment
# whether it passed or failed, so nothing is published or left behind.
#
#   scripts/publish-maven-central.sh dist/maven/pgmem-0.1.0-bundle.zip
#   scripts/publish-maven-central.sh --validate-only dist/maven/pgmem-0.1.0-bundle.zip
set -euo pipefail
validate_only=false
if [ "${1:-}" = --validate-only ]; then
  validate_only=true
  shift
fi
bundle=${1:?usage: scripts/publish-maven-central.sh [--validate-only] bundle.zip}

api=https://central.sonatype.com/api/v1/publisher
token=$(printf '%s:%s' "${CENTRAL_USERNAME:?}" "${CENTRAL_PASSWORD:?}" | base64 | tr -d '\n')
if [ -n "${GITHUB_ACTIONS:-}" ]; then echo "::add-mask::$token"; fi
auth="Authorization: Bearer $token"
type=AUTOMATIC
if $validate_only; then type=USER_MANAGED; fi

id=$(curl --fail-with-body -sS -H "$auth" \
  -F "bundle=@$bundle;type=application/octet-stream" \
  "$api/upload?name=$(basename "$bundle" .zip)&publishingType=$type")
echo "deployment $id ($type)"

deadline=$((SECONDS + 3600))
while :; do
  status=$(curl --fail-with-body -sS -X POST -H "$auth" "$api/status?id=$id")
  state=$(printf '%s' "$status" | python3 -c 'import json, sys; print(json.load(sys.stdin)["deploymentState"])')
  echo "$(date -u +%H:%M:%S) $state"
  case $state in
    PUBLISHED)
      exit 0 ;;
    VALIDATED)
      if $validate_only; then
        curl --fail-with-body -sS -X DELETE -H "$auth" "$api/deployment/$id"
        echo "validation passed; deployment dropped"
        exit 0
      fi ;;
    FAILED)
      printf '%s\n' "$status" | python3 -m json.tool
      if $validate_only; then
        curl --fail-with-body -sS -X DELETE -H "$auth" "$api/deployment/$id"
        echo "deployment dropped"
      fi
      exit 1 ;;
  esac
  if [ $SECONDS -ge $deadline ]; then
    echo "gave up waiting for deployment $id"
    exit 1
  fi
  sleep 15
done
