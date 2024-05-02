import { createInterface } from "readline";
import { createReadStream, read } from "fs";
import { filePath, roundHalfUp } from "./utils";

const fileStream = createReadStream(filePath);

const input = createInterface({
  input: fileStream,
});

type CityReading = { min: number; max: number; total: number; count: number };

const readings = new Map<string, CityReading>();

const parseLine = (line: string) => {
  const parts = line.split(";");
  const city = parts[0];
  const reading = Number(parts[1]);

  const current = readings.get(city);

  if (!current) {
    readings.set(city, {
      min: reading,
      max: reading,
      total: reading,
      count: 1,
    });
  } else {
    readings.set(city, {
      min: Math.min(current.min, reading),
      max: Math.max(current.max, reading),
      total: current.count + reading,
      count: current.count + 1,
    });
  }
};

const processResults = () => {
  const results = [...readings.keys()].sort().map((city) => {
    const values = readings.get(city) as CityReading;

    const min = roundHalfUp(values.min);
    const max = roundHalfUp(values.max);
    const avg = roundHalfUp(values.total / values.count);

    return `${city}=${min.toFixed(1)}/${avg.toFixed(1)}/${max.toFixed(1)}`;
  });

  console.log(`{${results.join(", ")}}`);
};

input.on("line", parseLine);
input.on("close", processResults);
