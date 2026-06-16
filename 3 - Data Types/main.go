package main

import (
	"errors"
	"fmt"
)

func main() {
	var i int8 = 10
	var n int16 = 1000

	fmt.Printf("n: %d\n", n)
	fmt.Printf("i: %d\n", i)

	var u uint = 100
	fmt.Printf("uint: %d\n", u)

	var n2 int32 = 200
	var n3 rune = 300 // rune is an alias for int32
	var c byte = 2    // byte is an alias for uint8
	fmt.Printf("n2: %d\n", n2)
	fmt.Printf("n3: %d\n", n3)
	fmt.Printf("c: %d\n", c)

	var f float32 = 3.14
	var d float64 = 3.141592653589793
	fmt.Printf("f: %f\n", f)
	fmt.Printf("d: %f\n", d)

	n4 := 4.5 // type inference, n4 is float64
	fmt.Printf("n4: %f\n", n4)

	var s string = "Hello, Go!"
	fmt.Printf("s: %s\n", s)
	char := 'A' // char is a rune (int32) representing the Unicode code point for 'A'
	fmt.Println("char:", char)

	// Zero values
	var zeroInt int
	var zeroFloat float64
	var zeroString string
	var zeroBool bool

	fmt.Printf("zeroInt: %d\n", zeroInt)
	fmt.Printf("zeroFloat: %f\n", zeroFloat)
	fmt.Printf("zeroString: '%s'\n", zeroString)
	fmt.Printf("zeroBool: %t\n", zeroBool)

	isActive := true // type inference, isActive is bool
	var isAvailable bool = false
	fmt.Printf("isActive: %t\n", isActive)
	fmt.Printf("isAvailable: %t\n", isAvailable)

	var err error = nil // error is an interface type, nil means no error
	fmt.Printf("err: %v\n", err)

	myError := errors.New("Internal error")
	fmt.Printf("err: %v\n", myError)
}
