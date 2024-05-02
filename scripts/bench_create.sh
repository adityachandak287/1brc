#!/bin/bash

set -e

for num_measurements in 1 10 100 1000 10000 100000 1000000 10000000 100000000 1000000000
do
  echo "Running benchmarks for $num_measurements"

  hyperfine \
    --min-runs 3 \
    --warmup 1 \
    --parameter-list create_method create_measurements.sh,create_measurements3.sh,create_measurements_fast.sh \
    --parameter-list num_measurements $num_measurements \
    -n '{create_method} {num_measurements}' 'bash scripts/{create_method} {num_measurements}'

done
