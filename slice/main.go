package main

import "fmt"

func main() {

	//Slice langsung
	var a1 = []int{1, 2, 3, 4, 5}
	a2 := []string{}

	fmt.Println(a1)
	fmt.Println(a2)

	//slice dari array
	array := [6]int{1, 2, 3, 4, 5, 6}
	slice1 := array[1:5]

	fmt.Printf("slice 1= %v\n", slice1)
	fmt.Printf("length 1= %d\n", len(slice1))
	fmt.Printf("capacity = %d\n", cap(slice1))

	//pas by reference

	var mapKosong1 = make(map[string]string)
	var mapKosong2 map[string]string

	fmt.Println(mapKosong1 == nil)
	fmt.Println(mapKosong2)
}
