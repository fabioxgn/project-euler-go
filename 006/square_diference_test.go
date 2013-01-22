package main

import "testing"

func TestSquareDiference(t *testing.T) {
	want := 2640
	if out := SumSquareDifference(10); want != out {
		t.Errorf("Want (%v) but was (%v)", want, out)
	}
}