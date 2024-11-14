package main 

import "fmt"

func main(){

	var name string = "Aman"
	var age int = 22
	
	fmt.Println("NAME : ",name)
	fmt.Println("AGE : ",age)

	//variable decalration shorthand , works inside function
	fName := "Aman"
	lName := "Kumar"

	fmt.Println("Full Name :",fName+" "+lName)

}