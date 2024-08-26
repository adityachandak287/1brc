package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
	"runtime/trace"

	bruteforce "github.com/adityachandak287/1brc/go/brute_force"
	track_aggregates "github.com/adityachandak287/1brc/go/track_aggregates"
)

func main() {
	inputFile := flag.String("file", "", "Measurements input file")
	impl := flag.String("impl", "", "Implementation to use [brute_force, track_aggregates, track_aggregates_v2]")
	cpuprofile := flag.String("cpuprofile", "", "Write cpu profile to file")
	tracefile := flag.String("trace", "", "Write traces to file")

	flag.Parse()

	if *inputFile == "" {
		log.Panic("Measurements input file requred!")
	}

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		err = pprof.StartCPUProfile(f)
		if err != nil {
			log.Fatal(err)
		}
		defer pprof.StopCPUProfile()
	}

	if *tracefile != "" {
		f, err := os.Create(*tracefile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		err = trace.Start(f)
		if err != nil {
			log.Fatal(err)
		}
		defer trace.Stop()
	}

	switch *impl {
	case "brute_force":
		bruteforce.Solution(*inputFile)
	case "track_aggregates":
		track_aggregates.Solution(*inputFile)
	case "track_aggregates_v2":
		track_aggregates.SolutionV2(*inputFile)
	default:
		log.Panic("Invalid input for impl!")
	}
}
