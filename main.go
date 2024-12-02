package main

import (
    "fmt"
    "time"
)

func task(id int) {
    fmt.Printf("Task %d started\n", id)
    time.Sleep(2 * time.Second) 
    fmt.Printf("Task %d completed\n", id)
}

func main() {
    for i := 1; i <= 3; i++ {
        go task(i) 
    }

    time.Sleep(4 * time.Second)
    fmt.Println("Main function completed")
}