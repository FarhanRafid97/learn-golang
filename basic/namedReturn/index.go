package main

import "fmt"

func getFUllaname()(num1,num2,num3 ,hasil int16) {
	num1 = 3
	num2=10
	num3 =2
	hasil = num1 + num2 + num3
	return 
	

}
func main(){
	num1,num2,num3,hasil := getFUllaname()
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(hasil)
}