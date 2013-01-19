package main

import "testing"

func TestLargestPrime(t *testing.T) {
	want := int64(29)
	number := int64(13195)
	if out := LargestPrime(number); out != want {
		t.Errorf("Want (%v) was (%v)", want, out)
	}
}