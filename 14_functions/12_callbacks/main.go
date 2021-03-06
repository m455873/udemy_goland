package main

import "fmt"

func visit(numbers []int, callback func(int)) {
	for _, n:= range numbers {
		callback(n)
	}
}


func main() {
	visit([]int{1,2,4,5}, func(n int) { // here, func is an anonymous function that takes an int
		fmt.Println(n)
	})
}
