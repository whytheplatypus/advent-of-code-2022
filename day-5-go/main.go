package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type stack []string

func (s *stack) pop() string {
	crate := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return crate
}

func (s *stack) append(crate string) {
	*s = append(*s, crate)
}

type yard map[string]*stack

func (y yard) move(num int, from, to string) {
	for i := 0; i < num; i++ {
		crate := y[from].pop()
		y[to].append(crate)
	}
}

func splitDrawing(drawing string) (string, string) {
	parts := strings.Split(drawing, "\n\n")
	return parts[0], parts[1]
}

var yardRegex = regexp.MustCompile(`(?:\s|\[)(.)(?:\s|\])(?:\s|$)`)

func loadYard(rawStart string) yard {
	y := yard{}
	drawing := strings.Split(rawStart, "\n")
	rows := strings.Split(drawing[len(drawing)-1], "   ")
	for _, row := range rows {
		y[strings.TrimSpace(row)] = &stack{}
	}

	drawing = drawing[:len(drawing)-1]

	for i := len(drawing) - 1; i > -1; i-- {
		line := drawing[i]
		rawCrates := yardRegex.FindAllStringSubmatch(line, -1)
		for k, crate := range rawCrates {
			row := strings.TrimSpace(rows[k])
			if crate[1] == " " {
				continue
			}
			y[row].append(crate[1])
		}
	}
	return y
}

var moveRegex = regexp.MustCompile(`^move (\d+) from (\d) to (\d)$`)

func parseMove(rawMove string) (num int, from string, to string) {
	parsedMove := moveRegex.FindAllStringSubmatch(rawMove, -1)
	var err error
	num, err = strconv.Atoi(parsedMove[0][1])
	if err != nil {
		log.Fatal(err)
	}
	from = parsedMove[0][2]
	to = parsedMove[0][3]
	return
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
	rawStart, rawMoves := splitDrawing(string(rawProblem))
	log.Println(rawStart)
	y := loadYard(rawStart)

	moves := strings.Split(rawMoves, "\n")
	for _, move := range moves {
		if move == "" {
			continue
		}
		num, from, to := parseMove(move)
		y.move(num, from, to)
	}

	//lol needed to get this result in order
	for i := 0; i < len(y); i++ {
		s := y[fmt.Sprintf("%d", i+1)]
		fmt.Print((*s)[len((*s))-1])
	}
}
