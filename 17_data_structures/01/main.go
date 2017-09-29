package main

import "fmt"

func main() {
	mySlice := []int{1,3,5,7,9,11}
	otherslice := []string{"a","b","c","g","m","z"}
	
	fmt.Printf("%T\n", mySlice)
	fmt.Println(mySlice)

	fmt.Println(otherslice)
	fmt.Println(otherslice[2:4])
	fmt.Println(otherslice[2])
	fmt.Println("otherslice"[2]) // indexing the "h" in "otherslice", not the variable above.  String is
	// an array of bytes (or runes)

	var slice2 []string
	slice2 = make ([]string,10,15)
	slice2[9] = "bob"
//	slice2[12] = "blah" // out of range error, even though cap is 15?

	slice3 := make([]int, 0, 5)

	fmt.Println("---")
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
	fmt.Println("---")

	for i:= 0;i<80;i++ {
		slice3 = append(slice3,i)
		fmt.Println("Len:", len(slice3), "Capacity:", cap(slice3), "Value: ",slice3[i])
	}
	
	slice4 := []int{1,3,5,7,9,11,}

	for i, value := range slice4 {
		fmt.Println(i, " - ", value)
	}
	
	slice5 := []string {
		"Good morning!",
		"bonjour!",
		"dias!",
		"bongiorno!",
		"ohayo!",
		"selamat pagi!",
	}

	for i, currentEntry := range slice5 {
		fmt.Println(i, currentEntry)
	}

	for j:=0;j < len(slice5);j++ {
		fmt.Println(slice5[j])
	}

	fmt.Print("[1:2] ")
	fmt.Println(slice5[1:2])
	fmt.Print("[:2] ")
	fmt.Println(slice5[:2])
	fmt.Print("[5:] ")
	fmt.Println(slice5[5:])
	fmt.Print("[:] ")
	fmt.Println(slice5[:]) // prints everything!

	// Index out of range error:
	slice6 := make([]string, 3, 5)

	slice6[0] = "one"
	slice6[1] = "two"
	slice6[2] = "three"
	//slice6[3] = "four" // triggers out of range error

	fmt.Println(len(slice6))
	fmt.Println(cap(slice6))
	
	// Appending
	slice6 = append(slice6, "four")
	slice6 = append(slice6, "five")
	slice6 = append(slice6, "six")
	
	
	fmt.Println(len(slice6))
	fmt.Println(cap(slice6))

	// Append slice to slice
	// Can do with ints, strings, other data types

	slice7 := []int{1,2,3,4,5}
	slice8 := []int{6,7,8,9}

	slice7 = append(slice7, slice8...)
	fmt.Println(slice7)

	// Deleting from slice, remove third element
	slice7 = append(slice7[:2], slice7[3:]...)
	fmt.Println(slice7)

	slice9 := []int{10,11,12,13,14,15}
	fmt.Println(slice9)

}
