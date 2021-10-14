#!/usr/bin/env sh
set -eax

scripts/build.sh && \
git tag `bin/commit version` && \
git push origin --tags
