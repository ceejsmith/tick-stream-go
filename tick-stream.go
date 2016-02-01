package main

import (
    "fmt"

    "github.com/ceejsmith/tick-stream-go/parse"
)

func main() {
    fmt.Println("This is a microservice to stream financial price ticks.")

    for _, tick := range parse.Parse() {
        fmt.Println(tick.Volume)
    }
}
