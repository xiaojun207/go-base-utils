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
