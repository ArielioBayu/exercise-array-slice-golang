package main

import "fmt"

func SchedulableDays(date1 []int, date2 []int) []int {

	result := make([]int, 0)

	for i := 1; i <= 31; i++ {

		isImamAvailable := false
		isPermanaAvailable := false

		for _, date := range date1 {
			if date == i {
				isImamAvailable = true
			}
		}
		for _, date := range date2 {
			if date == i {
				isPermanaAvailable = true
			}
		}

		if isImamAvailable && isPermanaAvailable {
			result = append(result, i)
		}
	}
	return result // TODO: replace this
}

func main() {
	fmt.Println(SchedulableDays([]int{1, 2, 3, 4}, []int{3, 4, 5}))
	fmt.Println(SchedulableDays([]int{11, 12, 13, 14, 15}, []int{5, 10, 12, 13, 20, 21}))
	fmt.Println(SchedulableDays([]int{2, 7, 12, 20, 21, 22}, []int{1, 3, 6, 10}))
	fmt.Println(SchedulableDays([]int{}, []int{}))
}
