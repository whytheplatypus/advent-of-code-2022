package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(input))

	grid := [][]int{
		[]int{},
	}

	for _, c := range input {
		if c == '\n' {
			grid = append(grid, []int{})
			continue
		}
		height, err := strconv.Atoi(string(c))
		if err != nil {
			log.Fatal(err)
		}
		grid[len(grid)-1] = append(grid[len(grid)-1], height)
	}
	// pop empty row
	grid = grid[:len(grid)-1]
	render(grid)

	total := 0
	for y, row := range grid {
		for x, _ := range row {
			if visible(grid, x, y) {
				total++
			}
		}
	}
	log.Println(total)
}

func render(grid [][]int) {
	for row, i := range grid {
		fmt.Printf("%d: ", row)
		for _, k := range i {
			fmt.Printf("%d ", k)
		}
		fmt.Println("")
	}
}

func visible(grid [][]int, x, y int) bool {
	visible := true
	height := grid[y][x]

	//from left
	for _, h := range grid[y][:x] {
		visible = visible && height > h
	}
	if visible {
		log.Println("visible from left")
		return visible
	}
	visible = true
	//from right
	for i := len(grid[y]) - 1; i > x; i-- {
		h := grid[y][i]
		visible = visible && height > h
	}
	if visible {
		log.Println("visible from right")
		return visible
	}

	visible = true
	//from top
	for i := 0; i < y; i++ {
		h := grid[i][x]
		visible = visible && height > h
	}
	if visible {
		log.Println("visible from top")
		return visible
	}

	visible = true
	//from bottom
	for i := len(grid) - 1; i > y; i-- {
		h := grid[i][x]
		visible = visible && height > h
	}
	if visible {
		log.Println("visible from bottom")
	}
	return visible
}
