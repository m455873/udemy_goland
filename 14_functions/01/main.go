package main

import "fmt"

func main() { // only one func main in all program, must be in package main
	// func main (parameters ...) returns here { ... }
	// All 

	greet("jane")
	greet("John")

	greetme("joe", "moe")
	greetme("mike","dike")

	fmt.Println(greetyou("Jane ", "Doe"))

	s:= greetyou("Joe ", "Flow")
	fmt.Println(s)

	fmt.Println(greetx("Zax ", "Slacks"))
	fmt.Println(greety("Bar ", "Fridge"))
}

func greet(name string){ // name variable, type string
	fmt.Println(name)
}

func greetme(firstname string, lastname string){
	fmt.Printf("%s, %s\n",firstname,lastname)
	fmt.Println(firstname,lastname)
}

func greetyou(fname, lname string) string {
	return fmt.Sprint(fname, lname) //Sprint, string print --> returns string rather than to STDOUT
}

func greetx(fname string, lname string) (s string) { // s string is the named return, it declares var, you
	// just have to use it (assign something to it) within the function
	s = fmt.Sprint(fname, lname)
	return
}

func greety(fname, lname string) (string, string) { // as above in greetx, can define type for each parameters
	// in call, or can lump together as in this one.  This has two return values, both of type string and unnamed
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname) // Note how return has two items separated by comma

	// could go func greety(...) (s string, t string) { s = "this", t = "that", return s,t}p
}


	
