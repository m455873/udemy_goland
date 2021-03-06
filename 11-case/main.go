package main

import "fmt"

func main() {
	switch "tim" { // regular switch expression (case) with fallthrough
	case "dan":
		fmt.Println("dan")
	case "joe":
		fmt.Println("joe")
		fallthrough // goes to next case statement, rather than default, goes to mike
	case "mike":
		fmt.Println("mike")
		fallthrough // goes to spooky/tim
	case "spooky", "tim":
		fmt.Println("spooky or tim")
	default:
		fmt.Println("none found")
	}

	myFriendsName := "joe"

	switch { // no expression, just runs whichever evaluates to true
	case len(myFriendsName) == 2:
		fmt.Println("2")
	case myFriendsName  == "tim":
		fmt.Println("tim")
	case myFriendsName == "spooky":
		fmt.Println("spooky")
	default:
		fmt.Println("no match")
	}

/*	x := "fiver"

	
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("none")
	}
	*/
}
