package main

import (
	"fmt"
)

func deleteFromSlice(arr []int, idx int) []int {
	copy(arr[idx:], arr[idx+1:])
	return arr[:len(arr)-1]
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	a = deleteFromSlice(a, 4)
	fmt.Println(a)
}

