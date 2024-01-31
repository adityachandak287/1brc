# https://github.com/gunnarmorling/1brc/blob/a8823a1f93e63999dfeec159cdfa21eac43a9cab/test.sh#L48

diff \
  --color=always \
  --strip-trailing-cr \
  <(bash ./scripts/tocsv.sh < average_baseline.txt) <(bash ./scripts/tocsv.sh < average.txt)
