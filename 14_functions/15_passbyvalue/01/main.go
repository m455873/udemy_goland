package main

import "fmt"

func changeMe(z []string) { // this is a reference type, no need to double deref
	z[0] = "todd"
	fmt.Println(z)
}

func changeThis(z map[string]int) {
	z["Todd"] = 44
}

func main() {
	m := make([]string,1,25)  // this is a reference type (slice, map or channel) so just pass
	// the value itself (the address) rather than a reference to...
	fmt.Println(m)
	changeMe(m)
	fmt.Println(m)

	fmt.Println("map:")

	n := make(map[string]int)
	changeThis(n)
	fmt.Println(n["Todd"])

	func() { // no receiver, no identifier, no parameters
		fmt.Println("Im loving it")
	}() // parens after curlys executes anonymous function
}


