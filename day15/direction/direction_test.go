package direction

import "testing"

func TestParsingValidRunes(t *testing.T) {

	testCases := []struct {
		input    rune
		expected Direction
	}{
		{'^', Up},
		{'v', Down},
		{'<', Left},
		{'>', Right},
	}

	for _, testCase := range testCases {
		res, err := TryParseFromRune(testCase.input)
		if err != nil {
			t.Fatalf("can not parse direction from valid rune '%c': %s", testCase.input, err.Error())
		}

		if res != testCase.expected {
			t.Fatalf("wrong direction parsed from rune '%c': expected %s, got %s", testCase.input, testCase.expected.String(), res.String())
		}
	}
}
