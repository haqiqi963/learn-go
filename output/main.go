package main

import "fmt"

func main() {
	var firstName, lastName = "Haqiqi", true
	a, b := 100, 200

	nickName, address := "Haqiqi", "Surabaya"

	var x = "Haqiqi"
	var y = 21

	fmt.Print(firstName, " ", lastName, "\n")
	fmt.Print(a, b, "\n")

	fmt.Println(nickName, address)

	fmt.Printf("value x adalah : %v dan tipe data x adalah %T\n", x, x)
	fmt.Printf("value x adalah : %v dan tipe data y adalah %T\n", y, y)

}
