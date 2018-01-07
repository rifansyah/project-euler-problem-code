package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	num := strconv.FormatFloat(math.Pow(2, 1000), 'f', 0, 64)
	fmt.Printf("Sum of the digits : %d\n", getSumOfDigit(num))
}

func getSumOfDigit(num string) int {
	result := 0
	for i := 0; i < len(num); i++ {
		tempNumb, _ := strconv.Atoi(string(num[i]))
		result += tempNumb
	}
	return result
}
