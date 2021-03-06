package main

import "fmt"

func main() {
	var myGreeting = make(map[string]string) // if you don't use make, you can't use map cause its nil
	//var map2 = map[string]string{} // also valid
	//var mapnil map[string]string // error to use, nil type
	myGreeting["Tim"] = "GOod morning"
	myGreeting["Jenny"] = "Bonjour"

	fmt.Println(myGreeting)

	// composite literal
	map3 := map[string]string{
		"Joe":"hello",
		"fred":"jello",
	}

	map3["Joe"] = "howdy"

	delete(map3,"fred")
	
	if val, exists := map3["fred"];exists {
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	} else {
		fmt.Println("That value doesn't exist")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	}
	
	map4 := map[int]string{
		0:"fred",
		1:"joe",
		2:"mike",
	}

	fmt.Println(map4)

}
