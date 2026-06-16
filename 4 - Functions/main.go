package main

import "fmt"

func sum(n1, n2 int8) int8 {
	return n1 + n2
}

func calcs(n1, n2 int8) (int8, int8) {
	sum := n1 + n2
	sub := n1 - n2

	return sum, sub
}

func main() {
	result := sum(2, 5)
	println(result)

	var f = func(txt string) {
		fmt.Println("Hello, " + txt)
	}

	f("Davi Silva ")

	sum, sub := calcs(5, 4)

	fmt.Println(sum)
	fmt.Println(sub)

	sum2, _ := calcs(7, 4)
	fmt.Println(sum2)
}
