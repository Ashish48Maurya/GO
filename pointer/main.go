package main

import "fmt"

func modifyNum(num *int){
	*num = *num+2;
}

func main() {
	// var num int = 4
	// fmt.Println("num:",num)
	// var ptr *int = &num
	// fmt.Println("Value stored in ptr:",ptr)
	// fmt.Println("Value pointing by ptr:",*ptr)


	//shorthand
	num := 5;
	fmt.Println("num:",num)
	ptr := &num
	fmt.Println("Value stored in ptr:",ptr)
	fmt.Println("Value pointing by ptr:",*ptr)

	var pointer *int;		//default value of the pointer is nil
	if pointer==nil{
		fmt.Println("nil")
	}
		
	change_num := 10
	fmt.Println("change_num",change_num)
	modifyNum(&change_num)
	fmt.Println("change_num",change_num)

}