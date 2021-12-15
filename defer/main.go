package main

import f "fmt"

func main() {
	defer f.Println("first")
	f.Println("second")
}