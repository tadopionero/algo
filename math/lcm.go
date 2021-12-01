package main

func lcm(temp1, temp2 int) int {
	var lcmnum int = 1
	if temp1 > temp2 {
		lcmnum = temp1
	} else {
		lcmnum = temp2
	}
	for {
		if lcmnum%temp1 == 0 && lcmnum%temp2 == 0 {
			break
		}
		lcmnum++
	}
	return lcmnum
}

func test_lcm() {
	println(gcd(5, 10))
}
