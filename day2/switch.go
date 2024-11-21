package main

import "fmt"

func main(){
	day := "Monday"

	switch day{
	case "Monday":
		fmt.Println("hi this is monday")
		fmt.Println("still monday is not over")
	case "Friday":
		fmt.Println("Now this is friday")
		fmt.Println("But Now it's over....")
	default:
		fmt.Println("I don't know what day is today, bye")
	}
}