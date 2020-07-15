package fibonacci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacciSequence(t *testing.T) {
	testSuites := []struct {
		expectedSequence []int
		input            int
	}{
		{
			expectedSequence: []int{},
			input:            0,
		},
		{
			expectedSequence: []int{0},
			input:            1,
		},
		{
			expectedSequence: []int{0, 1},
			input:            2,
		},
		{
			expectedSequence: []int{0, 1, 1},
			input:            3,
		},
		{
			expectedSequence: []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34},
			input:            9,
		},
	}

	for _, ts := range testSuites {
		actualSequence := Sequence(ts.input)
		assert.Equal(t, ts.expectedSequence, actualSequence)
	}

}
