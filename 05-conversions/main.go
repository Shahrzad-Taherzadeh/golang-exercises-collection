package main

import (
	"fmt"
	"strconv"
)

func main() {
	var inputNumber string
	fmt.Print("Please enter a number: ")
	fmt.Scanln(&inputNumber)

	floatValue, err := strconv.ParseFloat(inputNumber, 64)
	if err != nil {
		fmt.Println("Error converting input to float:", err)
		return
	}
	fmt.Println("Float value:", floatValue)

	floatStr := strconv.FormatFloat(floatValue, 'f', 2, 64)
	fmt.Println("Float as string (2 decimals):", floatStr)

	intValue := int(floatValue)
	fmt.Println("Float as int:", intValue)

	intStr := strconv.Itoa(intValue)
	fmt.Println("Int as string:", intStr)

	byteArray := []byte(intStr)
	runeArray := []rune(intStr)
	fmt.Println("String to []byte:", byteArray)
	fmt.Println("String to []rune:", runeArray)
}
