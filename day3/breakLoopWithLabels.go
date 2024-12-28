package main

import "fmt"

func main(){

	OuterLoop:
	for i := 1; i <= 10; i++ {
		fmt.Println("********* I : ",i," *********")
		InnerLoop:
		for j := 1 ; j <= 10; j++{
			fmt.Println("J : ",j)
			if i*j == 30 {
				break InnerLoop
			}

			if i*j == 49{
				break OuterLoop
			}
		}
	}
}