package main

import "testing"

func Test2DigitPalindromeMultiplication(t *testing.T) {
	expectedNumber := 9009
	number := Palindrome(2)

	if number != expectedNumber {
		t.Errorf("Expected (%v) but was (%v)", expectedNumber, number)
	}
}

func TestIsPalindrome(t *testing.T) {
	value := "90909"
	if !IsPalindrome(value) {
		t.Errorf("Value (%v) should be palindrome", value)	
	}

	value = "909009"
	if IsPalindrome(value) {
		t.Errorf("Value (%v) should not be palindrome", value)	
	}	
}