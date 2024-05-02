from utils import measurements_file, round_half_up

from typing import TypedDict


class CityReading(TypedDict):
  minimum: int
  maximum: int
  total: int
  count: int


if __name__ == "__main__":
  results: dict[str, CityReading] = {}

  with open(measurements_file) as measurements_file:
    for line in measurements_file:
      city, reading_str = line.strip().split(";")
      reading = float(reading_str)

      current = results.get(city, None)

      if current is None:
        results[city] = CityReading(
          minimum=reading,
          maximum=reading,
          total=reading,
          count=1,
        )
      else:
        current["minimum"] = min(current["minimum"], reading)
        current["maximum"] = max(current["maximum"], reading)
        current["total"] += reading
        current["count"] += 1
        results[city] = current

  cities = []
  for city in sorted(results.keys()):
    values = results[city]
    ans_min = round_half_up(values["minimum"])
    ans_avg = round_half_up(values["total"] / values["count"])
    ans_max = round_half_up(values["maximum"])
    cities.append(
      "%s=%.1f/%.1f/%.1f"
      % (
        city,
        ans_min,
        ans_avg,
        ans_max,
      )
    )

  print("{" + ", ".join(cities) + "}")
