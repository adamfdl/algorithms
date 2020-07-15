package prime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNPrimeSequence(t *testing.T) {
	testSuites := []struct {
		input            int
		expectedSequence []int
	}{
		{
			input:            5,
			expectedSequence: []int{2, 3, 5, 7, 11},
		},
	}

	for _, ts := range testSuites {
		actualSequence := firstNPrimeSequence(ts.input)
		assert.Equal(t, ts.expectedSequence, actualSequence)
	}
}
