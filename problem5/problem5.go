package problem5

const floor = 2
const ceiling = 20

func Execute() int {
	result := 0

	i := 1

	for true {
		if IsSmallestMultiple(i, floor, ceiling) {
			result = i
			break
		}
		i++
	}

	return result
}

func IsSmallestMultiple(val int, divisor int, ceiling int) bool {
	if val%divisor == 0 {
		if divisor == ceiling {
			return true
		}

		return IsSmallestMultiple(val, divisor+1, ceiling)
	}

	return false
}
