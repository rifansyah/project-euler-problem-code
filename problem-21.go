package main

import "fmt"

func main() {
	list := make(map[int]int)
	listResult := make(map[int]bool)
	result := 0

	fmt.Print("Amicable number : ")

	for i := 1; i < 10000; i++ {
		for j := 1; j < 10000; j++ {
			if i == j {
				continue
			}

			if _, ok := listResult[j]; ok {
				continue
			}

			if getDivisor(list, i) == j && getDivisor(list, j) == i {
				result += i + j
				fmt.Printf("[%d, %d], ", i, j)
				listResult[i] = true
			}
		}
	}
	fmt.Println("\nSum of amicable numbers :", result)
}

func getDivisor(list map[int]int, num int) int {
	total := 1
	temp := 0

	if val, ok := list[num]; ok {
		return val
	}

	for i := 2; i < num; i++ {
		if num%i == 0 && i*i == num {
			total += i
			break
		}

		if num%i == 0 && num/i == temp {
			break
		}

		if num%i == 0 {
			secondDivisor := num / i
			total += i + secondDivisor
			temp = i
		}
	}
	list[num] = total
	return total
}
