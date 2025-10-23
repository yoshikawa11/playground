package iteration

import (
	"testing"
)

func TestContains(t *testing.T) {
	expected := true
	got := Contains("seafood", "afo")
	if expected != got {
		t.Errorf("got %t but expected %t", got, expected)
	}
}
