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
		Status:  int(task3.Status_PENDING),
	}

	if err = mysql.CreateToDoList(newToDo); err != nil {
		return err
	}

	return nil
}

func (s *ToDoListService) QueryBatchToDoList(currentUserID int64, req *task3.QueryBatchToDoListsRequest) ([]*model.ToDoList, int64, error) {
	var err error

	todo_lists, total, err := mysql.QueryBatchToDoList(req.Keyword, (*int64)(req.Status), req.Page, req.PageSize)

	if err != nil {
		return nil, 0, err
	}

	return todo_lists, total, nil
}

func (s *ToDoListService) UpdateToDoList(currentUserID int64, req *task3.UpdateToDoListRequest) error {
	var err error

	existingToDo, err := mysql.QueryToDoListByID(req.TodoListID, currentUserID)

	if err != nil {
		return err
	}

	if req.Title != nil {
		existingToDo.Title = *req.Title
	}

	if req.Status != nil {
		existingToDo.Status = int(*req.Status)
	}

	if req.Context != nil {
		existingToDo.Context = *req.Context
	}

	if err = mysql.UpdateToDoListByID(existingToDo); err != nil {
		return err
	}

	return nil
}

func (s *ToDoListService) UpdateBatchStatus(currentUserID int64, req *task3.UpdateBatchStatusRequest) error {
	var err error

	err = mysql.UpdateBatchStatus(currentUserID, int(req.Status))

	if err != nil {
		return err
	}

	return nil
}

func (s *ToDoListService) DeleteToDoList(currentUserID int64, req *task3.DeleteToDoListRequest) error {
	var err error

	if err = mysql.DeleteToDoListByID(req.TodoListID, currentUserID); err != nil {
		return err
	}

	return nil
}

func (s *ToDoListService) DeletePendingToDos(currentUserID int64, req *task3.DeleteToDoListRequest) error {
	var err error

	if err = mysql.DeletePendingToDos(currentUserID); err != nil {
		return err
	}

	return nil
}

func (s *ToDoListService) DeleteCompletedToDos(currentUserID int64, req *task3.DeleteToDoListRequest) error {
	var err error

	if err = mysql.DeletePendingToDos(currentUserID); err != nil {
		return err
	}

	return nil
}

func (s *ToDoListService) DeleteAllToDos(currentUserID int64, req *task3.DeleteToDoListRequest) error {
	var err error

	if err = mysql.DeletePendingToDos(currentUserID); err != nil {
		return err
	}

	return nil
}
