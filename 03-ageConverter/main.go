package main

import(
	"fmt"
)

func main(){

	var  idNumber int

	type information struct{
		name string
		age int
	}

	informationValue := map[int]information{
		1234 : {"Shahrzad", 17},
		5678 : {"Donya", 22},
		9123 : {"Fatemeh", 17},
	}

	fmt.Println("Enter your ID number => ")
	fmt.Scanln(&idNumber)

	_, ok :=  informationValue[idNumber]
	if ok{
		fmt.Println("Exist")
	}else{
		fmt.Println("Not found")
	}
}