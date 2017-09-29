package vis

import "fmt" // this does not get package level scope (imports)


func PrintVar() {
	fmt.Println(MyName)
	fmt.Println(yourName) // yourName is unexported, but
	// yourName in this manner ---> this is similar to a 'getter' function from say java, method to access private variables
}
