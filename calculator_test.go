package calculator

import "testing"

func TestCalculator (t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("Test Input Empty String Return 0", func(t *testing.T) {
		got := Calculator("")
		want := "0"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Test Input 1 and Return 1", func(t *testing.T) {
		got := Calculator("1")
		want := "1"
		assertCorrectMessage(t, got, want)
	})
}