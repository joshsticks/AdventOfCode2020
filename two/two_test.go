package two

import (
	"testing"
)

func TestGetRegexVariables(t *testing.T) {
	expectedMin := "1"
	expectedMax := "3"
	expectedChar := "a"

	min, max, letter, _ := GetRegexVariables("1-3 a: abcde")

	if min != expectedMin {
		t.Errorf("GetRegexVariables() min = %q, want %q", min, expectedMin)
	}

	if max != expectedMax {
		t.Errorf("GetRegexVariables() max = %q, want %q", max, expectedMax)
	}

	if letter != expectedChar {
		t.Errorf("GetRegexVariables() letter = %q, want %q", letter, expectedChar)
	}
}

func TestGetGenerateRegex(t *testing.T) {
	want := `:\s([^a]*[a]){1,3}[^a]*$`

	if got := GenerateRegex("1-3 a: abcde"); got != want {
		t.Errorf("GenerateRegex() = %q, want %q", got, want)
	}
}
