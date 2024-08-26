| Command                                     |   Mean [ms] | Min [ms] | Max [ms] |     Relative |
| :------------------------------------------ | ----------: | -------: | -------: | -----------: |
| `python bruteforce measurements-100000.txt` | 146.7 ± 1.7 |    143.7 |    149.5 | 10.26 ± 0.22 |
| `bun bruteforce measurements-100000.txt`    |  79.3 ± 2.1 |     76.7 |     89.2 |  5.55 ± 0.17 |
| `node bruteforce measurements-100000.txt`   | 114.6 ± 7.3 |    106.8 |    135.2 |  8.02 ± 0.53 |
| `bun v1 measurements-100000.txt`            |  79.9 ± 3.1 |     76.6 |     93.8 |  5.59 ± 0.24 |
| `python v1 measurements-100000.txt`         | 222.9 ± 7.6 |    215.0 |    245.0 | 15.60 ± 0.60 |
| `node v1 measurements-100000.txt`           | 107.3 ± 1.9 |    105.0 |    112.4 |  7.51 ± 0.19 |
| `go brute_force`                            |  20.6 ± 0.3 |     20.0 |     22.9 |  1.44 ± 0.03 |
| `go track_aggregates`                       |  18.1 ± 0.3 |     17.6 |     19.6 |  1.27 ± 0.03 |
| `go track_aggregates_v2`                    |  14.3 ± 0.2 |     13.8 |     16.4 |         1.00 |
