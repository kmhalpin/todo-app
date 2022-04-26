package model

import "time"

type (
	Timestamp struct {
		TimeCreated time.Time
		TimeUpdated time.Time
	}
)
