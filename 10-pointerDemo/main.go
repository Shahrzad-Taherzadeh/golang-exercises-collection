package main

import "fmt"

func main() {
	var num int = 12
	p := &num
	fmt.Println(num)
	fmt.Println(*p)
}