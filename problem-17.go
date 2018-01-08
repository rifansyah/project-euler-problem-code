package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	result := 0
	for i := 1; i <= 1000; i++ {
		temp := getWordNumb(i)
		tempArr := strings.Split(temp, " ")
		temp = strings.Join(tempArr, "")
		tempArr = strings.Split(temp, "-")
		temp = strings.Join(tempArr, "")

		result += len(temp)
	}

	fmt.Println("total :", result)
}

func getWordNumb(num int) string {
	if num == 0 {
		return "zero"
	}

	smallNumb := []string{
		"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven",
		"twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen",
	}

	dozens := []string{
		"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety",
	}

	result := ""
	for i := 1; ; i++ {
		numbInString := strconv.Itoa(num)
		if len(numbInString) == 4 {
			temp := num / 1000
			result = fmt.Sprintf("%s thousand", smallNumb[temp])

			if num%1000 <= 0 {
				break
			}

			num %= 1000
		} else if len(numbInString) == 3 {
			temp := num / 100
			if len(result) == 0 {
				result = fmt.Sprintf("%s%s hundred", result, smallNumb[temp])
			} else {
				result = fmt.Sprintf("%s %s hundred", result, smallNumb[temp])
			}

			if num%100 > 0 {
				result = fmt.Sprintf("%s and", result)
			} else {
				break
			}

			num %= 100
		} else if len(numbInString) == 2 {
			temp := num / 10
			if len(result) == 0 {
				if temp == 1 {
					result = fmt.Sprintf("%s%s", result, smallNumb[num])
					return result
				} else {
					result = fmt.Sprintf("%s%s", result, dozens[temp])
				}
			} else {
				if temp == 1 {
					result = fmt.Sprintf("%s %s", result, smallNumb[num])
					return result
				} else {
					result = fmt.Sprintf("%s %s", result, dozens[temp])
				}
			}

			if num%10 > 0 {
				result = result + "-"
			} else {
				break
			}

			num %= 10
		} else {
			if len(result) > 0 && string(result[len(result)-1]) == "-" || len(result) == 0 {
				result = fmt.Sprintf("%s%s", result, smallNumb[num])
			} else {
				result = fmt.Sprintf("%s %s", result, smallNumb[num])
			}
			break
		}
	}
	return result
}
