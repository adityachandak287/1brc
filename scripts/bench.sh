#!/bin/bash

set -e

for measurements_file in $(ls -r measurements*.txt)
do
  echo "Running benchmarks for $measurements_file"

  hyperfine \
    --min-runs 1 \
    --warmup 1 \
    --parameter-list input_file $measurements_file \
    -n 'python bruteforce {input_file}' 'python python/brute_force.py {input_file}' \
    -n 'bun bruteforce {input_file}' 'bun run node/brute-force.ts {input_file}'

done
