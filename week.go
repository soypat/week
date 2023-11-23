package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	short := flag.Bool("s", false, "Non-verbose output.")
	flag.Parse()
	now := time.Now()
	dayOfYear := now.YearDay()
	offset := startWeekendDaysOfYear(now.Year())
	dayOfYear -= offset
	weeks := (dayOfYear / 7) + 1
	if *short {
		fmt.Println(weeks)
	} else {
		fmt.Printf("day %v of week %v (%v)\n", dayOfYear, weeks, now.Weekday())
	}
}

func startWeekendDaysOfYear(y int) int {
	day0 := time.Date(y, 1, 1, 1, 0, 0, 0, time.Local)
	// Return number of weekend days at start of year.
	switch day0.Weekday() {
	case time.Saturday:
		return 2
	case time.Sunday:
		return 1
	default:
		return 0
	}
}
