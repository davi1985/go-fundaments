package main

import "fmt"

type person struct {
	name    string
	surname string
	age     uint8
	height  uint8
}

type student struct {
	person
	course  string
	college string
}

func main() {
	p1 := person{"Davi", "Silva", 40, 160}

	fmt.Println(p1)

	e1 := student{p1, "Engenharia", "USP"}

	fmt.Println(e1.name)
}
