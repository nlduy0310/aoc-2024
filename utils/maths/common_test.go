package maths

import "testing"

func TestDigits(t *testing.T) {

	type TestCase struct {
		input, expectedOutput int
	}

	testCases := []TestCase{
		{input: 0, expectedOutput: 1},
		{input: 1, expectedOutput: 1},
		{input: 5, expectedOutput: 1},
		{input: 9, expectedOutput: 1},
		{input: 10, expectedOutput: 2},
		{input: 11, expectedOutput: 2},
		{input: 50, expectedOutput: 2},
		{input: 99, expectedOutput: 2},
		{input: 100, expectedOutput: 3},
		{input: 101, expectedOutput: 3},
		{input: -0, expectedOutput: 1},
		{input: -1, expectedOutput: 1},
		{input: -5, expectedOutput: 1},
		{input: -9, expectedOutput: 1},
		{input: -10, expectedOutput: 2},
		{input: -11, expectedOutput: 2},
		{input: -50, expectedOutput: 2},
		{input: -99, expectedOutput: 2},
		{input: -100, expectedOutput: 3},
		{input: -101, expectedOutput: 3},
	}

	for _, testCase := range testCases {
		actualOutput := Digits(testCase.input)

		if actualOutput != testCase.expectedOutput {
			t.Errorf("Digits() test failed, input=%d, expectedOutput=%d, received=%d", testCase.input, testCase.expectedOutput, actualOutput)
		}
	}
}
