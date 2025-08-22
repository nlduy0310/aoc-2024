package stone

import "testing"

func TestChange(t *testing.T) {

	type TestCase struct {
		input          int
		expectedOutput []int
	}

	testCases := []TestCase{
		{input: 0, expectedOutput: []int{1}},
		{input: 1, expectedOutput: []int{2024}},
		{input: 10, expectedOutput: []int{1, 0}},
		{input: 99, expectedOutput: []int{9, 9}},
		{input: 999, expectedOutput: []int{2021976}},
	}

	for _, testCase := range testCases {
		actualOutput := Stone{Val: testCase.input}.Change()
		if len(testCase.expectedOutput) != len(actualOutput) {
			t.Errorf("Stone.Change() test failed, input=%d, expected output length=%d, actual output length=%d", testCase.input, len(testCase.expectedOutput), len(actualOutput))
		}

		for idx, stone := range actualOutput {
			if testCase.expectedOutput[idx] != stone.Val {
				t.Errorf("Stone.Change() test failed, input=%d, expected value at index %d to be %d, got %d", testCase.input, idx, testCase.expectedOutput[idx], stone.Val)
			}
		}
	}
}
