package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Printf("\nCurrent time: %s\n", currentTime)
	fmt.Printf("\nCurrent time with format: %s\n", currentTime.Format(time.RFC1123))

	formattedTime := currentTime.Format("2006-01-02 15:04:05 Monday")
	fmt.Printf("\nFormatted time: %s\n", formattedTime)

	// Add 1 more day to the current time
	futureTime := currentTime.Add(24 * time.Hour)
	formatFutureTime := futureTime.Format("2006-01-02 15:04:05 Monday")
	fmt.Printf("\nFuture time (after adding 1 day) with format: %s\n", formatFutureTime)
}