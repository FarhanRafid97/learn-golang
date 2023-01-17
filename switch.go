package main

import "fmt"

func main ( ){
	 name := "farhan rafid syauqi"

	 switch name {
	 case "farhan":
		fmt.Println("hello",name)
		fmt.Println("hello",name)
		break
	 case "eko":
		fmt.Println("kamu bukan farhan",name)
		fmt.Println("haha",name)
	 default:
		fmt.Println("defualt",name)
		fmt.Println("sss",name)
 	 }

	  length :=len(name)
	 switch  {
	 case length < 5:
		fmt.Println("nama terlalu pendek")
		break
	 case length > 5 && length < 8:
		fmt.Println("nama sudah pas")
	 default:
		fmt.Println("nama terlau panjang")
		
 	 }
	 
}