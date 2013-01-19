package main

import (
	"testing"
)

func TestMultiplesOf3And5To10(t *testing.T) {
	values := []int {3, 5}
	expected := 23
	out := Sum(values, 10)
	if out != expected {
		t.Errorf("Sum(%v, %v) = %v, want %v", values[0], values[1], out, expected)
	}
}

func TestMultiplesOf3And5To1000(t *testing.T) {
	values := []int {3, 5}
	expected := 233168
	out := Sum(values, 1000)
	if out != expected {
		t.Errorf("Sum(%v, %v) = %v, want %v", values[0], values[1], out, expected)
	}
}
