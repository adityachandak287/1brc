package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Results struct {
	readings map[string][]float64
}

func NewResult() Results {
	return Results{
		readings: map[string][]float64{},
	}
}

func (res *Results) record(line string) {
	city, reading := parseLine(line)

	current, ok := res.readings[city]
	if !ok {
		current = []float64{reading}
	} else {
		current = append(current, reading)
	}
	res.readings[city] = current
}

func (res *Results) process() string {
	var cities []string
	for city := range res.readings {
		cities = append(cities, city)
	}
	sort.Strings(cities)

	output := make([]string, len(cities))

	for idx, city := range cities {
		readings := res.readings[city]

		min := math.MaxFloat64
		max := -1 * math.MaxFloat64

		total := float64(0)
		count := len(readings)

		for _, reading := range readings {
			if reading < min {
				min = reading
			}
			if reading > max {
				max = reading
			}

			total += reading
		}
		avg := total / float64(count)

		min = roundHalfUp(min, 1)
		max = roundHalfUp(max, 1)
		avg = roundHalfUp(avg, 1)

		output[idx] = fmt.Sprintf("%s=%.1f/%.1f/%.1f", city, min, avg, max)
	}

	return fmt.Sprintf("{%s}", strings.Join(output, ", "))

}

func parseLine(line string) (string, float64) {
	parts := strings.SplitN(line, ";", 2)

	if len(parts) != 2 {
		log.Panic("Failed to split line", parts)
	}

	city := parts[0]
	readingStr := parts[1]

	reading, err := strconv.ParseFloat(readingStr, 64)
	if err != nil {
		log.Panic("Could not parse string into float", readingStr)
	}

	return city, reading
}

var linePattern = regexp.MustCompile(`(.+);([-]?\d{1,2}\.\d)`)

func parseLineRegex(line string) (string, float64) {
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

func BruteForceSolution(inputFile string) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	results := NewResult()

	for scanner.Scan() {
		results.record(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(results.process())
}
