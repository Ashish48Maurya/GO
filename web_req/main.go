package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	res,err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if(err!=nil){
		fmt.Println("Error getting GET Response:",err)
		return
	}
	defer res.Body.Close() //closing the connection after getting the response to save the resource usage
	data,err := io.ReadAll(res.Body)
	if(err!=nil){
		fmt.Println("Error reading Response:",err)
		return
	}
	fmt.Println("Getting array of bytes:",data) //for converting array of bytes data to human readable data use => string(array_of_bytes)
	fmt.Println("Human Readbale data:",string(data))
}