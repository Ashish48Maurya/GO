package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId`
	Id        int    `json:"id`
	Title     string `json:"title`
	Completed bool   `json:completed`
}

func main() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	// if(err!=nil){
	// 	fmt.Println("Error getting GET Response:",err)
	// 	return
	// }

	if res.StatusCode != http.StatusOK {
		// if(res.StatusCode!=200){
		fmt.Println("Error getting GET Response:", res.Status)
		return
	}

	defer res.Body.Close()

											//1st method
	// data, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("Error reading Response:", err)
	// 	return
	// }

	// var todo Todo
	// err = json.Unmarshal(data,&todo)
	// if(err!=nil){
	// 	fmt.Println("Error while Marshalling:",err)
	// 	return
	// }
	// fmt.Println("Converting Json Data to Go Object: ",todo)



											//2nd Method - Shorthand
	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if(err!=nil){
		fmt.Println("Error decoding:",err)
		return
	}
	fmt.Println("Converting Json Data to Go Object: ",todo)
}
