package problem2

func isEven(val int) bool {
	return val%2 == 0
}

func Execute() int {
	sum := 0
	x := 1
	y := 2

	for y <= 4000000 {

		if isEven(y) {
			sum += y
		}

		z := x
		x = y
		y += z
	}

	return sum
}
