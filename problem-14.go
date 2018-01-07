package main

import "fmt"

func main() {
	list := make(map[int]int)
	longestChain := 0
	result := 0
	for i := 1; i < 1000000; i++ {
		temp := getSequence(list, i)
		if temp > longestChain {
			result = i
			longestChain = temp
		}
	}
	fmt.Println(result, ":", longestChain, "chain")
}

func getSequence(list map[int]int, num int) int {
	if val, ok := list[num]; ok {
		return val + 1
	}

	if num == 1 {
		return 1
	}

	if num%2 == 0 {
		temp := num / 2
		list[num] = getSequence(list, temp) + 1
	} else {
		temp := 3*num + 1
		list[num] = getSequence(list, temp) + 1
	}

	return list[num]
}
