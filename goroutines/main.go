package main

import (
    "fmt"
    "time"
)

func printMessage(msg string) {
    for index , val := range msg {
        fmt.Println(string(val), index)
        time.Sleep(500 * time.Millisecond)
    }
}

func main() {
    // 1. Launch a goroutine using the 'go' keyword!
    // This runs in the background independently.
    go printMessage("Hello World")

    // 2. Run this normally in the main function
    printMessage("synchron")
}