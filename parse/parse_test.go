package parse

import (
    "testing"
    "time"
)

func checkDate(expected []int, actual time.Time, t *testing.T) {
    year, month, day := actual.Date()
    if year != expected[0] { t.Errorf("Expected year %d. Got: %d", expected[0], year) }
    if int(month) != expected[1] { t.Errorf("Expected month %d. Got: %d", expected[1], year) }
    if day != expected[2] { t.Errorf("Expected day %d. Got: %d", expected[2], day) }

    hour, min, sec := actual.Clock()

    if hour != expected[3] { t.Errorf("Expected hour %d. Got: %d", expected[3], hour) }
    if min != expected[4] { t.Errorf("Expected minute %d. Got: %d", expected[4], min) }
    if sec != expected[5] { t.Errorf("Expected second %d. Got: %d", expected[5], sec) }
}

func checkLine(expectedTicker string, expectedDate []int,  expectedLastPrice float64, expectedVolume int, tick Tick, t *testing.T) {
    if tick.Ticker != expectedTicker  { t.Errorf("Expected ticker as %s. Got: %s", expectedTicker, tick.Ticker) } 

    checkDate(expectedDate, tick.Time, t)
    
    if tick.LastPrice != expectedLastPrice { t.Errorf("Expected last price as %f. Got: %f", expectedLastPrice, tick.LastPrice) }
    if tick.Volume != expectedVolume { t.Errorf("Expected volume as %d. Got %d", expectedVolume, tick.Volume) }
}

func TestParseDate(t *testing.T) {
    time := parseDate("20160127", "093101")

    checkDate([]int { 2016, 1, 27, 9, 31, 1 }, time, t) 
}

func TestParseLine(t *testing.T) {
    tick := parseLine("TICKER,0,20150512,211331,21.123000,34")

    checkLine("TICKER", []int { 2015, 5, 12, 21, 13, 31 }, 21.123, 34, tick, t) 
}

/*
var lines = []string {
    "AA,0,20160127,093001,7.070000000,500",
    "AA,0,20160127,093009,7.060000000,100",
    "AA,0,20160127,093100,7.050000000,100",
}*/
