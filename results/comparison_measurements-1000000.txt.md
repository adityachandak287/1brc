| Command                                      |    Mean [ms] | Min [ms] | Max [ms] |    Relative |
| :------------------------------------------- | -----------: | -------: | -------: | ----------: |
| `python bruteforce measurements-1000000.txt` | 984.0 ± 11.1 |    975.3 |    996.5 | 8.16 ± 0.15 |
| `bun bruteforce measurements-1000000.txt`    | 498.3 ± 29.9 |    476.2 |    557.0 | 4.13 ± 0.26 |
| `node bruteforce measurements-1000000.txt`   | 677.3 ± 10.5 |    662.8 |    686.3 | 5.62 ± 0.12 |
| `bun v1 measurements-1000000.txt`            |  477.0 ± 5.2 |    469.9 |    481.7 | 3.96 ± 0.07 |
| `python v1 measurements-1000000.txt`         |       1635.1 |   1635.1 |   1635.1 |       13.56 |
| `node v1 measurements-1000000.txt`           |  658.3 ± 2.7 |    655.0 |    661.3 | 5.46 ± 0.08 |
| `go brute_force`                             |  165.1 ± 2.1 |    161.9 |    169.7 | 1.37 ± 0.03 |
| `go track_aggregates`                        |  165.5 ± 2.1 |    162.1 |    170.0 | 1.37 ± 0.03 |
| `go track_aggregates_v2`                     |  120.6 ± 1.8 |    118.5 |    126.3 |        1.00 |
