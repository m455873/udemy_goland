package main

import "fmt"

func main() {
	for i := 0; i< 5; i++ {
		fmt.Printf("%d: ",i)
		fmt.Println(takesint(i))
	}
}

func takesint(z int) (m int,n int) {
	m = z % 2
	if m == 0 {
		n = 1
	}
	return m,n
}


