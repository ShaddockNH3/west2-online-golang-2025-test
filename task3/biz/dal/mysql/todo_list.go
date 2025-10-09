package mysql

import (
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task3/biz/model"
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task3/biz/model/task3"
)

func CreateToDoList(todo_lists *model.ToDoList) error {
	return DB.Create(todo_lists).Error
}

func DeleteToDoListByID(todoListID int64, userID int64) error {
	return DB.Where("id = ? AND user_id=?", todoListID, userID).Delete(&model.ToDoList{}).Error
}

func DeleteToDoListsComplete(userID int64) error {
	return DB.Where("user_id = ? AND status = ?", userID, task3.Status_Complete).Delete(&model.ToDoList{}).Error
}

func DeleteAllUserToDoLists(userID int64) error {
	return DB.Where("user_id = ?", userID).Delete(&model.ToDoList{}).Error
}

func UpdateToDoListByID(todo_list *model.ToDoList) error {
	return DB.Updates(todo_list).Error
}

func UpdateUserToDoListsStatus(userID int64, status int) error {
	return DB.Model(&model.ToDoList{}).Where("user_id = ?", userID).Update("status", status).Error
}

func QueryToDoList() {

}

func QueryToDoListByID() {

}

func QueryToDoListByStatus() {

}
