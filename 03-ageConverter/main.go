package main

import(
	"fmt"
	"strconv"
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

	person, ok :=  informationValue[idNumber]
	if ok{
		fmt.Println(".............................")
		ageString := strconv.Itoa(person.age)
		fmt.Printf("The student with name: %s and age: %s\n", person.name, ageString)
		fmt.Println(".............................")
	}else{
		fmt.Println("...Not found...")
	}

	
	
}