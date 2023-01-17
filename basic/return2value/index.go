package main

import "fmt"

func getFUllaname(num1 int16,num2 int16) (int16,int16){
	
	return num1 * num2,num1 + num2
}
func main(){
	result1, _ := getFUllaname(3,2)
	fmt.Println("Result-1",result1)
	
}