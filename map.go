package main

import "fmt"

func main(){
	name := map[string]string{
		"name":"farhan",
		"address":"padang",
	}
	name["saya"] = "hadir"
	fmt.Println(name["name"])
	fmt.Println(name["address"])
	fmt.Println(name["saya"])
}