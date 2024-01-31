import { createInterface } from "readline";
import { createReadStream } from "fs";
import { filePath, roundHalfUp } from "./utils";

const fileStream = createReadStream(filePath);

const input = createInterface({
  input: fileStream,
  crlfDelay: Infinity,
  terminal: false,
});

const readings = new Map<string, number[]>();

const parseLine = (line: string) => {
  const parts = line.split(";");
  const city = parts[0];
  const reading = Number(parts[1]);

  if (readings.has(city)) {
    readings.get(city)?.push(reading);
  } else {
    readings.set(city, [reading]);
  }
};

const processResults = () => {
  const results = [...readings.keys()].sort().map((city) => {
    let min = Number.MAX_SAFE_INTEGER;
    let max = Number.MIN_SAFE_INTEGER;
    let total = 0;

    const values = readings.get(city) || [];
    for (const val of values) {
      min = Math.min(val, min);
      max = Math.max(val, max);
      total += val;
    }

    min = roundHalfUp(min);
    max = roundHalfUp(max);
    const avg = roundHalfUp(total / values.length);

    return `${city}=${min.toFixed(1)}/${avg.toFixed(1)}/${max.toFixed(1)}`;
  });

  console.log(`{${results.join(", ")}}`);
};

input.on("line", parseLine);
input.on("close", processResults);
