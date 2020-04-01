package utils

import (
	"time"
)

func GetYYYYMMDD() string {
	return time.Now().Format("2006-01-02")
}

func GetYYYYMMDDHHMMSS() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func UtcToTime(utcDate string) time.Time {
	date, _ := time.ParseInLocation("2006-01-02T15:04:05.000Z", utcDate, time.UTC) // "2018-10-21T01:20:18.0Z"
	return date
}

func TimeToStr(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
