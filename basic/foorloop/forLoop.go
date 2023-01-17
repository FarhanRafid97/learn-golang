package main

import "fmt"

func main(){

	slice := [...]string{"farhan","rafid","syauqi"}
	
	for counter :=0; counter < len(slice);counter ++{
			fmt.Println(slice[counter])
	}
}