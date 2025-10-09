package service

import (
	"context"

	"github.com/ShaddockNH3/west2-online-golang-2025-test/task3/biz/dal/mysql"
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task3/biz/model"
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task3/biz/model/task3"
)

type ToDoListService struct {
	ctx context.Context
}

func NewToDoListService(ctx context.Context) *ToDoListService {
	return &ToDoListService{ctx: ctx}
}

func (s *ToDoListService) CreateToDoList(currentUserID int64, req *task3.CreateToDoListRequest) error {
	var err error

	newToDo := &model.ToDoList{
		UserID:  currentUserID,
		Title:   req.Title,
		Context: req.Context,
		Status:  int(task3.Status_ToDo),
	}

	if err = mysql.CreateToDoList(newToDo); err != nil {
		return err
	}

	return nil
}
