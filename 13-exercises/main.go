package main

import "fmt"

func main() {
	fmt.Println("ex.1: Hello World!\n")

	name := "Mike"
	fmt.Printf("ex.2: Hello %s\n\n",name)

	fmt.Println("ex.3: Please enter your name to STDIN, press enter to continue:")
	
	fmt.Scan(&name) // need to send reference to  variable

	fmt.Printf("ex.3: Your name is: %s\n\n",name)

	var smallno int32
	var bigno int32
	
	fmt.Println("ex.4 Enter a small number:")
	fmt.Scan(&smallno)
	fmt.Println("ex.4 Enter a large number:")
	fmt.Scan(&bigno)

	fmt.Printf("ex.4: big div small remainder: %d\n\n", bigno % smallno)

	fmt.Println("ex. 5: even numbers between 0 and 100")
	for i:= 0;i <= 100;i++ {
		if i % 2 == 0 {
			fmt.Println(i)
		}
	}

	fmt.Println("ex.6: fizzbuzz")
	for j := 1;j<= 100;j++ {
		fmt.Printf("Number: %d\t",j)
		if j % 3 == 0 {
			fmt.Printf("fizz")
		}
		if j % 5 == 0 {
			fmt.Printf("buzz")
		}

		fmt.Printf("\n")

	}

	var dasum int
	
	for k := 1; k < 1000; k++ {
		if k % 3 == 0 || k % 5 == 0 {
			dasum += k
		}
	}
	fmt.Printf("ex.7: sum of 3 or 5 multiples below 1000: %d \n",dasum)
		
		
}
