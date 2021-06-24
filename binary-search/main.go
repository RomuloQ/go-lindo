package main

import (
	"fmt"
)

func main() {
	numbers := makeRange(1,100000000)

	var exist = search(numbers, 2)

	fmt.Println(exist)
}

func search(numbers []int, paramSearch int) bool {

	var half = len(numbers) / 2

	if paramSearch > numbers[half] && len(numbers) != 1 {
		return search(numbers[half:], paramSearch)
	} else if paramSearch < numbers[half] && len(numbers) != 1 {
		return search(numbers[:half],paramSearch)
	} else if paramSearch == numbers[half] {
		return true
	}

	return false
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}


