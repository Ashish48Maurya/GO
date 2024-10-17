package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// fmt.Printf("Enter U'r Name: ")  //for Ashish Maurya getting only Ashish, coz Scan reads only till the white space
	// var name string;
	// fmt.Scan(&name)
	// fmt.Println("U'r Name is:",name)

	fmt.Printf("Enter U'r Full Name: ")  //for Ashish Maurya getting only Ashish, coz Scan reads only till the white space
	reader := bufio.NewReader(os.Stdin)
	fullname,_ := reader.ReadString('\n')
	fmt.Println("Hello Mr.",fullname)
}