package main

import "fmt"
func main() {
	for i := 5000;i <= 5100;i++{
		fmt.Println(i, " - ", string(i), " - ", [] byte(string(i))) // cast (convert) string of i into byte array, print
	}

	foo := 'a'
	fmt.Println(foo) // prints 97 -- the byte equivalent, why is it stored as a number? (int32... remember)
	fmt.Printf("%q\n",foo)
	fmt.Printf("%T\n",foo)
	fmt.Printf("%T\n",rune(foo))

	
}
