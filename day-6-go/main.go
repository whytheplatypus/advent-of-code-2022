package main

import (
	"flag"
	"log"
	"os"
)

const MARKER_LEN = 4

type window map[string]struct{}

func (w window) check() bool {
	return len(w) == MARKER_LEN
}

func processBuffer(input string) int {
	processed := []string{}
	for _, c := range input {
		processed = append(processed, string(c))
		if len(processed) > MARKER_LEN-1 {
			win := window{}
			for _, p := range processed[len(processed)-MARKER_LEN : len(processed)] {
				win[p] = struct{}{}
			}
			if win.check() {
				return len(processed)
			}
		}
	}
	return 0
}

var input string

func init() {
	flag.StringVar(&input, "input", "test.txt", "path to the input file")
}

func main() {
	flag.Parse()
	rawProblem, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}

	count := processBuffer(string(rawProblem))
	log.Println(count)
}
