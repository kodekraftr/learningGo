package main
import "fmt"

func main(){
	var num int  // default 0
	var decimal float64 //default 0
	var str string  //default ""
	var isBool bool  //default false 

	fmt.Println("Num : ",num)
	fmt.Println("Decimal Number : ", decimal)
	fmt.Println("String : ",str)
	fmt.Println("Boolean value : ",isBool)

	num = 56 // default 0
	decimal  = 3.14 //default 0.0
	str  = "hiiii" //default ""
	isBool = true //default false 

	fmt.Println("Num : ",num)
	fmt.Println("Decimal Number : ", decimal)
	fmt.Println("String : ",str)
	fmt.Println("Boolean value : ",isBool)
}