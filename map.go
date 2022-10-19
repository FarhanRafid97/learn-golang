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

	var song map[string]string = make(map[string]string)
	song["title"] = "That should be me"
	song["artist"] = "bruno mars"
	song["realese"] = "2022"

	fmt.Println(song)
	fmt.Println(len(song))
	delete(song,"realese")
	fmt.Println(song)
	fmt.Println(len(song))
}
