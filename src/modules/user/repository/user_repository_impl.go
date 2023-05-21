package repository

import (
	"syamsf/learning-gin/src/modules/user/model"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db}
}

func (repository *UserRepositoryImpl) FindAll() []model.User {
	var users []model.User

	_ = repository.db.Find(&users)

	return users
}

func (repository *UserRepositoryImpl) FindById(id int) model.User {
	var user model.User

	_ = repository.db.First(&user, id)
	
	return user
}

func (repository *UserRepositoryImpl) Save(user model.User) (*model.User, error) {
	result := repository.db.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (repository *UserRepositoryImpl) Update(user model.User) (*model.User, error) {
	result := repository.db.Save(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (repository *UserRepositoryImpl) Delete(user model.User) (*model.User, error) {
	result := repository.db.Delete(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
	