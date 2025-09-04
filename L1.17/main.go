package main

import (
	"fmt"
)

func binarySearch(arr []int, searched int) int {
	l, r := 0, len(arr) - 1
	for l <= r {
		i := l + (r - l)/ 2
		if arr[i] > searched {
			r = i - 1
		} else if arr[i] < searched {
			l = i + 1
		} else if arr[i] == searched {
			return i
		}
	}
	return -1
}

func main(){
	var arr = []int{1,2,3,4} 
	fmt.Println(binarySearch(arr, 4))
}
