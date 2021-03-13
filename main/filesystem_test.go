package main

import "testing"

func TestGetHostname(t *testing.T) {
	got := GetHostname()
	want := "PC"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func Test1(t *testing.T) {
	t.Helper()
	GetHDDStats("/")
}