package main

import "strconv"

func main() {
	println(Palindrome(3))
}

func Palindrome(digits int) (int) {
	result := 0
	for first := getMaxNumber(digits); first > 0; first-- {
		for second := first; second > 0; second -- {
			str := strconv.Itoa(first * second)		
			if IsPalindrome(str) {
				number, _ := strconv.Atoi(str)
				if number > result {
					result = number
				}
			}
		}
	}		
	return result
}

func getMaxNumber(digits int) (int) {
	str := ""
	for i := 0; i < digits; i++ {
		str += strconv.Itoa(9)
	}

	result, _ := strconv.Atoi(str)
	return result
}

func IsPalindrome(str string) (bool) {	
	i := 0
	j := len(str) - 1
	word := []rune(str)	
	for j > i {
		if word[i] != word[j] {
			return false
		}
		i++
		j--
	}	
	return true
}