package todo

import (
	"time"

	"github.com/kmhalpin/todoapp/internal/model"
)

type (
	Todo struct {
		ID     string
		Title  string
		Note   string
		Time   time.Time
		UserID string

		Timestamp
	}

	Timestamp = model.Timestamp
)
