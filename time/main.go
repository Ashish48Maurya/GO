package main

import (
	"fmt"
	"time"
)

func main() {

	// => 02/01/2006 03:04:05 PM Monday  is the official time when the GO was Launch
	curr_time := time.Now();
	fmt.Println("curr_time:",curr_time)

	formatted_time := time.Now().Format("02-01-2006")   // DD-MM-YYYY
	fmt.Println("formatted time:",formatted_time)

	formatted_time = time.Now().Format("02/01/2006")
	fmt.Println("formatted time-1:",formatted_time)

	formatted_time = time.Now().Format("02/01/2006 15:04:05 Monday")  //15 for 24 hour format , 3 for 12 hour format
	fmt.Println("formatted time along with time and Day:",formatted_time)

	formatted_time = time.Now().Format("02/01/2006 3:04 PM Monday")  //15 for 24 hour format , 3 for 12 hour format
	fmt.Println("formatted time along with time(without second) and Day with timezone:",formatted_time)

	formatted_time = time.Now().Format("2006/02/01")
	fmt.Println("formatted time-2:",formatted_time)


	new_date := curr_time.Add(24*time.Hour) //Add 1 day to current day
	fmt.Println("new_date:",new_date)
	formatted_new_date := new_date.Format("02/01/2006 Monday")
	fmt.Println("formatted_new_date:",formatted_new_date)

	new_date = curr_time.Add(48*time.Hour) //Add 2 day to current day
	fmt.Println("new_date:",new_date)
	formatted_new_date = new_date.Format("02/01/2006 Monday")
	fmt.Println("formatted_new_date:",formatted_new_date)

	res_date := "2006-01-02" //layout
	format_date := "2003-01-25" //value
	formatted_date,_ := time.Parse(res_date,format_date)
	fmt.Println("Formatted Date:",formatted_date)
}