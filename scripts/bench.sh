#!/bin/bash

set -e

npm run build

for measurements_file in $(ls -r measurements*.txt)
do
  echo "Running benchmarks for $measurements_file"

  hyperfine \
    --min-runs 1 \
    --warmup 1 \
    --parameter-list input_file $measurements_file \
    --export-markdown comparison_$measurements_file.md \
    -n 'python bruteforce {input_file}' 'python python/brute_force.py {input_file}' \
    -n 'bun bruteforce {input_file}' 'bun run node/brute-force.ts {input_file}' \
    -n 'node bruteforce {input_file}' 'node node/dist/brute-force.js {input_file}'

done
