#!/bin/bash

set -e

## Export variables
#CONF=test/conf/env.sh
#if [ -f "$CONF" ]; then
#    . "$CONF"
#fi

# Get test names from args, run all if empty
TESTS="$1"
if [ -z "$TESTS" ]; then
    TESTS=("controlplane")
fi

# Run tests
for TEST in ${TESTS[@]}; do
    bats "bats/${TEST}.bats"
done