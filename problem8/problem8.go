package problem8

import (
	"os"
	"strconv"
)

const adjacentDigits = 13
const numDigits = 1000

func readDataFileAndParse() [numDigits]int {

	file, err := os.Open("data.txt")
	if err != nil {
		// handle the error here
		panic("Error: Could not open file")
	}
	defer file.Close()

	// read the file
	buffer := make([]byte, 1)
	eof := 1
	var parsedInts [numDigits]int
	index := 0

	for eof == 1 {
		eof, err = file.Read(buffer)

		if eof == 0 {
			break
		} else if err != nil {
			panic(err.Error())
		}

		parsedInts[index], err = strconv.Atoi(string(buffer))

		// surpress newlines
		if err == nil {
			index++
		}
	}

	return parsedInts
}

func findProduct(nums []int) int {
	sum := 1
	for i := 0; i < len(nums); i++ {
		sum *= nums[i]
	}

	return sum
}

func Execute() int {
	result := 0

	data := readDataFileAndParse()

	for i := 0; (i + adjacentDigits) < numDigits; i++ {
		product := findProduct(data[i:(i + adjacentDigits)])
		if result < product {
			result = product
		}
	}

	return result
}
