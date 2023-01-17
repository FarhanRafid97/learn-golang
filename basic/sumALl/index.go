package main

import "fmt"

func sumAll(numbers ...int)int{
total := 0
for _, number := range numbers{
	total += number
}
return total
}
func main(){
fmt.Println(sumAll(2,3,4,4,4,))
slice := []int{5,5,5,5,5}
fmt.Println(sumAll(slice...))
}