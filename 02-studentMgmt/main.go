package main

import(
	"fmt"
)

func main(){

	var idNumber int

	type information struct{
		name string
		age int
	}

	check := map[int]information{
		123 : {"shahrzad" , 17},
		456 : {"fatemeh" , 17},
		789 : {"donya" , 17},
	}

	fmt.Print("Enter your ID number => ")
	fmt.Scan(&idNumber)

	_, ok := check[idNumber]
	if ok{
		fmt.Println("Exist")
	}else{
		fmt.Println("Not found")
	}
	
}