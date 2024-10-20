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
		Title:     "Ashish",
		Completed: true,
	}

	res, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling:", err)
		return
	}
	jsonString := string(res)
	jsonReader := strings.NewReader(jsonString)
	fmt.Println(jsonString, jsonReader)
	myURL := "https://jsonplaceholder.typicode.com/todos"

	res1, err1 := http.Post(myURL, "application/json", jsonReader)
	if err1 != nil {
		fmt.Println("Error sending req:", err1)
		return
	}
	defer res1.Body.Close()
	data,_ := io.ReadAll(res1.Body)
	fmt.Println("Data Saved: ",string(data))
}
