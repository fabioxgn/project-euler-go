package main

func main() {
	println(SumSquareDifference(100))
}

func SumSquareDifference(limit int) (result int) {
	result = 0
	for i := 1; i <= limit; i++ {
		result += i * i
	}
	sum := 0
	for i := 1; i <= limit; i++ {
		sum += i;
	}
	result = (sum * sum) - result
	return result
}