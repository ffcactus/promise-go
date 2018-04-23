package main

import (
	"fmt"
)

type Show interface {
	Print(v Value)
}

type Value interface {
	Value() string
}

type MyShow struct {
}

type MyValue struct {
}

func (s *MyValue) Value() string {
	return "MyValue"
}

func (s *MyShow) Print(v *MyValue) {
	fmt.Println(v.Value())
}

func main() {
	var (
		value Value
		show  Show
	)
	value = new(MyValue)
	show = new(MyShow)
	show.Print(value)
}
