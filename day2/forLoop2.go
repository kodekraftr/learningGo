package main
import "fmt"

func main(){
	items := []string{"Apple","Banana"}
	fmt.Println(items)

	for index,item := range items{
		// fmt.Printf("Index: %d, Item: %s\n", index, item)
		fmt.Println("Index:",index,", Item:",item)
	}
}