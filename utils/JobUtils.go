package utils

import "time"

func NewFixdDelayJob(d time.Duration, f func()) {
	go func() {
		for {
			f()
			time.Sleep(d)
		}
	}()
}
