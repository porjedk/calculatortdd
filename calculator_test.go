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
}