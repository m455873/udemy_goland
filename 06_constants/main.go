package main

import "fmt"

const p string = "death & taxes"

const (
	PI = 3.14159
	Language = "Go"

	// iotas are basically just going ++ on the variable, imagining that first iota is 0
	// second iota is ++, etc.
	
	A = iota // 2  -- the above variables count as iotas?
	B = iota // 3
	C = iota // 4
)

const (
	D = iota // 0 -- here there is nothing above, so starts from 0
	E = iota // 1
	F = iota // 2
)

const (
	_ = iota // 0 (blank identifier) --> throw away the zero
	G = iota * 10 //1 * 10
	H = iota * 10 //2 * 10
)	

const (
	_ = iota
	KB = 1 << (iota * 10) // bit shift, to left 10 bits
	MB = 1 << (iota * 10) // bit shift to left 20 bits
	GB = 1 << (iota * 10) 
)

func main () {
	const q = 42  // guessing that this is limited to func main scope

	fmt.Println("p- ",p)
	fmt.Println("q - ",q)

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)

	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)

	fmt.Println(G)
	fmt.Println(H)

	fmt.Println("binary\t\tdecimal")
	fmt.Printf("%b\t",KB)
	fmt.Printf("%d\n",KB)
	fmt.Printf("%b\t",MB)
	fmt.Printf("%d\n",MB)
	fmt.Printf("%b\t",GB)
	fmt.Printf("%d\n",GB)

	
}

