package common

import "time"

func GetCurrentTime() string {
	currentTime := time.Now()
	dateLayout := "2006-01-02 15:04:05"
	return currentTime.Format(dateLayout)
}
