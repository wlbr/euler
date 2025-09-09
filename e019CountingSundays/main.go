/*
Counting Sundays
Problem 19

You are given the following information, but you may prefer to do some research for yourself.

  - 1 Jan 1900 was a Monday.
  - Thirty days has September,
    April, June and November.
    All the rest have thirty-one,
    Saving February alone,
    Which has twenty-eight, rain or shine.
    And on leap years, twenty-nine.
  - A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.

How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?

Answer:  171

*/

package main

import (
	"time"
)

func CheckForSunday(year int, month int, day int) bool {
	t := time.Date(year, time.Month(month), day, 12, 0, 0, 0, time.UTC)
	return t.Weekday() == time.Sunday
}

func CheckForSundayTime(t time.Time) bool {
	return t.Weekday() == time.Sunday
}

func searchFirstSunday(start time.Time) time.Time {
	res := start
	for i := 0; i < 7; i++ {
		current := start.AddDate(0, 0, i)
		if CheckForSundayTime(current) {
			res = current
		}
	}
	return res
}

func main() {
	start := time.Date(1901, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2000, 12, 31, 0, 0, 0, 0, time.UTC)

	current := searchFirstSunday(start)
	result := 0
	if CheckForSundayTime(current) && current.Day() == 1 {
		result++
	}

	for i := 0; ; i++ {
		current = current.AddDate(0, 0, 7)
		if current.After(end) {
			break
		}
		if CheckForSundayTime(current) && current.Day() == 1 {
			result++
		}
	}
	println(result)
}
