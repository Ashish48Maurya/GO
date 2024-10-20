package main

import (
	"encoding/json"
	"fmt"
)
type Person struct{
	Name string `json:"name"`
	Age int `json:"age"`
	IsStudent bool `json:"student"`
}
func main(){
	person := Person{Name:"Ashish" , Age:20, IsStudent:true}
	jsonData,err := json.Marshal(person)
	if err!=nil{
		fmt.Println("Error while Marshalling:",err)
		return
	}
	fmt.Println("Converting Go Object to Json String: ",string(jsonData))


	fmt.Println("UnMarshalling the data: Converting Json Data to GO Object")
	var personData Person
	err = json.Unmarshal(jsonData,&personData) //storing unmarshalled data into the refferenced variable
	if(err!=nil){
		fmt.Println("Error while Marshalling:",err)
		return
	}
	fmt.Println("Converting Json Data to Go Object: ",personData)
}