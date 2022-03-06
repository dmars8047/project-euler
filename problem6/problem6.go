package problem6

import (
	"math"
)

const ceiling = 100

func sumOfSquares(sumOfSquares chan int) {
	sum := 0.0
	for i := 1.0; i <= ceiling; i++ {
		sum += math.Pow(i, 2)
	}

	sumOfSquares <- int(sum)
}

func squareOfSums(squareOfSum chan int) {
	sum := 0.0
	for i := 1.0; i <= ceiling; i++ {
		sum += i
	}

	squareOfSum <- int(math.Pow(sum, 2))
}

func Execute() int {
	result := 0

	sumOfSquaresChan := make(chan int)
	squareOfSumsChan := make(chan int)
	squareOfSumResult := 0
	sumOfSquaresResult := 0

	go sumOfSquares(sumOfSquaresChan)
	go squareOfSums(squareOfSumsChan)

	for squareOfSumResult == 0 || sumOfSquaresResult == 0 {
		select {
		case sumSqs := <-sumOfSquaresChan:
			sumOfSquaresResult = sumSqs
		case sqSum := <-squareOfSumsChan:
			squareOfSumResult = sqSum
		}
	}

	result = squareOfSumResult - sumOfSquaresResult

	return result
}
