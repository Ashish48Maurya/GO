package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Todo struct {
	UserId    int    `json:"userId`
	Id        int    `json:"id`
	Title     string `json:"title`
	Completed bool   `json:completed`
}

func main() {
	todo := Todo{
		UserId:    123,
		Title:     "Ashish Maurya",
		Completed: false,
		Id: 1,
	}
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error while Marshalling:", err)
		return
	}

	// const myURL = `https://jsonplaceholder.typicode.com/todos/${todo.ID}`
	const baseURL = "https://jsonplaceholder.typicode.com/todos/"
	myURL := fmt.Sprintf("%s%d", baseURL, todo.Id)

	jsonReader := strings.NewReader(string(jsonData))
	// req,err := http.NewRequest("PUT",myURL,jsonReader)       OR
	req,err := http.NewRequest(http.MethodPut,myURL,jsonReader)
	if(err!=nil){
		fmt.Println("Error creating PUT Request:", err)
		return;
	}
	req.Header.Set("Content-type","application/json")

	client := http.Client{}
	res,err:=client.Do(req)
	if(err!=nil){
		fmt.Println("Error sending req:", err)
		return
	}
	defer res.Body.Close()

	data,_ := io.ReadAll(res.Body)
	fmt.Println("data Updated Successfully:",string(data))
	fmt.Println("Res Status:",res.Status)
}
