package operator

import "testing"

func TestApplyOpConcatenate(t *testing.T) {

	type TestCase struct {
		firstOperand   int
		secondOperand  int
		expectedResult int
	}

	positiveTestCases := []TestCase{
		{0, 1, 1},
		{0, 2, 2},
		{1, 0, 10},
		{2, 0, 20},
		{123, 456, 123456},
		{1, 23456, 123456},
		{12345, 6, 123456},
	}

	for _, testCase := range positiveTestCases {
		actualOutput := OpConcatenate.Apply(testCase.firstOperand, testCase.secondOperand)

		if actualOutput != testCase.expectedResult {
			t.Errorf("test apply OpConcatenate failed: %d concat %d gives %d, expected %d", testCase.firstOperand, testCase.secondOperand, actualOutput, testCase.expectedResult)
		}
	}
}
