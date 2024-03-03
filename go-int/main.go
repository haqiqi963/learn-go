package main

import "fmt"

func main() {
	//uint == Unsigned tdk boleh minus
	var a uint = 1000000000000000

	//int == signed boleh
	var b int = -40000000

	fmt.Printf("Tipe data: %T, value: %v\n", a, a)
	fmt.Printf("Tipe data: %T, value: %v\n", b, b)
}
