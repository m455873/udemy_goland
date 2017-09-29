package main

import "fmt"

func zero(x int){
	x = 0
}

func hero(x int){
	fmt.Printf("%p\n",&x) //same twice
	fmt.Println(&x)
	x = 0
}

func pero(x *int){
	fmt.Println(x)
	*x = 0
}


func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a

	fmt.Println(b) // prints memory address of a
	fmt.Println(*b) // prints 42, asterisk means dereference.

	*b = 43 // so, to get at pointer (where it points to), use *b, b is just an memory address
	fmt.Println(a) // 43


	x := 5
	zero(x) // just passing value -- passing copy
	fmt.Println(x)

	fmt.Println("next part 04:")
	
	fmt.Printf("%p\n",&x) // two ways of doing print
	fmt.Println(&x) //same
	hero(x) //passing value again, not memory address
	fmt.Println(x)  //no change in x

	fmt.Println("--- pero")
	// Address of x in both scopes (main and pero) stays same)
	// x still 5 now
	fmt.Println(&x)
	pero(&x) // must pass reference to
	// x not 5 anymore
	fmt.Println(x)

}
