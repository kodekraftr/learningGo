package main
import "fmt"

func main(){
	num := 0

	increment := func() int{
		num += 1
		return num
	}

	fmt.Println("Value of num : ", increment())
	fmt.Println("Value of num : ", increment())
	fmt.Println("Value of num : ", increment())

	numbers := []int{1,2,3,4,5,6,7,8,9}
	isEven := func(n int) bool{
		return n%2 == 0
	}

	fmt.Println("Even Numbers : ", filter(numbers,isEven))

	isOdd := func(n int) bool{
		return n%2 != 0
	}
	fmt.Println("Odd Numbers : ", filter(numbers,isOdd))
}

func filter(numbers []int, test func(n int) bool) []int{
	result := []int{}
	for _,num := range numbers{
		if test(num){
			result = append(result,num)
		}
	}
	return result
}