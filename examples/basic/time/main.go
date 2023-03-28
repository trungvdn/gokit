package main

import (
	"fmt"
	"time"
)

func main() {
	r := RecurringTime{
		startTime: time.Date(2022, time.June, 1, 7, 0, 0, 0, time.UTC),
		endTime:   time.Date(2022, time.June, 1, 9, 0, 0, 0, time.UTC),
		endDate:   time.Now(),
		freq:      "Weekly",
	}
	chanTime := r.calculate()
	fmt.Println("chainTime", chanTime)
	// fmt.Println(chanTime[1:])

	for _, v := range chanTime[1:] {
		fmt.Println(v.startTime)
	}
	// result := make(map[string]*string)
	// v, ok := result["s"]
	// fmt.Println(v, ok)
}

type YMD struct {
	year  int
	month int
	days  int
}

var Frequencies = map[string]YMD{
	"Daily":   {days: 1},
	"Weekly":  {days: 7},
	"Monthly": {month: 1},
	"Yearly":  {year: 1},
}

type Frequency string

type RecurringTime struct {
	startTime time.Time
	endTime   time.Time
	endDate   time.Time
	freq      Frequency
}

type TimePare struct {
	startTime time.Time
	endTime   time.Time
}

func (r RecurringTime) calculate() []TimePare {
	fmt.Println("[debug]", r.startTime, r.endDate)
	timePare := []TimePare{}
	for r.startTime.Before(r.endDate) {
		fmt.Println("[debug]", r.startTime, r.endTime)
		freq := Frequencies[string(r.freq)]
		timePare = append(timePare, TimePare{
			startTime: r.startTime,
			endTime:   r.endDate,
		})
		r.startTime = r.startTime.AddDate(freq.year, freq.month, freq.days)
		r.endTime = r.endTime.AddDate(freq.year, freq.month, freq.days)
	}
	return timePare
}
