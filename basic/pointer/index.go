package main

import "fmt"

type Address struct{
	City,Province,Country string
}
type Man struct{
	Name string
}

func changeCity(address *Address){
	*address  = Address{"manggopoh","lubuk basung","danau toba"}

}
func (man *Man) Married(){
	man.Name = "MR " + man.Name
}
func main (){
	adress := Address{"padang","sumatera barat","indonesia"}
	adress2 := &adress
	changeCity(&adress)
	adress2.Province = "riau"

	adress3 := &adress
	// *adress2 = Address{"jawa","jawa barat","singapore"}
	fmt.Println(adress)
	fmt.Println(adress2)
	fmt.Println(adress3)

	farhan := Man{"farhan"}
	farhan2 := &farhan
	farhan2 .Married()
	fmt.Println(farhan)
	fmt.Println(farhan2)

}