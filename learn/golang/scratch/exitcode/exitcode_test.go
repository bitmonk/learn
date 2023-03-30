package exitcode

import "testing"

func TestExitCode(t *testing.T) {
	assertCorrectReturn := func(t testing.TB, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("Explicitly request a true value from sh", func(t *testing.T) {
		got := GetExitCode("sh", "-c", "true")
		want := 0
		assertCorrectReturn(t, got, want)
	})

	t.Run("Explicitly request a false value from sh", func(t *testing.T) {
		got := GetExitCode("sh", "-c", "false")
		want := 1
		assertCorrectReturn(t, got, want)
	})
}
