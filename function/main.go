package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func multiply(a,b,c,d,e int) int {
	return a*b*c*d*e
}

func main() {
	fmt.Println("hello")
	fmt.Println("Sum is:",add(3,4))
	fmt.Println("Mul is:",multiply(1,2,3,4,5))
}
