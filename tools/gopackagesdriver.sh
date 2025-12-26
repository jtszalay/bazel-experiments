#!/usr/bin/env bash

# The following logic is done because we have multiple Bazel workspaces in the same root directory. We need to scope the bazel run below to match the workspace we're in.

# put bazel on the PATH
export PATH=$(git rev-parse --show-toplevel)/bin:$PATH

# Extract file path from arguments
file_path=""
for arg in "$@"; do
  if [[ "$arg" == file=* ]]; then
    file_path="${arg#file=}"
    break
  fi
done

# Find nearest MODULE.bazel starting from file path
if [ -n "$file_path" ]; then
  dir=$(dirname "$file_path")
  while [ "$dir" != "/" ]; do
    if [ -f "$dir/MODULE.bazel" ]; then
      cd "$dir"
      break
    fi
    dir=$(dirname "$dir")
  done
fi

exec bazel run --norun_validations -- @rules_go//go/tools/gopackagesdriver "${@}"