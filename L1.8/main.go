package main

import (
	"fmt"
	
)

func main() {
	var x int64
	fmt.Scan(&x)
	
	var bit int
	fmt.Scan(&bit)
	
	var needBit int
	fmt.Scan(&needBit)
	
	if needBit == 1 {
		x = x | (1 << bit)
	} else {
		x = x &^ (1 << bit)
	}
	fmt.Println(x)
}