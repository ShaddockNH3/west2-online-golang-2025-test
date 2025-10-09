package pack

import (
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task3/biz/model"
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task3/biz/model/task3"
)

func ToDoLists(models []*model.ToDoList) []*task3.ToDoList {
	todo_lists := make([]*task3.ToDoList, 0, len(models))
	for _, m := range models {
		if u := ToDoList(m); u != nil {
			todo_lists = append(todo_lists, u)
		}
	}
	return todo_lists
}

func ToDoList(model *model.ToDoList) *task3.ToDoList {
	if model == nil {
		return nil
	}
	return &task3.ToDoList{
		TodoListID: int64(model.ID),
		UserID:     model.UserID,
		Title:      model.Title,
		Context:    model.Context,
		Status:     task3.Status(model.Status),
	}
}
