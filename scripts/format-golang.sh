#!/bin/bash
set -e
if [ "$#" -eq 0 ]; then
  exit 0
fi

# Format files
gofmt -w "$@"
# Optionally, run your linter on just the staged files
# golangci-lint run "$@"
