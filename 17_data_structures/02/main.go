package main

import "fmt"

func main() {
	records := make([][]string,0)

	student1 := make([]string, 4)
	student1[0] = "Foster"
	student1[1] = "Nathan"
	student1[2] = "100.00"
	student1[3] = "74.00"

	records = append(records, student1)

	student2 := make([]string, 4)
	student2[0] = "Gomez"
	student2[1] = "Lisa"
	student2[2] = "92.00"
	student2[3] = "96.00"

	records = append(records, student2)

	fmt.Println(records)

	transactions := make([][]int, 0, 3) // size 0, cap 3
	for i:= 0;i < 3;i++ {
		transaction := make([]int,0)
		for j := 0; j < 4; j++ {
			transaction = append(transaction, j)
		}
		transactions = append(transactions,transaction)
	}

	fmt.Println(transactions)


	// initializing slices
	var slice1 []string // value is initially nil 
	slice2 := make([]string, 35,100)
	 // actually builds an array, datatype would be []string

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice1 == nil)

	// Can assign values up to 35 without doing append

	slice2[34] = "blag"

	// Beyond 35, would need to do append, but wouldn't trigger underlying array recreation
	// of large size until hit 100
	// Always have to do append using "var slicex []string", because your underlying size is 0

	//slice2[35] = "blarg" // error
	slice2 = append(slice2, "narg")

	slice3 := make([]int,1)
	fmt.Println(slice3[0])
	slice3[0] = 6
	fmt.Println(slice3[0])
	slice3[0]++ // add one (integer) to the value at slice element [0]
	fmt.Println(slice3[0])
	
	
}
