package main

import "fmt"

func generic(contract interface{}) {
	fmt.Println(contract)
}

func main() {
	generic("String")
	generic(1)
	generic(true)

}
