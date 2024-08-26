# The One Billion Row Challenge

Learn more: https://github.com/gunnarmorling/1brc

## Getting Started

### 1. Create Measurements

```bash
bash scripts/create_measurements.sh 1000000
```

Above example will create measurements file `measurements-1000000.txt`

### 2. Run baseline

```bash
bash scripts/calculate_average_baseline_original_rounding.sh measurements-1000000.txt
```

Baseline average calculation output saved to `average_baseline.txt`

### 3. Run implementation

```bash
# Golang
go run ./go/main.go --file measurements-1000000.txt --impl track_aggregates_v2 > average.txt

# Bun (TypeScript)
bun run node/brute-force.ts measurements-1000000.txt > average.txt

# NodeJS (Javascript)
npm run build
node ./node/dist/v1.js measurements-1000000.txt > average.txt

# Python
python ./python/brute_force.py measurements-1000000.txt > average.txt
```

### 4. Test implementation output against baseline

```bash
bash scripts/test.sh
```

If output is incorrect, test script will exit with non-zero exit code and diff output from baseline.

## Results Comparison

### Run benchmark script

The [benchmark script](scripts/bench.sh) runs all implementations against different number of rows of inputs (10, 100, 1k, 10k, 100k, etc) using [hyperfine](https://github.com/sharkdp/hyperfine) to compare runtime among different implementations.

```bash
bash scripts/bench.sh
```

Refer to [results](results/) directory for comparison results.
