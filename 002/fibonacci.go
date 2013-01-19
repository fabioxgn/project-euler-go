package main

func SumEven(limit int) (int) {
	fib := []int {1, 2, 3}
	result := 2
	for fib[2] < limit {		
		fib[0] = fib[1]
		fib[1] = fib[2]
		fib[2] = fib[0] + fib[1]				
		if fib[2] % 2 == 0 {
			result += fib[2]
		}		
	}
	return result
}

func main() {
	println(SumEven(4000000));
}
