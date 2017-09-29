package main

import "fmt"

func main() {
	if true {
		fmt.Println("This ran")
	}
	if !false { // exclamation mark negates
		fmt.Println("this did not run")
	}

	b := true
	if food := "Chocolate"; b {
		fmt.Println(food)
	}
	
	if false{
		fmt.Println("First")
	} else {
		fmt.Println("second")
	}

	if false {
		fmt.Println("Third")
	} else if true {
		fmt.Println("Fourth")
	} else {
		fmt.Println("Fifth")
	}

	for i:= 0; i<= 100;i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}


}
