package problem9

import (
	"math"
)

const targetSum = 1000

func Execute() int {
	a, b, c, result := 0, 0, 0, 0

	for result == 0 {
		a++
		b = a + 1
		c = targetSum - b - a
		for b < c {
			if a+b+c == targetSum && math.Pow(float64(a), 2)+math.Pow(float64(b), 2) == math.Pow(float64(c), 2) {
				result = a * b * c
				break
			}
			c--
			b++
		}
	}

	return result
}
