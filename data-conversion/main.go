package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 123
	fmt.Printf("Type of num is:%T\n", num)
	fmt.Println("num:",num)
	num_string := strconv.Itoa(num)
	fmt.Printf("Type of num_string is:%T\n", num_string)
	fmt.Println(num,num_string)

	num_float := float64(num)
    // float_num  := float64(6.4) + num  //type mismatch because num is int
    float_num  := float64(6.4) + num_float  //type mismatch because num is int
    fmt.Println(float_num)
	fmt.Printf("Type of num_float is:%T\n", num_float)
	fmt.Println("num_float: ",num_float)

	str := "Ashish"
	fmt.Printf("Type of str is %T\n",str)
	str_num,_ := strconv.Atoi(str)
	fmt.Printf("Type of str_num is %T",str_num)
}
