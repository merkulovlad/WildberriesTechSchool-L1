package main

import (
	"fmt"
	"reflect"
)

func checkerType(value interface{}) {
    v := reflect.ValueOf(value)
    if v.Kind() == reflect.Chan {
        fmt.Println("chan")
        return
    }
    switch value.(type) {
    case int:
        fmt.Println("int")
    case string:
        fmt.Println("string")
    case bool:
        fmt.Println("bool")
    default:
        fmt.Println("unknown")
    }
}

func main() {
	checkerType(1)
	checkerType("hello")
	checkerType(true)
	checkerType(make(chan int))
	checkerType(3.14)
	checkerType([]int{1, 2, 3})
}
