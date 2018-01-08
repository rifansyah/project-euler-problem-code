package main

import "fmt"

func main() {
	listDayInMonth := map[int]int{
		1: 31, 2: 28, 3: 31, 4: 30, 5: 31, 6: 30, 7: 31, 8: 31, 9: 30, 10: 31, 11: 30, 12: 31,
	}

	day := 7
	month := 1
	year := 1900

	count := 0

	for {
		if day >= 1 && month >= 1 && year >= 2001 {
			break
		}

		if isLeapYear(year) {
			listDayInMonth[2] = 29
		} else {
			listDayInMonth[2] = 28
		}

		day += 7

		if day > listDayInMonth[month] {
			day -= listDayInMonth[month]
			month++
		}

		if month > 12 {
			year++
			month -= 12
		}

		if year >= 1901 && year <= 2000 {
			if day == 1 {
				count++
				fmt.Printf("day : %d, month : %d, year : %d\n", day, month, year)
			}
		}

	}

	fmt.Println("total : ", count)
}

func isLeapYear(year int) bool {
	if year%100 == 0 && year%400 == 0 {
		return true
	} else if year%100 == 0 {
		return false
	} else if year%4 == 0 {
		return true
	}
	return false
}

// question
/*


You are given the following information, but you may prefer to do some research for yourself.

    1 Jan 1900 was a Monday.
    Thirty days has September,
    April, June and November.
    All the rest have thirty-one,
    Saving February alone,
    Which has twenty-eight, rain or shine.
    And on leap years, twenty-nine.
    A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.

How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?

*/
