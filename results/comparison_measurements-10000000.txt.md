| Command                                       |      Mean [s] | Min [s] | Max [s] | Relative |
| :-------------------------------------------- | ------------: | ------: | ------: | -------: |
| `python bruteforce measurements-10000000.txt` |         9.387 |   9.387 |   9.387 |     7.82 |
| `bun bruteforce measurements-10000000.txt`    |         4.579 |   4.579 |   4.579 |     3.81 |
| `node bruteforce measurements-10000000.txt`   |         6.164 |   6.164 |   6.164 |     5.13 |
| `bun v1 measurements-10000000.txt`            |         4.406 |   4.406 |   4.406 |     3.67 |
| `python v1 measurements-10000000.txt`         |        15.472 |  15.472 |  15.472 |    12.88 |
| `node v1 measurements-10000000.txt`           |         6.095 |   6.095 |   6.095 |     5.07 |
| `go brute_force`                              |         1.647 |   1.647 |   1.647 |     1.37 |
| `go track_aggregates`                         |         1.623 |   1.623 |   1.623 |     1.35 |
| `go track_aggregates_v2`                      | 1.201 ± 0.005 |   1.197 |   1.205 |     1.00 |
