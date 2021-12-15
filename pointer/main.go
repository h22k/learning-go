package main

import . "fmt"

func main() {
	a := 5
	Println(&a)
	b := &a
	*b = 10
	Println(a)
	Println(b)
	Println(*b)

	if b == &a {
		Println("pointers are equal")
	} else {
		Println("pointers are not equal")
	}

	var arr []string
	arr = append(arr, "first")
	arr2 := arr
	Println(&arr)
	Println(arr2)
	Println(&arr2)
	arr2 = appendArr(arr)
	Println(arr)
	Println(arr2)
}

func appendArr(arr []string) []string {
	return append(arr, "second", "third")
}