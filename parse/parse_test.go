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

func TestParseDate(t *testing.T) {
    time := parseDate("20160127", "093101")

    checkDate([]int { 2016, 1, 27, 9, 31, 1 }, time, t) 
}

func TestParseLine(t *testing.T) {
    tick := parseLine("TICKER,0,20150512,211331,21.123000,34")

    if tick.Ticker != "TICKER" { t.Errorf("Expected ticker as TICKER. Got: %s", tick.Ticker) } 

    checkDate([]int { 2015, 5, 12, 21, 13, 31 }, tick.Time, t)

    if tick.LastPrice != 21.123 { t.Errorf("Expected last price as 21.123. Got: %f", tick.LastPrice) }
    if tick.Volume != 34 { t.Errorf("Expected volume as 34. Got %d", tick.Volume) }
}
