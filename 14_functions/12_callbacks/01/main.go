package main

import "fmt"

func filter(numbers []int, callback func(int) bool) []int { // slice int array, callbackfunc, takes int, returns bool... filter function returns int array slice
	xs := []int{}  // empty list
	for _, n := range numbers {
		if callback(n) { // if a callback returns true, append number  from list onto array xs
			xs = append(xs,n)
		}
	}

	return xs
}



func main () {
	xs := filter([]int{1,2,3,4}, func(n int) bool {
		return n > 1
	})
	fmt.Println(xs)
}

// Call anonymous function filter, filter does callback on the return n>1 which is checked in the
// for loop for values of 1,2,3,4... So, where those numbers are greater than one, the get appended
// onto the xs array (returned from filter), and printed... unfortunately here, all the variables are same
// names between scopes (silly... but legal)
// filter is anonymous
// the anonymous function passed in for callback starts here... func(n int) bool {...}
