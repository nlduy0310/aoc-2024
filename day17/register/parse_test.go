package register

import "testing"

func TestParseValidString(t *testing.T) {
	testCases := []struct {
		line string
		name string
		val  int
	}{
		{"Register A: 0", "A", 0},
		{"Register   B:   012", "B", 12},
		{"Register\tC:\t120", "C", 120},
	}

	for _, testCase := range testCases {
		register, err := ParseLine(testCase.line)
		if err != nil {
			t.Fatalf(
				"parsing register line \"%s\" failed with error: %s",
				testCase.line, err.Error(),
			)
		}
		if register.Name != testCase.name {
			t.Fatalf(
				"parsing register line \"%s\" failed: expected name to be \"%s\", got \"%s\"",
				testCase.line, testCase.name, register.Name,
			)
		}
		if register.LiteralValue != testCase.val {
			t.Fatalf(
				"parsing register line \"%s\" failed: expected value to be %d, got %d",
				testCase.line, testCase.val, register.LiteralValue,
			)
		}
	}
}
