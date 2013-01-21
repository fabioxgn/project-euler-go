package main

func main() {
	println(FindMultiple(20))
}

func FindMultiple(limit int) int {
	result := 2
	for ; ; result++ {
		remainder := false
		for i := 1; i <= limit; i++ {
			remainder = (result % i) != 0
			if remainder {
				break
			}
		}
		if !remainder {
			return result
		}
	}
	return 0
}
