package main

import (
	"fmt"
)

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)
	sum()

	num := []int{1, 5, 9}
	sum(num...)
}

