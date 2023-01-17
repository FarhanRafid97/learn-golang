package main

import "fmt"

func main(){
	var month = [...]string{
		"januari",
		"febuary","maret","april","mei","junni","jully","januari2","januari3",
	}
	var slice= month[1:3]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	 

	var slice2 = month[7:]
	
	ass := append(slice2, "saya")
	fmt.Println(ass)

}