package main

import "fmt"

func mathCalcs(n1, n2 int) (sum, subtraction int) {
	sum = n1 + n2
	subtraction = n1 - n2
	return
}

func main() {
	sum, sub := mathCalcs(2, 4)

	fmt.Println(sum, sub)
}
