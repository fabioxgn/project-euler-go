package main

import "testing"

func TestFibonacci(t *testing.T) {
	limit := 10
	want := 10
	if out := SumEven(limit); out != want {
		t.Errorf("Want (%v) was (%v)", want, out)
	}
}