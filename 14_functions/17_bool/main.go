package main

import "fmt"

func main() {
	if !true {
		fmt.Println("not true")
	}

	if !false {
		fmt.Println("not false")
	}
	
	if true || false {
		fmt.Println("This ran")
	}

	if true && false {
		fmt.Println("This did not run")
	}
}
