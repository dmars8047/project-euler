package problem3

const target = 600851475143

func isItPrime(val int) bool {

	for i := 2; i < val; i++ {
		if val%i == 0 {
			return false
		}
	}

	return true
}

func Execute() int {

	maxPrimeFactor := 1
	ceiling := target

	for i := 1; i < ceiling; i++ {
		if target%i == 0 {

			ceiling = target / i

			if isItPrime(i) {
				maxPrimeFactor = i
			}
		}
	}

	return maxPrimeFactor
}
