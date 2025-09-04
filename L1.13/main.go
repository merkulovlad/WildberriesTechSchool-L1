package main

import (
	"fmt"
)

func swap(a, b *int) {
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
}

func main() {
	a := 5
	b := 10
	swap(&a, &b)
	fmt.Println(a, b)
}
