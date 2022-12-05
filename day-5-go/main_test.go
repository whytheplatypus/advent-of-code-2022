package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPop(t *testing.T) {
	example := stack{"a", "b", "c"}
	crate := example.pop()
	if example[len(example)-1] != "b" {
		t.Fatal("Failed", example)
	}
	if crate != "c" {
		t.Fatal("Failed", crate)
	}
	t.Log(example)
	t.Log(crate)
}

func TestAppend(t *testing.T) {
	example := stack{"a", "b"}
	example.append("c")
	t.Log(example)
	assert.Equal(t, example, stack{"a", "b", "c"})
}

func TestMove(t *testing.T) {
	example := yard{
		"1": &stack{"A", "B"},
		"2": &stack{},
	}

	example.move(1, "1", "2")

	t.Log(example)

	assert.Equal(t, example, yard{
		"1": &stack{"A"},
		"2": &stack{"B"},
	})
}

func TestRegex(t *testing.T) {
	example := "    [D]    "
	result := yardRegex.FindAllStringSubmatch(example, -1)
	t.Logf("%q\n", result)
	assert.Equal(t, result, [][]string{
		[]string{"    ", " "},
		[]string{"[D] ", "D"},
		[]string{"   ", " "}})
}

func TestLoadYard(t *testing.T) {
	input := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 `
	result := loadYard(input)
	assert.Equal(t, yard{
		"1": &stack{"Z", "N"},
		"2": &stack{"M", "C", "D"},
		"3": &stack{"P"},
	}, result)
}

func TestParseMove(t *testing.T) {
	input := "move 1 from 2 to 1"
	num, from, to := parseMove(input)
	assert.Equal(t, 1, num)
	assert.Equal(t, "2", from)
	assert.Equal(t, "1", to)
}
