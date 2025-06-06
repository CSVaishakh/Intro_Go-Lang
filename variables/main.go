package main

import "fmt"

func main() {
	//ways for declaring variables of a type
	
	// 1) Strings 

	// variable values are mutable
	// d type of a variable cannot be muted

	var string1 string = "C S" // d type is defined

	var string2 = "Vaishakh" // d type is detected

	var string3 string
	string3 = "C S Vaishakh"

	fmt.Println(string1, string2, string3)

	// another way to decalre a varibale is :
	// this type of variable declaratuion is only possbile inside functions
	// automaticly detects the type
	string4 := "C S V"
	fmt.Println(string4)

	// 2)Integers
	//  using var keyword
	var age1 int = 20 // d type is specified
	
	var age2 = 25 // dtype is detected
	
	//  using : and =
	age3 := 16 // dtype is detected

	fmt.Println(age1,age2,age3)

}