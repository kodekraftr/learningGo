package main
import "fmt"

func main(){
	counter := 0
	for{
		counter++

		fmt.Println("Counter ",counter)
		if counter == 100{
			break
		}
	}
}