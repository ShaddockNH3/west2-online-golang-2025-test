package pack

import (
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task3/biz/model"
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task3/biz/model/task3"
)

func Users(models []*model.User) []*task3.User {
	users := make([]*task3.User, 0, len(models))
	for _, m := range models {
		if u := User(m); u != nil {
			users = append(users, u)
		}
	}
	return users
}

func User(model *model.User) *task3.User {
	if model == nil {
		return nil
	}
	return &task3.User{
		UserID:    int64(model.ID),
		Name:      model.Name,
		Introduce: model.Introduce,
	}
}
