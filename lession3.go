package main

import "fmt"

func main() {

	// control flow

	// if statement
	// for
	// switch

	age := 30 // 30 is an integer literal

	// 30 > 18 then condition is false so else will be executed

	if age < 18 { // if is control flow statement,
		//means it will check the condition and execute the code inside the block if the condition is true
		fmt.Println("You are underage.")
	} else {
		// else is used to execute the code inside the block if the condition is false
		fmt.Println("You are an adult.")
	}

	// for loop
	// for is to repeat statement 100 times, i is the counter variable, it starts from 0 and ends at 99
	for i := 0; i < 100; i++ { // 99 because 100 is not included in the loop
		// here print 0 to 99
		fmt.Print(i, " ")
	}

	// here is an array of 3 strings
	arr := []string{"a", "b", "c"}
	// we use range to loop 3 times, a is the value of the array at the current index
	// remember index is first in order in for range
	for index, a := range arr {
		fmt.Println(a, " ", index)
	}

	// we have also for infinite

	//	for {
	//	fmt.Println("Hello, World!")
	//	}

	// finally we have switch statement
	// switch is used to perform different actions based on different conditions

	age = 18
	// becuase 18 is not 19 or 20 so default will be executed
	switch age {
	// case is used to check the condition and execute the code inside the block if the condition is true
	case 19:
		fmt.Println("You are 19 years old.")
	case 20:
		fmt.Println("You are 20 years old.")
	default:
		fmt.Println("You are not 19 or 20 years old.")
	}

	fmt.Println("Hello, World!")

}
