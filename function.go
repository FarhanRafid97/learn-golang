package main

import "fmt"



func penambahan(a int16,b int16)int16{
	var hasil int16 = a + b 
	return hasil
	
}
func main (){
	hasil := penambahan(20000,3002)
	fmt.Println(hasil)
}
