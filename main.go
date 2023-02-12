package main

import (
	"algorithm/customSort"
	"fmt"
)

func main() {
	slice1 := []int{4, 2, 7, 4, 5}

	customSort.QuickSort(slice1)
	fmt.Println(slice1)
}
