package main // define program as execuatable

import "fmt" // package for printing to console

func main() { // define main function

	// hello, we can declare variable in two way
	// 1. var name type = value
	// 2. name := value (only inside function)

	var name string = "Mahmoud" // declare variable name with type string and value "Mahmoud"

	age := 32 // shorthand declaration of variable age with value 32

	fmt.Println(name, age) // print name and age to console

}
