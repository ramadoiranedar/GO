package main

import (
	"fmt"
	"time"
)

// TimeIn returns the time in UTC if the name is "" or "UTC".
// It returns the local time if the name is "Local".
// Otherwise, the name is taken to be a location name in
// the IANA Time Zone database, such as "Africa/Lagos".
func TimeIn(t time.Time, name string) (time.Time, error) {
	loc, err := time.LoadLocation(name)
	if err == nil {
		t = t.In(loc)
	}
	return t, err
}

func main() {
	now := time.Now()
	fmt.Println("LOCAL", now.Local())
	fmt.Println("DAY", now.Day())
	fmt.Println("MONTH", now.Month())
	fmt.Println("YEAR", now.Year())
	fmt.Println("HOUR", now.Hour())
	fmt.Println("MINUTE", now.Minute())
	fmt.Println("SECOND", now.Second())
	fmt.Println("NANOSECOND", now.Nanosecond())
	fmt.Println("UNIX SECOND", now.Unix())
	fmt.Println("UNIX MILIS", now.UnixMilli())

	manual := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("MANUAL", manual)

	layout := time.RFC3339
	parse, _ := time.Parse(layout, "2022-08-06T16:33:22Z")
	fmt.Println("PARSE", parse)

	// TIME ZONE
	for _, name := range []string{
		"",
		"Local",
		"Asia/Shanghai",
		"Asia/Jakarta",
	} {
		t, err := TimeIn(time.Now(), name)
		if err == nil {
			fmt.Println(t.Location(), t.Format("15:04"))
		} else {
			fmt.Println(name, "<time unknown>")
		}
	}
}
