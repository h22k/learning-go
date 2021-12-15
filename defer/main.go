package main

import . "fmt"

func main() {
	//defer f.Println("first 1")
	//defer f.Println("first 2")
	//defer f.Println("first 3")
	//defer f.Println("first 4")
	//defer f.Println("first 5")
	//f.Println("second")

	deger := 0
	for deger < 10 {
		defer Println(deger) // its bad idea btw.
		deger++
	}

	Println("After for loop")

}