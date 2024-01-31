if (process.argv.length <= 2) {
  console.error("Measurements file path argument required!");
  process.exit(0);
}

export const filePath = process.argv.at(2) as string;

// https://realpython.com/python-rounding/#rounding-half-up
export const roundHalfUp = (n: number, decimals = 1) => {
  const mult = Math.pow(10, decimals);
  return Math.floor(n * mult + 0.5) / mult;
};
