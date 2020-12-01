package one

import "testing"

func TestFindTwentyTwenty(t *testing.T) {
	want := 514579
	if got := FindTwentyTwenty(1721, 299); got != want {
		t.Errorf("FindTwentyTwenty() = %q, want %q", got, want)
	}
}
