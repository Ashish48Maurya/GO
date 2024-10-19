package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "cat,bat,mat"
	str_arr := strings.Split(str, ",")
	fmt.Println("str_arr:",str_arr)

	str2 := "one two three one one four two"
	ans := strings.Count(str2,"one")
	fmt.Printf("one appears %d times in [%s]\n",ans,str2)

	str3 := "Ashish"
	fmt.Println(strings.HasPrefix(str3,"s"))

	string1 := "Ashish"
	string2 := "Maurya"
	string3 := strings.Join([]string{string1,string2}, " ")
	fmt.Println("string3:",string3)
}