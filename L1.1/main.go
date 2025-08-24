package main

import (
	"fmt"
)

type Human struct {
	Name string
}

func (h *Human) SayHi() {
	fmt.Printf("Hi, I'm %s\n", h.Name)
}

type Action struct {
	Human
}

func main() {
	action := Action{
		Human: Human{Name: "Lol"},
	}
	action.SayHi()
}
