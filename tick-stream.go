package main

import (
    "os"
    "fmt"
    "io/ioutil"
    "strings"

    "github.com/ceejsmith/tick-stream-go/parse"
    "github.com/ceejsmith/tick-stream-go/stream"
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
    volumes := make([]string, len(lines) - 2)
    for i, tick := range parse.Parse(lines[1:len(lines) - 1]) {
        volumes[i] = fmt.Sprintf("%d", tick.Volume)
    }

    stream.Strings(volumes)
}
