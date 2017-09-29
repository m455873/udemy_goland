package main

import (
	"fmt"
	"github.com/m455873/udemy_goland/04_scope/vis"
	// above is the fully qualified package name
)

var x int = 42 // accessible in this whole package, so package level scope... outside of function or block

func max(x int) int { // int return value
	return 42 + x
}


func main() { // block level scope, variables here can't be seen outside this block
	fmt.Println(x)
	foo()
	y := 17
	fmt.Println(y) // accessible
	fmt.Println(z) // accessible

	max := max(7) // the result and function call are same name, but max now refers to variable declared here
	// above line is variable shadowing -- can't call function max anymore because of non-unique name
	// will not throw error, but considered sloppy
	fmt.Println(max)

	fmt.Println(vis.MyName)
	vis.PrintVar()
}

func foo() {
	fmt.Println(x)
	//fmt.Println(y) // not accessible --> error
}

var z = 52 // accessible, because scope outside of blocks is declared first


