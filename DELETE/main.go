package main

import (
	"fmt"
	"net/http"
)

func main() {
	const myURL = "https://jsonplaceholder.typicode.com/todos/1"
	req, err := http.NewRequest(http.MethodPut, myURL, nil)
	if err != nil {
		fmt.Println("Error creating PUT Request:", err)
		return
	}
	client := http.Client{}
	res,err:=client.Do(req)
	if(err!=nil){
		fmt.Println("Error sending req:", err)
		return
	}
	defer res.Body.Close()

	fmt.Println("Res Status:",res.Status)
}
