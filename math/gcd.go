package main

func gcd(a, b int) int {
	var gcdnum int
	for i := 1; i <= a && i <= b; i++ {
		if a%i == 0 && b%i == 0 {
			gcdnum = i
		}
	}
	return gcdnum
}

func test_gcd() {
	println(gcd(5, 10))
}
