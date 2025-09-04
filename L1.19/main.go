package main 
import (
	"fmt"
)
func reverseString(strings []rune) []rune {
	n := len(strings)
	for i := 0; i < n/2; i++ {
		strings[i], strings[n-i-1] = strings[n-i-1], strings[i]
	}
	return strings
}

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Println(string(reverseString([]rune(str))))
}
