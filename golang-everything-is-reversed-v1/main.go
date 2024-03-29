package main

import (
	"fmt"
	"strconv"
)

func reversePosition(arr [5]int) [5]int {
	result := [5]int{}

	for i := 0; i < len(arr); i++ {
		result[i] = arr[len(arr)-1-i]
	}
	return result
}

func reverseDigit(number int) int {
	str := strconv.Itoa(number)

	cap := ""
	for _, value := range str {
		cap = string(value) + cap
	}
	result, _ := strconv.Atoi(cap)
	return result

}

func ReverseData(arr [5]int) [5]int {

	x := reversePosition(arr)
	// fmt.Println("Hasil reverse position", x)

	// [123, 456, 789] => awalnya
	// [789, 456, 123] => reverse posisi
	// [987, 654, 321] => reverse digit

	for i := 0; i < len(x); i++ {
		x[i] = reverseDigit(x[i])
	}
	// fmt.Println("Hasil reverse digit a", x)

	return x // TODO: replace this
}

func main() {
	fmt.Println(ReverseData([5]int{123, 456, 11, 1, 2}))
	fmt.Println(ReverseData([5]int{456789, 44332, 2221, 12, 10}))
	fmt.Println(ReverseData([5]int{10, 10, 10, 10, 10}))
	fmt.Println(ReverseData([5]int{23456, 789, 123, 456, 500}))
	fmt.Println(ReverseData([5]int{0, 0, 0, 0, 0}))
}
