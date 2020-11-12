#!/bin/bash
#
# Do a complete release. Run on CI.
# Upload assets, run goreleaser, and notify Tilt Cloud of the new release binaries.

set -ex

if [[ "$GITHUB_TOKEN" == "" ]]; then
    echo "Missing GITHUB_TOKEN"
    exit 1
fi

if [[ "$DOCKER_TOKEN" == "" ]]; then
    echo "Missing DOCKER_TOKEN"
    exit 1
fi

if [[ "$TILT_CLOUD_TOKEN" == "" ]]; then
    echo "Missing Tilt release token"
    exit 1
fi

DIR=$(dirname "$0")
cd "$DIR/.."

echo "$DOCKER_TOKEN" | docker login --username "$DOCKER_USERNAME" --password-stdin

mkdir -p ~/.tilt-dev
echo "$TILT_CLOUD_TOKEN" > ~/.tilt-dev/token

git fetch --tags
./scripts/upload-assets.py latest
goreleaser --rm-dist --skip-publish --snapshot

VERSION=$(git describe --abbrev=0 --tags)

./scripts/release-update-tilt-repo.sh "$VERSION"
./scripts/release-update-tilt-docs-repo.sh "$VERSION"
./scripts/record-release.sh "$VERSION"
