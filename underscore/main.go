package main
import "fmt"

// func divide(a, b int) (int, string) {
// 	if b == 0 {
// 		return 0, "Cannot divide by zero"
// 	}
// 	return a / b, "nil"
// }

// func main(){
// 	quotient, err := divide(10, 5)
//     if err!= "nil" {
//         println(err)
//     } else {
//         fmt.Println("Quotient:", quotient)
//     }
// }

// func main() {
// 	ans, _ := divide(10, 5) //use underscore if you want to handle only first result of return statement when want to ignore the error part
// 	fmt.Println("Ans is:",ans)
// }

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by 0")
	}
	return a / b , nil
}

func main(){
	quotient, err := divide(10, 2)
	// quotient, err := divide(10, 0)
    if err!= nil {
        fmt.Println(err)
    } else {
        fmt.Println("Quotient:", quotient)
    }
}