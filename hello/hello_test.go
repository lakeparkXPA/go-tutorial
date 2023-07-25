package main

import (
	"rsc.io/quote"
	"testing"
)

func TestHello(t *testing.T) {
	got := Quote()
	want := quote.Go()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
