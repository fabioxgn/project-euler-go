package main

func Sum(values []int, limit int) (int) {
	result := 0;
	for i := 1; i < limit; i++ {		
		if isMultiple(i, values) {
			result += i
		}
	}
	return result
}

func isMultiple(number int, values[]int) (bool) {
	for _, value := range values {
		if number % value == 0 {
			return true
		}
	}	
	return false;
}

func main() {
	println(Sum([]int {3, 5}, 1000))
}