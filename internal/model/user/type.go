package user

import "github.com/kmhalpin/todoapp/internal/model"

type (
	User struct {
		ID         string
		Username   string
		Password   string
		Permission permissionType

		Timestamp
	}

	permissionType string

	Timestamp = model.Timestamp
)
