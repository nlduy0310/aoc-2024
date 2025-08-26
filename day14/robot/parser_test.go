package robot

import "testing"

func TestParsingValidString(t *testing.T) {

	inp := "p=0,4 v=3,-3"

	_, err := TryParseFromString(inp)
	if err != nil {
		t.Fatalf("can not parse robot from a valid string: %s", err.Error())
	}
}

func TestParsingInvalidString(t *testing.T) {

	inp := "p=0,4, v=3,-3"

	robot, err := TryParseFromString(inp)
	if err == nil {
		t.Fatalf("parsing robot from an invalid string succeeded:\nstring: '%s'\nrobot: %s", inp, robot.String())
	}
}
