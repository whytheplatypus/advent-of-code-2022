package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessBuffer(t *testing.T) {
	tests := []struct {
		input  string
		result int
	}{
		{
			"bvwbjplbgvbhsrlpgdmjqwftvncz",
			5,
		},
		{
			"nppdvjthqldpwncqszvftbrmjlhg",
			6,
		},
		{
			"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			10,
		},
		{
			"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			11,
		},
	}

	for i, test := range tests {
		result := processBuffer(test.input)
		assert.Equal(t, test.result, result, "results should be equal for test", i)
	}
}
