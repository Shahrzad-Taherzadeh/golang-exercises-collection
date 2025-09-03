package main

import(
	"fmt"
	"time"
)

func main(){

	start := time.Now()

	for i:=5; i>0; i--{
		fmt.Println("Take a break!")
		time.Sleep(3 * time.Second)
	}

	fmt.Println("All reminders done!")

	end := time.Now()

	elapsed := end.Sub(start)
	fmt.Println("Total elapsed time:", elapsed)
}