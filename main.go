package main

import "fmt"

// func divider(x, y int) (int, error) {
// 	if y == 0 {
// 		return 0, errors.New("error divide by zero")
// 	} else {
// 		result := x / y
// 		return result, nil
// 	}
// }

type Student struct {
	Name    string
	Address string
	Age     int
}

func (student *Student) GetInfo() {
	student.Address = "Address : " + student.Address
}

func main() {
	var student Student
	student.Address = "Sidoarjo"

	student.GetInfo()
	fmt.Println(student.Address)
}
