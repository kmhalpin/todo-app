package http

import (
	httpCommon "github.com/kmhalpin/todoapp/common/http"
	tModel "github.com/kmhalpin/todoapp/internal/model/todo"
)

func (h HTTPTodoDelivery) mapTodoModelToResponse(t tModel.Todo) httpCommon.Todo {
	return httpCommon.Todo{
		ID:          t.ID,
		Title:       t.Title,
		Note:        t.Note,
		Time:        t.Time,
		UserID:      t.UserID,
		TimeCreated: t.TimeCreated,
		TimeUpdated: t.TimeUpdated,
	}
}

func (h HTTPTodoDelivery) mapTodoBodyToModel(t httpCommon.Todo) tModel.Todo {
	return tModel.Todo{
		ID:     t.ID,
		Title:  t.Title,
		Note:   t.Note,
		Time:   t.Time,
		UserID: t.UserID,

		Timestamp: tModel.Timestamp{
			TimeCreated: t.TimeCreated,
			TimeUpdated: t.TimeUpdated,
		},
	}
}
