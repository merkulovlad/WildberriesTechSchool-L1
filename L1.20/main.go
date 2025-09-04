package main

import (
	"fmt"
)

func reverseWordsInPlace(b []byte) {
	reverse(b, 0, len(b)-1)
	start := 0
	for i := 0; i <= len(b); i++ {
		if i == len(b) || b[i] == ' ' {
			reverse(b, start, i-1)
			start = i + 1
		}
	}
}

func reverse(b []byte, left, right int) {
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}

func main() {
	s := []byte("sun dog snow")
	reverseWordsInPlace(s)
	fmt.Println(string(s))
}

