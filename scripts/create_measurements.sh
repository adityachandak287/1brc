#!/bin/bash

MEASUREMENTS_FILE="measurements.txt"
NUM_MEASUREMENTS=${1:-1000}

if ! [ -f $MEASUREMENTS_FILE ]; then
  echo "Measurements file does not exist, creating."
  touch $MEASUREMENTS_FILE
fi

docker run \
  --rm \
  -it \
  --workdir /scripts \
  -v "$PWD/$MEASUREMENTS_FILE:/scripts/$MEASUREMENTS_FILE" \
  -v "$PWD/scripts/CreateMeasurements.java:/scripts/CreateMeasurements.java" \
  amazoncorretto:21 \
  bash -c "javac CreateMeasurements.java && java CreateMeasurements.java ${NUM_MEASUREMENTS}"
