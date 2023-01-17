package main

import "fmt"

type Blocked func(name string)bool

func register(name string, blacklist Blocked) {
	 if(blacklist(name)){
		fmt.Println("Youre  Blocked", name)
		return
	}
	fmt.Println("Welcome ", name)
}
func main(){
	isBlacklist := func (name string)bool  {
		return name == "admin"
	}
	 register("admin",isBlacklist)
	 register("farhan",func(name string) bool {
		 return name == "farhan"
	 })
}