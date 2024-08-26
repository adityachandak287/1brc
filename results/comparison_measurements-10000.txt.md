| Command                                    |  Mean [ms] | Min [ms] | Max [ms] |     Relative |
| :----------------------------------------- | ---------: | -------: | -------: | -----------: |
| `python bruteforce measurements-10000.txt` | 70.2 ± 0.8 |     68.5 |     73.0 | 21.75 ± 0.93 |
| `bun bruteforce measurements-10000.txt`    | 38.2 ± 0.9 |     36.7 |     40.3 | 11.84 ± 0.56 |
| `node bruteforce measurements-10000.txt`   | 45.5 ± 0.7 |     44.3 |     48.5 | 14.11 ± 0.62 |
| `bun v1 measurements-10000.txt`            | 37.6 ± 1.9 |     35.4 |     51.7 | 11.64 ± 0.75 |
| `python v1 measurements-10000.txt`         | 84.8 ± 1.3 |     83.4 |     90.6 | 26.29 ± 1.15 |
| `node v1 measurements-10000.txt`           | 44.5 ± 1.5 |     43.2 |     54.7 | 13.81 ± 0.73 |
| `go brute_force`                           |  3.6 ± 0.2 |      3.4 |      5.9 |  1.10 ± 0.07 |
| `go track_aggregates`                      |  3.4 ± 0.2 |      3.2 |      5.3 |  1.05 ± 0.06 |
| `go track_aggregates_v2`                   |  3.2 ± 0.1 |      3.0 |      5.6 |         1.00 |
