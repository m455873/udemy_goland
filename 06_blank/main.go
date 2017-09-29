package main

import (
	"net/http"
//	"log"
	"io/ioutil"
	"fmt"
)

func main() {
	a := "stored in a"
//	b := "stored in b"
	fmt.Println("a - ", a)
	// b is not being used, this is an error

	res, _ := http.Get("http://www.mcleods.com/") // the _ underscore throws away the error result
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s",page)
}
