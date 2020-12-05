package five

import "testing"

func TestFindRow(t *testing.T) {
	want := 44
	if got := FindRow("FBFBBFFRLR"); got != want {
		t.Errorf("FindRow() = %v, want %v", got, want)
	}
}

func TestFindColumn(t *testing.T) {
	want := 5
	if got := FindColumn("FBFBBFFRLR"); got != want {
		t.Errorf("FindColumn() = %v, want %v", got, want)
	}
}

func TestFindSeatID(t *testing.T) {
	want := 357
	if got := FindSeatID("FBFBBFFRLR"); got != want {
		t.Errorf("FindSeatID() = %v, want %v", got, want)
	}
}
