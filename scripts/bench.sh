#!/bin/bash

set -e

for measurements_file in $(ls -r measurements*.txt)
do
  echo "Running benchmarks for $measurements_file"

  hyperfine \
    --min-runs 1 \
    --warmup 1 \
    --parameter-list implementation brute_force.py \
    --parameter-list input_file $measurements_file \
    -n '{implementation} {input_file}' \
    'python python/{implementation} {input_file}'

done
