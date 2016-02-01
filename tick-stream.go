package main

import (
    "os"
    "fmt"
    "io/ioutil"
    "strings"

    "github.com/ceejsmith/tick-stream-go/parse"
)

func main() {
    filename := os.Args[1]

    file, err := ioutil.ReadFile(filename)

    if err != nil {
        panic(fmt.Sprintf("Cannot open %s", filename))
    }

    lines := strings.Split(string(file), "\r\n")

    fmt.Printf("Found %d lines\n", len(lines))

    // Discard header and empty trailing line
    for _, tick := range parse.Parse(lines[1:len(lines) - 1]) {
        fmt.Println(tick.Volume)
    }
}
