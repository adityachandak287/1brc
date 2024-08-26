package utils

import (
	"log"
	"math"
	"regexp"
	"strconv"
)

// https://realpython.com/python-rounding/#rounding-half-up
func RoundHalfUp(n float64, decimals int) float64 {
	mult := math.Pow(10, float64(decimals))
	return math.Floor(n*mult+0.5) / mult
}

func ParseLine(line string) (string, float64) {
	semiIdx := -1
	for idx, ch := range line {
		if ch == ';' {
			semiIdx = idx
			break
		}
	}

	readingStr := line[semiIdx+1:]
	reading, err := strconv.ParseFloat(readingStr, 64)
	if err != nil {
		log.Panic("Could not parse string into float", readingStr)
	}

	return line[:semiIdx], reading
}

var linePattern = regexp.MustCompile(`(.+);([-]?\d{1,2}\.\d)`)

func ParseLineRegex(line string) (string, float64) {
	parts := linePattern.FindStringSubmatch(line)

	if len(parts) != 3 {
		log.Panic("Failed to split line", parts)
	}

	city := parts[1]
	readingStr := parts[2]

	reading, err := strconv.ParseFloat(readingStr, 64)
	if err != nil {
		log.Panic("Could not parse string into float", readingStr)
	}

	return city, reading
}
