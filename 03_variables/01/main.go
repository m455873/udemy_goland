package main

import "fmt" // the format/printing standard library thingy
// ardlan labs go course

func main() {

	var a int // default 0
	var b string  // default empty string
	var c float32 // default 0
	var d bool // default false

	fmt.Printf("%v, %T \n", a,a)
	fmt.Printf("%v, %T \n", b,b)
	fmt.Printf("%v, %T \n", c,c)
	fmt.Printf("%v, %T \n", d,d)

	// can't do a:= 10 below because its already declared!!!
	
	a = 10
	b = "golang"
	c = 4.17
	d = true

	// this is called the shorthand method, rather than the var, assign.  shorthand is prefered
	e := 10
	f := "golang"
	g := 5.17
	h := false
	
	fmt.Printf("%v, %T \n", a,a)
	fmt.Printf("%v, %T \n", b,b)
	fmt.Printf("%v, %T \n", c,c)
	fmt.Printf("%v, %T \n", d,d)

	fmt.Printf("%v, %T \n", e,e)
	fmt.Printf("%v, %T \n", f,f)
	fmt.Printf("%v, %T \n", g,g)
	fmt.Printf("%v, %T \n", h,h)
}
