package main

import "fmt"

func main(){

	arr := [6]int{1,2,3,4,5}
	slice := arr[1:] //slice from array
	//arr = append(arr,1) //error - first argument to append must be a slice; have arr (variable of type [6]int)
	fmt.Println(arr)
	slice = append(slice,1) //we can add in slice but not arr
	fmt.Println(slice)

	slice2 := make([]int,0)
	slice2 = append(slice2,2,3,4,5)
	fmt.Println(slice2)
	
	slice3 := []int{1} //no length mentioned
	fmt.Println(slice3) 

	slice4 := slice2
	slice4 = append(slice4,1)
	fmt.Println(slice4)
}