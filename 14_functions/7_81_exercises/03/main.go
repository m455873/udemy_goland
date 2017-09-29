package main

import "fmt"

func main() {
	data := []int{4,10,15,20,25}
	fmt.Println(vari(data...))
	
}

func vari(jf ...int ) int {
	var retval int
	for _, v:= range jf {
		if v > retval {
			retval = v
		}
	}
	return retval
}
	
	
