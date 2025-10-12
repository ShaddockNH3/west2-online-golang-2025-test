package pack

import "github.com/ShaddockNH3/west2-online-golang-2025-test/task4/biz/dal/db"

func Users(models []*db.User) []*db.User {
	users := make([]*db.User, 0, len(models))
	for _, m := range models {
		if u := User(m); u != nil {
			users = append(users, u)
		}
	}
	return users
}

func User(model *db.User) *db.User {
	if model == nil {
		return nil
	}
	return &db.User{
		ID:        model.ID,
		Username:  model.Username,
		Password:  model.Password,
		AvatarUrl: model.AvatarUrl,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		DeletedAt: model.DeletedAt,
	}
}
