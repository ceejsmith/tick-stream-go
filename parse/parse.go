package parse

import (
    "strings"
    "strconv"
    "time"
)

// The header of the data files is:
// TICKER,PER,DATE,TIME,LAST,VOL
// A sample line is:
// AA,0,20160127,093133,7.030000000,400

type Tick struct {
    Ticker string
    Time time.Time
    LastPrice float64
    Volume int
}

func parseDate(datePart, timePart string) time.Time {
    // Ignoring error returns is bad practice, but OK for prototyping with clean data
    year, _ := strconv.Atoi(datePart[:4])
    month, _ := strconv.Atoi(datePart[4:6])
    day, _ := strconv.Atoi(datePart[6:])

    hour, _ := strconv.Atoi(timePart[:2])
    min, _ := strconv.Atoi(timePart[2:4])
    sec, _ := strconv.Atoi(timePart[4:])

    return time.Date(year, time.Month(month), day, hour, min, sec, 0, time.UTC)
}

func parseLine(line string) Tick {
    fields := strings.Split(line, ",")
    lastPrice, _ := strconv.ParseFloat(fields[4], 64)
    volume, _ := strconv.Atoi(fields[5])
    return Tick{
        fields[0],
        parseDate(fields[2], fields[3]),
        lastPrice,
        volume,
    }
}

func Parse(lines []string) []Tick {
    result := make([]Tick, len(lines))
    for i, line := range lines {
        result[i] = parseLine(line)
    }
    return result
}
