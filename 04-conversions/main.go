package main

import (
	"fmt"
	"strconv"
)

func main() {
	price := 23.23
	name := "shahrzad"
	count := int64(123)
	letterByte := byte(65) 
	letterRune := 'a'

	fmt.Println("Float to Int:", int(price))

	
	if val, err := strconv.Atoi("123"); err == nil {
		fmt.Println("String to Int (Atoi):", val)
	} else {
		fmt.Println("Error converting string to int:", err)
	}

	if val, err := strconv.ParseInt("101", 2, 64); err == nil {
		fmt.Println("Binary String to Int64 (ParseInt):", val)
	} else {
		fmt.Println("Error:", err)
	}

	fmt.Println("Int to String (Itoa):", strconv.Itoa(int(count)))

	fmt.Println("String to []byte:", []byte(name))

	fmt.Println("Byte to String:", string(letterByte))

	fmt.Println("Rune to String:", string(letterRune))

	fmt.Println("String to []rune:", []rune(name))
}
