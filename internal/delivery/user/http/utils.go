package http

import (
	httpCommon "github.com/kmhalpin/todoapp/common/http"
	uModel "github.com/kmhalpin/todoapp/internal/model/user"
)

func (h HTTPUserDelivery) mapUserModelToResponse(u uModel.User) httpCommon.User {
	return httpCommon.User{
		ID:          u.ID,
		Username:    u.Username,
		Password:    u.Password,
		Permission:  u.GetUserPermissionString(),
		TimeCreated: u.TimeCreated,
		TimeUpdated: u.TimeUpdated,
	}
}

func (h HTTPUserDelivery) mapUserBodyToModel(u httpCommon.AddUser) uModel.User {
	return uModel.User{
		ID:       u.ID,
		Username: u.Username,
		Password: u.Password,

		Timestamp: uModel.Timestamp{
			TimeCreated: u.TimeCreated,
			TimeUpdated: u.TimeUpdated,
		},
	}.SetUserPermissionString(u.Permission)
}
