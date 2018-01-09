package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	result := 0

	file, err := ioutil.ReadFile("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fileInString := string(file)
	listName := strings.Split(fileInString, ",")

	for i := 0; i < len(listName); i++ {
		listName[i] = strings.Replace(listName[i], "\"", "", -1)
	}

	sort.Strings(listName)

	for i := 0; i < len(listName); i++ {
		temp := getTotalValueOfString(listName[i])
		temp *= i + 1
		result += temp
	}

	fmt.Println("Total :", result)
}

func getTotalValueOfString(name string) int {
	result := 0
	for i := 0; i < len(name); i++ {
		result += getValueOfEachChar(string(name[i]))
	}
	return result
}

func getValueOfEachChar(char string) int {
	return int(char[0]) - 64
}

// question :
// Using names.txt (right click and 'Save Link/Target As...'), a 46K text file containing over five-thousand first names, begin by sorting it into alphabetical order. Then working out the alphabetical value for each name, multiply this value by its alphabetical position in the list to obtain a name score.
// For example, when the list is sorted into alphabetical order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list. So, COLIN would obtain a score of 938 × 53 = 49714.
// What is the total of all the name scores in the file?
