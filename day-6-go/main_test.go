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
			23,
		},
		{
			"nppdvjthqldpwncqszvftbrmjlhg",
			23,
		},
		{
			"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			29,
		},
		{
			"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			26,
		},
	}

	for i, test := range tests {
		result := processBuffer(test.input)
		assert.Equal(t, test.result, result, "results should be equal for test", i)
	}
}
