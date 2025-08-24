package machine

import (
	"strings"
	"testing"
)

func TestParsingFromValidLines(t *testing.T) {

	lines := []string{
		"Button A: X+94, Y+34",
		"Button B: X+22, Y+67",
		"Prize: X=8400, Y=5400",
	}

	_, err := TryParseFromLines(lines)
	if err != nil {
		t.Fatalf("got an error while parsing a valid string: %s", err.Error())
	}

	// too lazy to write the rest
}

func TestParsingFromInvalidLines(t *testing.T) {

	lines := []string{
		"Button A: X+94 Y+34",
		"Button B: X+22 Y+67",
		"Prize: X=8400 Y=5400",
	}

	machine, err := TryParseFromLines(lines)
	if err == nil {
		t.Fatalf("parsing invalid lines succeeded:\nlines:\n%s\nmachine: %s", strings.Join(lines, "\n"), machine.String())
	}
}
