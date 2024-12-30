package main

import (
	"fmt"
	"math"
)

func main(){
	arr := [5]int{11,12,13,14,15}
	arr2 := [5]float64{}
	fmt.Println(arr) //[11 12 13 14 15]
	fmt.Println(arr2) //[0 0 0 0 0]

	for i := 0 ; i < 5 ; i++ {
		fmt.Println(arr[i]) // 11 12 13 14 15
	}

	for index,value := range arr2{
		fmt.Println(index , value)
		arr2[index] = math.Pow(10,float64(index)) * float64(index)
		fmt.Println(arr2[index])
	
	}
}