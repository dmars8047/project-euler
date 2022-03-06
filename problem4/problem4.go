package problem4

import "fmt"

func isItAPalindrome(val int) bool {
	valString := fmt.Sprintf("%d", val)
	backwardValString := ""

	index := len(valString) - 1

	for index > -1 {
		backwardValString = fmt.Sprintf("%s%c", backwardValString, valString[index])
		index--
	}

	return valString == backwardValString
}

func Execute() int {
	result := 0

	x := 100
	y := 100

	for x < 1000 {
		for y < 1000 {
			z := x * y

			if isItAPalindrome(z) && z > result {
				result = z
			}
			y++
		}
		y = 100
		x++
	}

	return result
}
