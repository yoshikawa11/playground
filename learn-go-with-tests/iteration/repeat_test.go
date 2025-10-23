package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("repeat 5 times.", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("repeat 0 time.", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := ""
		assertCorrectMessage(t, repeated, expected)
	})
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}
