package main

import (
	"fmt"
	"reflect"
)

func main() {
	var array1 [5]int
	fmt.Println(array1)

	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array1))

	slice = append(slice, 18)
	fmt.Println(slice)

	// slice is a pointer to an array

	// internal array

	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)

	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
}
