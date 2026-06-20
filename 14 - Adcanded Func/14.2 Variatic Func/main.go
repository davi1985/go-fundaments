package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}

	return total
}

func write(txt string, numbers ...int) {
	for _, value := range numbers {
		fmt.Println(string(txt[value]))
	}
}

func main() {
	total := sum(1, 2, 4, 5, 67, 0, 3)

	fmt.Println(total)

	write("Hello World", 1, 2, 3, 4, 5, 9)
}
