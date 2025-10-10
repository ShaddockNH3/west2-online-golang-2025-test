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

func DeleteCompletedToDos(userID int64) error {
	return DB.Where("user_id = ? AND status = ?", userID, task3.Status_COMPLETED).Delete(&model.ToDoList{}).Error
}

func DeletePendingToDos(userID int64) error {
	return DB.Where("user_id = ? AND status = ?", userID, task3.Status_PENDING).Delete(&model.ToDoList{}).Error
}

func DeleteAllToDos(userID int64) error {
	return DB.Where("user_id = ?", userID).Delete(&model.ToDoList{}).Error
}

func UpdateToDoListByID(TodoListID int64, todo_list *model.ToDoList) error {
	return DB.Where("id=?", TodoListID).Updates(todo_list).Error
}

func UpdateBatchStatus(userID int64, status int) error {
	return DB.Model(&model.ToDoList{}).Where("user_id = ?", userID).Update("status", status).Error
}

func QueryToDoListByID(todoListID int64, userID int64) (*model.ToDoList, error) {
	var todoList model.ToDoList
	if err := DB.Where("id = ? AND user_id = ?", todoListID, userID).First(&todoList).Error; err != nil {
		return nil, err
	}
	return &todoList, nil
}

func QueryBatchToDoList(currentUserID int64, keyword *string, status *int64, page, pageSize int64) ([]*model.ToDoList, int64, error) {
	db := DB.Model(model.ToDoList{})

	db = db.Where("user_id = ?", currentUserID).First(db)

	if keyword != nil && len(*keyword) != 0 {
		db = db.Where(DB.Or("name like ?", "%"+*keyword+"%").
			Or("introduce like ?", "%"+*keyword+"%"))
	}

	if status != nil {
		db = db.Where("status = ?", status).First(db)
	}

	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var res []*model.ToDoList
	if err := db.Limit(int(pageSize)).Offset(int(pageSize * (page - 1))).Find(&res).Error; err != nil {
		return nil, 0, err
	}
	return res, total, nil
}
