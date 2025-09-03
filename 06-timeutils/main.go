package main

import (
    "fmt"
    "time"
)

func main() {
    
    current := time.Now()
    fmt.Println("Current time:", current)

   
    tenMinutesLater := current.Add(10 * time.Minute)
    fmt.Println("Ten minutes later:", tenMinutesLater)

   
    start := time.Now()
    time.Sleep(3 * time.Second) 
    end := time.Now()
    elapsed := end.Sub(start)
    fmt.Println("Elapsed time:", elapsed)

    today := current.Format("2006/01/02")
    fmt.Println("Today's date:", today)

    parsedDate, err := time.Parse("2006-01-02", "2025-12-25")
    if err != nil {
        fmt.Println("Error parsing date:", err)
    } else {
        fmt.Println("Parsed date:", parsedDate)
    }
}
