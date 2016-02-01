package main

import (
    "fmt"

    "github.com/ceejsmith/tick-stream-go/parse"
)

var lines = []string {
    "AA,0,20160127,093001,7.070000000,500",
    "AA,0,20160127,093009,7.060000000,100",
    "AA,0,20160127,093100,7.050000000,100",
    "AA,0,20160127,093100,7.050000000,100",
    "AA,0,20160127,093100,7.050000000,100",
    "AA,0,20160127,093100,7.040000000,100",
    "AA,0,20160127,093100,7.030000000,1000",
    "AA,0,20160127,093100,7.030000000,100",
    "AA,0,20160127,093101,7.050000000,100",
    "AA,0,20160127,093118,7.040000000,100",
}

func main() {
    fmt.Println("This is a microservice to stream financial price ticks.")

    for _, tick := range parse.Parse(lines) {
        fmt.Println(tick.Volume)
    }
}
