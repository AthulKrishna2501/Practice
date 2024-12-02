package main

import (
    "fmt"
    "time"
)


func friend(id int, tasks <-chan int) {
    for task := range tasks {
        fmt.Printf("Friend %d started task %d\n", id, task)
        time.Sleep(time.Second) 
        fmt.Printf("Friend %d finished task %d\n", id, task)
    }
}

func main() {
    tasks := make(chan int, 5) 
    for f := 1; f <= 3; f++ {
        go friend(f, tasks)
    }

    
    for t := 1; t <= 5; t++ {
        tasks <- t
    }
    close(tasks) 

    time.Sleep(5 * time.Second) 
}
