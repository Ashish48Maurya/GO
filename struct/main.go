package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	var person Person
	fmt.Println(person) // => { 0} because default value of string is " " and of int is 0
	person.name = "Ashish"
	person.age = 20
	fmt.Println(person)

	person1 := Person{
		name : "Ashish",
		age: 20,
	}
	fmt.Println("Person1:",person1)
}