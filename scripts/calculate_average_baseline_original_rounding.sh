#!/bin/bash

set -e

MEASUREMENTS_FILE=${1:-measurements-1000.txt}
OUTPUT_FILE="average_baseline.txt"

echo "INFO: Running calculation for $MEASUREMENTS_FILE"
echo "INFO: Output will be written to $OUTPUT_FILE"

if ! [ -f $MEASUREMENTS_FILE ];
then
  echo "ERROR: Measurements file [$MEASUREMENTS_FILE] does not exist!"
  exit 1
fi

# Run docker container with JDK 21 to compile and run the `CalculateAverage_baseline_original_rounding.java` file
docker run \
  --rm \
  -it \
  --workdir /scripts \
  -v "$PWD/$MEASUREMENTS_FILE:/scripts/measurements.txt:rw" \
  -v "$PWD/scripts/CalculateAverage_baseline_original_rounding.java:/scripts/CalculateAverage_baseline_original_rounding.java:ro" \
  amazoncorretto:21 \
  bash -c "javac CalculateAverage_baseline_original_rounding.java && java CalculateAverage_baseline_original_rounding.java" > $OUTPUT_FILE

echo "INFO: Average calculation output saved to $OUTPUT_FILE"
