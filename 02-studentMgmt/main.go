package main

import(
	"fmt"
)

func main(){

	var idNumber int

	name := map[int]string{
		1 : "shahrzad",
		2 : "Fatemeh",
		3 : "Donya",
	}
	// age := map[int]int{
	// 	1 : 17,
	// 	2 : 17,
	// 	3 : 22,
	// }

	fmt.Print("Enter your ID number => ")
	fmt.Scan(&idNumber)

	_, ok := name[idNumber]
	if ok{
		fmt.Println("The desired student exists")
	}else{
		fmt.Println("Not found")
	}

}