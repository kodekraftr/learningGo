package main
import "fmt"

func main(){
	fmt.Println(greet("World"))
	fmt.Println("Area : ",area(4,5))

	q, r := divide(10, 3)
    fmt.Println("Quotient:", q, "Remainder:", r)

	a,p := rectangle(10,12)
	fmt.Println("Area:", a, "Perimeter:", p)

	//Anonymous function
	add := func(a int,b int) int{
		return a+b
	}

	fmt.Println("Sum : ", add(5,5))
}


//Single return 
func greet(name string) string{
	return "Hello "+ name
}

func area(a int, b int) int{
	return a*b
}

//Multiple return
func divide(a int,b int) (int,int){
	quot := a/b
	rem := a%b
	return quot,rem
}

func rectangle(length int, width int) (area int,perimeter int){
	area = length*width
	perimeter = 2*(length+width)
	return
}