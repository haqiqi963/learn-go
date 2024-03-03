package main

import "fmt"

func main() {
	var a1 = [4]string{"Haqiqi", "Ganteng", "Banget"}
	a2 := [...]int{1, 2, 3, 4, 5, 6, 7, 8}

	a2[2] = 100

	fmt.Println(a1[0])
	fmt.Println(a2[2])
}
