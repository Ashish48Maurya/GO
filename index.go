package main //in main directory(i.e Project folder) only one package can exist ( if u try to create the another file containing another package than it will throw an error) if you need to use another package create new folder(Sub-Folder) in that create file and  in this use the pacakge
import (
	"firstCode/myutil"
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	// NOTE :=> if first letter of the variable is Captital than that variable can be used in other packages(file) as well ,(e.g) Println  is the function in fmt file(package) but we can using it in other pacakges simillarly SayHelloAshish (user defind function in index1.go file)
	fmt.Println(myutil.SayHelloAshish())
	myutil.Greet("Ashish")

	//variables
	var name string = "Ashish"
	fmt.Println(name)
	name = "Golang"
	fmt.Println("After Changing:", name)

	var number int = 23
	// var number int = "Ashish"    throw an error because type is int and u'r trying to assigned string to it
	var currency float64 = 23.03
	fmt.Println(number, currency)

	var anyType = "Ashish" //can store any value {string , float, int} if datatype is not assigned
	var anyType1 = 123
	fmt.Println(anyType, anyType1)

	//shorthand
	surname := "maurya" //widely used synatx ,(useCase => ) when u want to store data comming from api's
	fmt.Println(surname)

	//Println and Printf
	fmt.Println("Name:", name, ",Number:", number, ",Surname:", surname)
	fmt.Printf("Name: %s\n",name)
	fmt.Printf("Name: %s, Surname: %s, number: %d\n",name,surname,number)

	fmt.Println("Line-1")
	fmt.Println("Line-2 on new Line")

	fmt.Printf("Line-1 ")
	fmt.Printf("Line-2 on same Line")
}
