package main

import "fmt"

func main() {
	var x [58]string // numbers in brackets is array, slice has no numbers in brackets
	// var y []string --> slice
	// Arrays are not dynamic
	// string empty (default) value is space

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])

	for i := 65; i <= 122; i++ {
		x[i-65] = string(i)
	}

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])

	fmt.Println("----------- int")
	var y [256]int
	fmt.Println(len(x))
	fmt.Println(x[42])

	for j := 0;j < 256;j++ {
		y[j] = j
	}
	for i,v := range x {
		fmt.Printf("%v - %T - %b\n",v ,v ,v)
		if i > 50 {
			break
		}
	}
	fmt.Println("------------ byte type")

	var z [256]byte

	fmt.Println(len(z))
	fmt.Println(z[0])
	for i := 0;i < 256;i++{
		z[i] = byte(i)
	}

	for i,v := range z {
		fmt.Printf("%v - %T - %b\n", v,v,v)
		if i > 50 {
			break
		}
	}
	
}
