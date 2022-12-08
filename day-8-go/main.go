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

	highScore := 0
	for y, row := range grid {
		fmt.Printf("%d: ", y)
		for x, h := range row {
			score := score(grid, x, y)
			if score > highScore {
				fmt.Printf("!")
				highScore = score
			}
			fmt.Printf("%d ", h)
		}
		fmt.Println("")
	}
	log.Println(highScore)
}

func render(grid [][]int) {
	for row, i := range grid {
		fmt.Printf("%d: ", row)
		for _, k := range i {
			fmt.Printf("%d ", k)
		}
		fmt.Println("")
	}
	fmt.Println("---\n")
}

func score(grid [][]int, x, y int) int {
	height := grid[y][x]
	score := 0

	//left
	for i := x - 1; i > -1; i-- {
		score++
		h := grid[y][i]
		if height <= h {
			break
		}
	}

	//from right
	count := 0
	for i := x + 1; i < len(grid[y]); i++ {
		count++
		h := grid[y][i]
		if height <= h {
			break
		}
	}
	score *= count

	count = 0
	//from top
	for i := y - 1; i > -1; i-- {
		count++
		h := grid[i][x]
		if height <= h {
			break
		}
	}
	score *= count

	count = 0
	//from bottom
	for i := y + 1; i < len(grid); i++ {
		count++
		h := grid[i][x]
		if height <= h {
			break
		}
	}
	score *= count
	return score
}
