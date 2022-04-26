package http

import "time"

type (
	Error struct {
		Code    int               `json:"code"`
		Message string            `json:"message"`
		Errors  map[string]string `json:"errors"`
	}

	Todo struct {
		ID          string    `json:"id"`
		Title       string    `json:"title" binding:"required"`
		Note        string    `json:"note"`
		Time        time.Time `json:"time" binding:"required"`
		UserID      string    `json:"owner"`
		TimeCreated time.Time `json:"created_at"`
		TimeUpdated time.Time `json:"updated_at"`
	}

	User struct {
		ID          string    `json:"id"`
		Username    string    `json:"username" binding:"required"`
		Password    string    `json:"password" binding:"required"`
		Permission  string    `json:"permission"`
		TimeCreated time.Time `json:"created_at"`
		TimeUpdated time.Time `json:"updated_at"`
	}

	AddUser struct {
		ID          string    `json:"id"`
		Username    string    `json:"username" binding:"required"`
		Password    string    `json:"password" binding:"required"`
		Permission  string    `json:"permission" binding:"required,oneof=r r/w"`
		TimeCreated time.Time `json:"created_at"`
		TimeUpdated time.Time `json:"updated_at"`
	}
)
