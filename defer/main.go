package main

import "fmt"

func main() {
	//defer is used to free the resources used while performing some task (e.g) file read,write etc.


	fmt.Println("------------Learning differ------------")
	// fmt.Println("Start of Execution")
	// fmt.Println("Middle of Execution")
	// fmt.Println("End of Execution")

	fmt.Println("Start of Execution")
	defer fmt.Println("Middle of Execution")  //will execute just before the end of main block  
	fmt.Println("End of Execution")

	fmt.Println("if there are 2 differs, than they follow LIFO order")
	defer fmt.Println("Hello-2")
	defer fmt.Println("Hello-3")
	fmt.Println("Hello-4")
}