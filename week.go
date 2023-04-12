package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

const day = 24 * time.Hour

func main() {
	short := flag.Bool("s", false, "Non-verbose output.")
	flag.Parse()

	now := time.Now()
	y := now.Year()
	date0 := time.Date(y, 1, 1, 1, 0, 0, 0, time.Local)
	i := 0
	for date0.Weekday() == time.Saturday || date0.Weekday() == time.Sunday {
		date0 = date0.Add(day)
		i++
	}
	for ; i < 365; i++ {
		d := date0.Add(day * time.Duration(i))
		weekday := i/7 + 1
		if now.Before(d) {
			if *short {
				fmt.Println(weekday)
			} else {
				fmt.Printf("day %v of week %v (%v)\n", i+1, weekday, d.Weekday())
			}

			return
		}
	}
	fmt.Println("failed with bad day of year")
	os.Exit(1)
}
