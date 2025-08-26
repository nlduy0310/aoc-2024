package position

import "testing"

func TestClampX(t *testing.T) {

	type TestCase struct {
		initial            int
		rangeMin, rangeMax int
		expected           int
	}

	testCases := []TestCase{
		{1, 5, 7, 7},
		{2, 5, 7, 5},
		{3, 5, 7, 6},
		{9, 5, 7, 6},
		{10, 5, 7, 7},
		{11, 5, 7, 5},
		{5, 5, 5, 5},
		{6, 5, 5, 5},
		{4, 5, 5, 5},
	}

	for _, testCase := range testCases {
		p := NewPosition(testCase.initial, 0)
		p.ClampX(testCase.rangeMin, testCase.rangeMax)
		if p.X != testCase.expected {
			t.Fatalf("expected X value of %d clamped in range [%d, %d] to be %d, got %d", testCase.initial, testCase.rangeMin, testCase.rangeMax, testCase.expected, p.X)
		}
	}
}

func TestClampY(t *testing.T) {

	type TestCase struct {
		initial            int
		rangeMin, rangeMax int
		expected           int
	}

	testCases := []TestCase{
		{1, 5, 7, 7},
		{2, 5, 7, 5},
		{3, 5, 7, 6},
		{9, 5, 7, 6},
		{10, 5, 7, 7},
		{11, 5, 7, 5},
		{5, 5, 5, 5},
		{6, 5, 5, 5},
		{4, 5, 5, 5},
	}

	for _, testCase := range testCases {
		p := NewPosition(0, testCase.initial)
		p.ClampY(testCase.rangeMin, testCase.rangeMax)
		if p.Y != testCase.expected {
			t.Fatalf("expected Y value of %d clamped in range [%d, %d] to be %d, got %d", testCase.initial, testCase.rangeMin, testCase.rangeMax, testCase.expected, p.Y)
		}
	}
}
