from utils import measurements_file, round_half_up

if __name__ == "__main__":
  results = {}

  with open(measurements_file) as measurements_file:
    for line in measurements_file:
      city, reading_str = line.strip().split(";")
      reading = float(reading_str)

      if city not in results:
        results[city] = [float(reading)]
      else:
        results[city].append(float(reading))

  averages = []
  for city in sorted(results.keys()):
    ans_min = round_half_up(min(results[city]))
    ans_avg = round_half_up(sum(results[city]) / len(results[city]))
    ans_max = round_half_up(max(results[city]))
    averages.append(
      "%s=%.1f/%.1f/%.1f"
      % (
        city,
        ans_min,
        ans_avg,
        ans_max,
      )
    )

  print("{" + ", ".join(averages) + "}")
