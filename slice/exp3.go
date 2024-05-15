package main
import "fmt"

func main(){
	slice1 := make([]*int, 0)
	for i := 0; i < 5; i++{
		slice1 = append(slice1, &i)
	}

	num := make([]int, 0)
	for _, v := range slice1{
		num = append(num, *v)
	}
	// Go1.21之前是5，5，5，5，5
	fmt.Println(num)
}