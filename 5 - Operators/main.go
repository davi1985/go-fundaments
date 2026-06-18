package main

import "fmt"

func main() {
	sum := 1 + 2
	sub := 1 - 2
	division := 3 / 2
	mult := 1 * 2
	mod := 5 % 2

	fmt.Println(sum)
	fmt.Println(sub)
	fmt.Println(division)
	fmt.Println(mult)
	fmt.Println(mod)

	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)

	isTrue, isFalse := true, false
	fmt.Println(isTrue && isFalse)
	fmt.Println(isTrue || isFalse)
	fmt.Println(!isTrue)
	fmt.Println(!isFalse)

	n := 10
	n++
	fmt.Println(n)
	n--
	fmt.Println(n)
	n += 5
	fmt.Println(n)
	n *= 5
	fmt.Println(n)

	var txt string
	if n > 5 {
		txt = "N = 5"
	} else {
		txt = "N diff 5"
	}

	fmt.Println(txt)
}
