package problem1

func isMultipleOfThree(val int) bool {
	return val%3 == 0
}

func isMultipleOfFive(val int) bool {
	return val%5 == 0
}

func Execute() int {
	sum := 0

	for i := 1; i < 1000; i++ {
		if isMultipleOfThree(i) || isMultipleOfFive(i) {
			sum += i
		}
	}

	return sum
}
