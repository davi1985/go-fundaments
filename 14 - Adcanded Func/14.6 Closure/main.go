package main

import "fmt"

func closure() func() {
	text := "Inside closure"

	fn := func() {
		fmt.Println(text)
	}
	return fn
}

func main() {
	txt := "Fn main"
	fmt.Println(txt)
	fn := closure()
	fn()
}
