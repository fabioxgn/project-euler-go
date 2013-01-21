package main

import "testing"

func TestMultiple(t *testing.T) {
	out := FindMultiple(10)
	want := 2520
	if out != want {
		t.Errorf("Got %v want %v", out, want)
	}
}
