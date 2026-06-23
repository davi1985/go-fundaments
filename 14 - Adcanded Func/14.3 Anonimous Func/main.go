package main

import "fmt"

func main() {
	result := func(txt string) string {
		return fmt.Sprintf("Get -> %s", txt)
	}("Params")

	fmt.Println(result)
}
