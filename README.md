# Time Translator
Command line time parser to convert from/to different date formats and timezones specified with the Go "reference time"

# Usage
echo "1997-07-16" | tt -from "2006-01-02" -to "Mon, Jan 1, 2006" && echo

# Reference Time
See https://golang.org/src/time/format.go
