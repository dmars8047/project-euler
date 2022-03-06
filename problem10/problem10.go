package problem10

var ceiling = 2000000

func isItPrime(val int) bool {

	upperBound := val

	for i := 2; i < upperBound; i++ {
		q, r := divideAndGetRemainder(val, i)

		if r == 0 {
			return false
		}
		upperBound = q
	}

	return true
}

// This function returns remainder of num/divisor
// without using % (modulo) operator
func divideAndGetRemainder(dividend, divisor int) (quotient, remainder int) {
	q, r := 0, 0

	q = dividend / divisor

	r = (dividend - divisor*q)

	return q, r
}

func Execute() int {
	result := 2

	for i := 3; i < ceiling; {
		if isItPrime(i) == true {
			result += i
		}

		i += 2
	}

	return result
}
