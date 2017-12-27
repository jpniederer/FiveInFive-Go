package printHi

import (
	"testing"
)

func Hi() string {
	return "Hi"
	// Output: Hi
}

func TestHi(t *testing.T) {
	output := Hi()
	if output != "Hi" {
		t.Errorf("Expected Hi but got %s", output)
	}
	// Output: Hi
}
