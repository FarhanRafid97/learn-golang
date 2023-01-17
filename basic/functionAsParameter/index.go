package main

import "fmt"

type TypeFilter func(string)string

func sayWord(name string,filter TypeFilter)string{
	return ("Hello " + filter(name))
}

func filterWord(name string)string{
	if(name == "anjing"){
		return "****"
	}
	return name

}
func main (){
	fmt.Println(sayWord("farhan",filterWord))
}