package filesystem

import (
	"fmt"
	"testing"
)

func TestGetHostname(t *testing.T) {
	got := GetHostname()
	want := "PC"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHDDStats(t *testing.T) {
	t.Helper()
	GetHDDStats("/")
}

func TestCommandExec(t *testing.T) {
	if CanRunCommands() {
		t.Helper()

		got, err := ExecuteCommand("whoami")
		want := "pi"
		fmt.Printf("%s", got)
		if err != nil {
			if got != want {
				t.Errorf("got %q want %q", got, want)
			}
		} else {
			t.Failed()
		}
	}
}
