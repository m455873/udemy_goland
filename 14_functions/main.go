package main

import "fmt"

func main() {
	avg := average (43,56,87,12,45,57) // 43, 56... are of type 'kind', as a literal constant -- don't have to do Int(43), Int(56)... etc  perl and python do something similar
	fmt.Println(avg)

	data := []float64{43,65,92,15,23,89} // slice of float64 -- we aren't passing in a list,
	// each item gets passed in individually out of the slice
	m := average(data...)
	fmt.Println(m)

	// NExt here we have just using a list
	dataq := []float64{45,32,23,13,79,90} // notice, curly brackets around list of ints
	q := listaverage(dataq)
	fmt.Println(q)
}

func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T\n",sf)
	total := 0.0
	for _, v:= range sf { // range sf gives number of items in the ... what is it? A list of float64
		// above, notice the _ --> blank identifier.  Why not use len(sf) rather than range sf
		total += v
	}

	return total / float64(len(sf)) // len gives you length of list
}


func listaverage(sf []float64) (total float64) {
	for _,v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

	
