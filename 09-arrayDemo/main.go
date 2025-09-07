package main

import "fmt"

func main() {

	var num [5]int = [5]int{1,2,3,4,5}
	
	for _, v := range num {
		fmt.Println(v)
	}

	num[2] = 100
	fmt.Println(num)
}