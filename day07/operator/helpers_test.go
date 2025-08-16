package operator

import (
	"testing"
)

func TestDigitsCount(t *testing.T) {

	type TestCase struct {
		input          int
		expectedOutput int
	}

	testCases := []TestCase{
		{0, 1},
		{-0, 1},
		{1, 1},
		{-1, 1},
		{5, 1},
		{-5, 1},
		{10, 2},
		{-10, 2},
		{50, 2},
		{-50, 2},
		{99, 2},
		{-99, 2},
		{101, 3},
		{-101, 3},
		{1000, 4},
		{-1000, 4},
	}

	for _, testCase := range testCases {
		actualOutput := digits(testCase.input)

		if actualOutput != testCase.expectedOutput {
			t.Errorf("test digits count failed: input=%d, expectedOutput=%d, actualOutput=%d", testCase.input, testCase.expectedOutput, actualOutput)
		}
	}
}
