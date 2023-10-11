package longcons

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongCons(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output int
	}{
		{"Case 1", []int{100, 4, 200, 1, 3, 2}, 4},
		{"Case 2", []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
		{"Empty", []int{}, 0},
		{"Single elements", []int{1, 3, 5}, 1},
		{"All same", []int{1, 1, 1, 1, 1, 1}, 1},
		{"All same + 1", []int{1, 1, 1, 1, 1, 1, 2}, 2},
		{"Negative and positive", []int{-7, -1, 3, -9, -4, 7, -3, 2, 4, 9, 4, -9, 8, -7, 5, -1, -7}, 4},
		{"Case 3", []int{-2, -3, -3, 7, -3, 0, 5, 0, -8, -4, -1, 2}, 5},
		{"Case 4", []int{-4, -1, 4, -5, 1, -6, 9, -6, 0, 2, 2, 7, 0, 9, -3, 8, 9, -2, -6, 5, 0, 3, 4, -2}, 12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := longestConsecutive(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}
