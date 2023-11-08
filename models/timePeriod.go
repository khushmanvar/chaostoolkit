package models

import "time"

type TimePeriod struct {
	From time.Time
	To   time.Time
}