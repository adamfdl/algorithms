package calculator

import "testing"

func TestCalculator(t *testing.T) {
	testSuites := []struct {
		expectedResult int
		cmd            calculator
	}{
		{
			cmd:            addCommand{1, 2},
			expectedResult: 3,
		},
		{
			cmd:            addCommand{2, 3},
			expectedResult: 5,
		},
		{
			cmd:            multiplyCommand{2, 3},
			expectedResult: 6,
		},
		{
			cmd:            multiplyCommand{10, 10},
			expectedResult: 100,
		},
	}

	for _, ts := range testSuites {
		actualResult := ts.cmd.calculate()
		if actualResult != ts.expectedResult {
			t.Errorf("expected %v, got %v", ts.expectedResult, actualResult)
		}
	}

}
