package problem12

const ceiling = 500

func getNumDivisors(num int) int {
	numDivisors, lo, hi := 0, 1, num

	for lo < hi {
		if num%lo == 0 {
			hi = num / lo
			numDivisors += 2
		}
		lo++
	}

	return numDivisors
}

func Execute() int {
	result := 0

	i := 1
	val := 0

	for true {
		val += i
		numDivisors := getNumDivisors(val)
		if numDivisors > ceiling {
			result = val
			break
		}

		i++
	}

	return result
}
