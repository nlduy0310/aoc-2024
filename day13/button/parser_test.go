package button

import "testing"

func TestButtonParseFromString(t *testing.T) {

	t.Run("test parsing a valid string", func(t *testing.T) {

		input := "Button A: X+94, Y+34"
		button, err := TryParseFromString(input, 1)
		if err != nil {
			t.Fatalf("parsing a valid string failed for: %s", err.Error())
		}

		if button.Name != "A" {
			t.Fatalf("expected button to have name 'A', got: '%s'", button.Name)
		}

		if button.XShift != 94 {
			t.Fatalf("expected button to have xShift = 94, got: %d", button.XShift)
		}

		if button.YShift != 34 {
			t.Fatalf("expected button to have yShift = 34, got: %d", button.YShift)
		}

		if button.Cost != 1 {
			t.Fatalf("expected button to have cost = 1, got: %d", button.Cost)
		}
	})

	t.Run("test parsing an invalid string", func(t *testing.T) {

		input := "Button B: X+22 Y+67"

		button, err := TryParseFromString(input, 1)
		if err == nil {
			t.Fatalf("expected invalid string '%s' to not be parsed, got: %s", input, button.String())
		}
	})
}
