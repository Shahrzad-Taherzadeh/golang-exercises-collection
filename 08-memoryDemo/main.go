package main

import "fmt"

type User struct{
	Name string
	Age int
}

func CreatUser(name string, age int) *User{
	return &User{Name: name, Age: age}
}

func main() {

	u1 := CreatUser("Shahrzad", 17)
	fmt.Println("User1", u1)

	u2 := CreatUser("Fateme", 17)
	fmt.Println("User1", u2)

}

// PowerShell👇
// $env:GODEBUG="gctrace=1"
// go run memory_demo.go

// CMD (Command Prompt)👇
// set GODEBUG=gctrace=1
// go run memory_demo.go

// Git Bash👇
// GODEBUG=gctrace=1 go run memory_demo.go

// Linux/Mac (bash/zsh)👇
// GODEBUG=gctrace=1 go run memory_demo.go

