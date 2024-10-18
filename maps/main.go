package main

import "fmt"

func main() {
	studentsGrade := make(map[string]int)
	studentsGrade["Ashish"] = 10
	studentsGrade["Shubham"] = 8
	studentsGrade["Krishna"] = 5

	for index, value := range studentsGrade {
		fmt.Printf("%s Grade is %d\n", index, value)
	}
	fmt.Println("Grade of Krishna:", studentsGrade["Krishna"])

	Grade, Exist := studentsGrade["Krishna"]
	fmt.Println("Krishna Grade:", Grade)
	fmt.Println("Krishna Exist? :", Exist)

	delete(studentsGrade, "Krishna")

	grade, exist := studentsGrade["Krishna"]
	fmt.Println("Krishna Grade:", grade) //returns 0 if object doesn't exists
	fmt.Println("Krishna Exist? :", exist)

	for index, value := range studentsGrade {
		fmt.Printf("%s Grade is %d\n", index, value)
	}
}
