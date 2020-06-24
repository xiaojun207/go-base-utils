package utils

import "time"

func NewFixedDelayJob(d time.Duration, f func()) {
	go func() {
		for {
			f()
			time.Sleep(d)
		}
	}()
}
