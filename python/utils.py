import math
import sys

if not len(sys.argv) > 1:
  print("Measurements file path argument required!")
  sys.exit(1)

measurements_file = sys.argv[1]


# https://realpython.com/python-rounding/#rounding-half-up
def round_half_up(n, decimals=1):
  multiplier = 10**decimals
  return math.floor(n * multiplier + 0.5) / multiplier
