package main

import "fmt"

func main() {
	n := -1

	if n > 15 {
		fmt.Println("> 15")
	} else {
		fmt.Println("<= 15")
	}

	if n1 := n; n1 > 0 {
		fmt.Println("> Zero")
	}

}
