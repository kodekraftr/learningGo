package main

import "fmt"

func main(){
	age := 18
	if age <= 16{
		fmt.Println("You are crazy child")
	}else if age <= 35 {
		fmt.Println("You are energetic young")
	}else {
		fmt.Println("You are going old")
	}
}