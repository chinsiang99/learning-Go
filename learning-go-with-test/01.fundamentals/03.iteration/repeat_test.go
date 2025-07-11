package iteration

import "testing"

func TestRepeat(t *testing.T) {
	result := Repeat("a")
	expected := "aaaaa"

	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}
