package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"time"
)

var from = flag.String("from", "", "Time layout to convert from, specified in reference time - see https://golang.org/src/time/format.go#L61")
var to = flag.String("to", "", "Time layout to convert to, specified in reference time - see https://golang.org/src/time/format.go#L61")
var convertTime = flag.Bool("convert", false, "Convert all times to timezone (default UTC)")
var totimezone = flag.String("totimezone", "UTC", "Timezone to convert date to with -convert")
var fromtimezone = flag.String("fromtimezone", "UTC", "Timezone to parse date from (default UTC)")

var helpURL = "https://golang.org/src/time/format.go#L64"

// takes file from stdin and outputs to stdout in the correct time format, one date per line
func main() {
	flag.Parse()
	if *from == "" {
		log.Printf("No -from time layout specified - see %s", helpURL)
		flag.Usage()
		return
	}
	if *to == "" {
		log.Printf("No -to time layout specified - see %s", helpURL)
		flag.Usage()
		return
	}
	lineReader := bufio.NewScanner(bufio.NewReader(os.Stdin))
	lineWriter := bufio.NewWriter(os.Stdout)
	defer lineWriter.Flush()
	var totz *time.Location
	var fromtz *time.Location
	var err error
	if *convertTime {
		totz, err = time.LoadLocation(*totimezone)
		if err != nil {
			log.Fatalf("Error loading timezone - %s - %v", *totimezone, err)
		}
	}
	fromtz, err = time.LoadLocation(*fromtimezone)
	if err != nil {
		log.Fatalf("Error loading timezone - %s - %v", *fromtimezone, err)
	}
	for {
		if !lineReader.Scan() {
			break
		}
		line := lineReader.Text()
		date, err := time.ParseInLocation(*from, line, fromtz)
		if err != nil {
			log.Fatalf("Error parsing input time - %v", err)
		}
		if *convertTime {
			date = date.In(totz)
		}
		_, err = lineWriter.WriteString(date.Format(*to))
		if err != nil {
			log.Fatalf("Error writing output time - %v", err)
		}
		lineWriter.WriteString("\n")
	}
}
