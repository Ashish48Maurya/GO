//slice is dynamic array in GO
package main
import "fmt"
var numbers = [5]int { 1, 2, 3, 4, 5}
var names = [4]string {"a", "b", "c", "d"}
var arr1 = [5]int{1:10,2:40}

var myslice = []int{1,2,3}

func main() {
	fmt.Println("Arr1:",arr1)
	fmt.Println("Array Res:", numbers)
	fmt.Println("Length of Array:",len(numbers))
	fmt.Println("Names:",names)


	//slice
	fmt.Println(myslice)
	myslice = append(myslice, 4, 5, 6,7)
	fmt.Println(myslice)
	fmt.Println("Capacity: ",cap(myslice))
	fmt.Println("Length:",len(myslice))

}
