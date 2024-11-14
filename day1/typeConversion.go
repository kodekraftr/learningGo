package main

import "fmt"
import "reflect" 
import "strconv"

func main(){
	x := 45.5
	y := 45

	s := x+ float64(y)

	fmt.Println("Type of s : ", reflect.TypeOf(s))
	fmt.Println("S : ",s)

	t := int(s)
	fmt.Println("Type of T : ", reflect.TypeOf(t))
	fmt.Println("T : ",t)

	num := 123
	str := strconv.Itoa(num)
	fmt.Println("Type of str : ", reflect.TypeOf(str))
	fmt.Println("str : ",str)

	n,err := strconv.Atoi(str)
	if err == nil{
		fmt.Println("Type of n : ", reflect.TypeOf(n))
		fmt.Println("str : ",n)
	}

}