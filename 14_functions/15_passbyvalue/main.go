package main

import "fmt"

func changeMe(z *int) {
	fmt.Println(z)
	fmt.Println(*z)

	*z = 24

	fmt.Println(z)
	fmt.Println(*z)
}

func main() {
	age := 44
	fmt.Println(&age) // address of age

	changeMe(&age)

	fmt.Println(&age)
	fmt.Println(age)
}
