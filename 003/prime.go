package main

func LargestPrime(number int64) (int64) {
	for number % 2 == 0 {
		number /= 2
	}
	for i := int64(3); i*i < number; i += 2 {
		for number % i == 0 {
			number /= i
		}
	}	
	return number;
}

func main() {
	println(LargestPrime(600851475143))
}