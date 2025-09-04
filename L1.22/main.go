package main

import (
	"fmt"
	"math/big"
)

func main() {
	var aStr, bStr string
	fmt.Scan(&aStr, &bStr)

	a := new(big.Int)
	b := new(big.Int)
	a.SetString(aStr, 10)
	b.SetString(bStr, 10)

	sum := new(big.Int).Add(a, b)
	diff := new(big.Int).Sub(a, b)
	prod := new(big.Int).Mul(a, b)
	quot := new(big.Rat).SetFrac(a, b)

	fmt.Println("Сумма:", sum)
	fmt.Println("Разность:", diff)
	fmt.Println("Произведение:", prod)
	fmt.Println("Деление:", quot)
}
