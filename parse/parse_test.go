package parse

import "testing"

func TestParseDate(t *testing.T) {
    time := parseDate("20160127", "093101")

    year, month, day := time.Date()
    if year != 2016 { t.Errorf("Expected year 2016. Got: %d", year) }
    if month != 1 { t.Errorf("Expected month 1. Got: %d", month) }
    if day != 27 { t.Errorf("Expected day 27. Got: %d", day) }

    hour, min, sec := time.Clock()

    if hour != 9 { t.Errorf("Expected hour 9. Got: %d", hour) }
    if min != 31 { t.Errorf("Expected minute 31. Got: %d", min) }
    if sec != 1 { t.Errorf("Expected second 1. Got: %d", sec) }
}
