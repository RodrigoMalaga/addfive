package addfive

import "testing"

func TestAddFive(t *testing.T) {
	result := AddFive(10)
	expected := 15

	if result != expected {
		t.Errorf("expected %d but got %d", expected, result)
	}
}
