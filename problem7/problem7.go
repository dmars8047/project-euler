package problem7

var numPrimesSearchingFor = 10001

func isItPrime(num int) bool {
	for i := 1; i <= num; i++ {
		if i != 1 && i != num {
			if num%i == 0 {
				return false
			}
		}
	}

	return true
}

func Execute() int {
	result := 0
	numPrimesFound := 0

	for i := 2; numPrimesFound < numPrimesSearchingFor; i++ {
		if isItPrime(i) == true {
			numPrimesFound++
			if numPrimesFound == numPrimesSearchingFor {
				result = i
			}
		}
	}

	return result
}
