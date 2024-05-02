#!/bin/bash

set -e

NUM_MEASUREMENTS=${1:-1000}
MEASUREMENTS_FILE="measurements-$NUM_MEASUREMENTS.txt"

if ! [ -f $MEASUREMENTS_FILE ];
then
  echo "INFO: Measurements file [$MEASUREMENTS_FILE] does not exist, creating."
  touch $MEASUREMENTS_FILE
else
  echo "WARN: Measurements file [$MEASUREMENTS_FILE] already exists, it will be overwritten."
fi

# Run docker container with JDK 21 to compile and run the `CreateMeasurements.java` file
docker run \
  --rm \
  --workdir /scripts \
  -v "$PWD/$MEASUREMENTS_FILE:/scripts/measurements.txt:rw" \
  -v "$PWD/scripts/CreateMeasurements.java:/scripts/CreateMeasurements.java:ro" \
  amazoncorretto:21 \
  bash -c "javac CreateMeasurements.java && java CreateMeasurements.java ${NUM_MEASUREMENTS}"

# Display measurements file size
du -sh $MEASUREMENTS_FILE
