package main

func checkPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i < n; i++ {

		if n%i == 0 {
			return false
		}
	}
	return true
}

func test_prime() {
	println(checkPrime(2))
}
