package calculator

import "testing"

func TestCalculator (t *testing.T) {
	t.Run("Test Input Empty String Return 0", func(t *testing.T) {
		got := Calculator("")
		want := "0"

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})

	t.Run("Test Input 1 and Return 1", func(t *testing.T) {
		got := Calculator("1")
		want := "1"
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}