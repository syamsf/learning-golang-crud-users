package repository

import "syamsf/learning-gin/src/modules/user/model"

type UserRepository interface {
	FindAll() []model.User
	FindById(id int) model.User
	Save(user model.User) (*model.User, error)
	Update(user model.User) (*model.User, error)
	Delete(user model.User) (*model.User, error)
}
