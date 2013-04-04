package main

import (
	"testing"
)

func TestPrimes(t *testing.T) {
	want := 13
	got := FindPrime(6)
	if got != want {
		t.Errorf("want %v got %v", want, got)
	}
}
