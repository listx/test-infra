#!/usr/bin/env bash

set -euo pipefail

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd -P)"
cd "${REPO_ROOT}"

echo "Ensuring go version."
source ./hack/build/setup-go.sh

_bin/protoc/bin/protoc \
    "--proto_path=${REPO_ROOT}/_bin/protoc/include/google/protobuf" \
    "--proto_path=${REPO_ROOT}/_bin/protoc/include/googleapis" \
    "--proto_path=${REPO_ROOT}/prow/gangway" \
    --include_imports \
    --include_source_info \
    --descriptor_set_out prow/gangway/api_descriptor.pb \
    gangway.proto
