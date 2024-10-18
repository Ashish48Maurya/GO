package main
import "fmt"
func main(){
	num := 9;
	if  num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }

	//switch statement
	switch num{
		case 0: fmt.Println(num, "is zero")
        case 1,2,3: fmt.Println(num, "is greater than 4")
        default: fmt.Println(num, "is greater than or equal to 4")
	}

	//for loop
	for i:=0; i<10; i++{
        fmt.Printf("%d ",i)
    }

    //while loop
    i:=0;
    for i<10{
        fmt.Println(i)
        i++
    }

    // //infinite loop
    // for{
    //     fmt.Println("Infinite loop")
    // }

    //break statement
	fmt.Println("Break Statement")
    for i:=0; i<10; i++{
        if i==5 {
            break
        }
        fmt.Println(i)
    }

    // //continue statement
	fmt.Println("continue Statement")
    for i:=0; i<10; i++{
        if i==5{
            continue
        }
        fmt.Println(i)
    }


	number := []int {1,2,3,4,5,6,7,8,9}
	for index,value := range number{
		fmt.Println(index,value)
	}
}