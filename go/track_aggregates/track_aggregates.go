package track_aggregates

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"

	"github.com/adityachandak287/1brc/go/utils"
)

type CityAgg struct {
	min   float64
	max   float64
	total float64
	count float64
}

type Results struct {
	readings map[string]CityAgg
}

func NewResult() Results {
	return Results{
		readings: make(map[string]CityAgg),
	}
}

func (res *Results) record(line string) {
	city, reading := utils.ParseLine(line)

	current, ok := res.readings[city]
	if !ok {
		current = CityAgg{
			min:   reading,
			max:   reading,
			total: reading,
			count: 1,
		}
	} else {
		current.min = math.Min(current.min, reading)
		current.max = math.Max(current.max, reading)
		current.total += reading
		current.count += 1
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
		cityAgg := res.readings[city]

		min := utils.RoundHalfUp(cityAgg.min, 1)
		max := utils.RoundHalfUp(cityAgg.max, 1)
		avg := utils.RoundHalfUp(cityAgg.total/cityAgg.count, 1)

		output[idx] = fmt.Sprintf("%s=%.1f/%.1f/%.1f", city, min, avg, max)
	}

	return fmt.Sprintf("{%s}", strings.Join(output, ", "))
}

func Solution(inputFile string) {
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

const WRITE_BATCH_SIZE = 1000

func writer(scanner *bufio.Scanner, lineCh chan<- []string) {
	lines := make([]string, WRITE_BATCH_SIZE)

	lineIdx := 0
	for scanner.Scan() {
		lines[lineIdx] = scanner.Text()
		lineIdx += 1
		if lineIdx == WRITE_BATCH_SIZE {
			batch := make([]string, WRITE_BATCH_SIZE)
			copy(batch, lines[0:lineIdx])
			lineCh <- batch
			lineIdx = 0
		}
	}

	if lineIdx > 0 {
		lineCh <- lines[:lineIdx]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	close(lineCh)
}

func reader(results Results, lineCh <-chan []string, doneCh chan<- bool) {
	for lines := range lineCh {
		for _, line := range lines {
			results.record(line)
		}
	}

	doneCh <- true
}

// Spawn 1 batch writer and batch reader goroutine
func SolutionV2(inputFile string) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineCh := make(chan []string, 1024)
	go writer(scanner, lineCh)

	results := NewResult()
	doneCh := make(chan bool)
	go reader(results, lineCh, doneCh)

	<-doneCh

	fmt.Println(results.process())
}
