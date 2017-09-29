package main

import "fmt"

func main() {
	joe := func(y int) (m int, n int) {
		m = y % 2
		if m == 0 {
			n = 1
		}
		return m,n
	}
	
	for i := 0; i< 5; i++ {
		fmt.Printf("%d: ",i)
		fmt.Println(joe(i))
	}

	c,d := takesint(15)
	fmt.Println(c,d)

}

func takesint(z int) (m int,n int) {
	m = z % 2
	if m == 0 {
		n = 1
	}
	return m,n
}


