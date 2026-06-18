package main

import (
	"fmt"
	"reflect"
)

func main() {
	var arr1 [5]string
	arr1[0] = "Pos 1"
	fmt.Println(arr1)

	arr2 := [5]string{"Pos 1", "Pos 2", "Pos 3", "Pos 4", "Pos 5"}
	fmt.Println(arr2)

	arr3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr3)

	slice := []int{2, 3, 4, 5}

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(arr3))

	slice = append(slice, 20)
	fmt.Println(slice)

	slice2 := arr2[1:3]
	fmt.Println(slice2)
	arr2[1] = "Pos Alt"
	fmt.Println(slice2)

	fmt.Println("---------------------")

	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
	fmt.Println(slice3)

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
	fmt.Println(slice4)
}
