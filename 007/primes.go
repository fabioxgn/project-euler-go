package main

func FindPrime(count int) (prime int) {
	sequence := 1
	isPrime := true
	for i := 1; sequence <= count; i++ {
		if i == 2 {
			continue
		}

		isPrime = true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			prime = i
			sequence++
		}
	}
	return prime
}

func main() {
	println(FindPrime(10001))
}
