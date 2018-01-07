package main

import "fmt"

func main() {
	triangleMap := make(map[int]int)

	for i := 1; ; i++ {
		triangle := getTriangle(i, triangleMap)
		divisor := getDivisor(triangle)
		if divisor > 500 {
			fmt.Println(triangle)
			return
		}
		// fmt.Println(triangle, divisor)
	}
}

func getTriangle(num int, mapnya map[int]int) int {

	if num == 1 {
		mapnya[1] = 1
		return 1
	}

	mapnya[num] = mapnya[num-1] + num
	return mapnya[num]
}

func getDivisor(num int) int {
	if num == 1 {
		return 1
	}
	temp := 0
	count := 0
	for i := 1; i <= num; i++ {

		if num%i == 0 && i*i == num {
			temp = i
			count++
			break
		} else if num%i == 0 {
			if num/i == temp {
				break
			}
			temp = i
			count += 2
		}
	}
	return count
}
