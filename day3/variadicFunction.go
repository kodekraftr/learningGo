package main

import "fmt"

func main(){
	fmt.Println("Largest number is ", largest(2,3,42,1,9))
}

func largest(numbers ...int) int {
	large := 0
	for _,num := range numbers{
		if large < num {
			large = num
		}
	}
	return large
}