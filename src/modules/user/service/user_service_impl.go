package service

import (
	"strconv"
	"syamsf/learning-gin/src/modules/user/repository"
	"syamsf/learning-gin/src/modules/user/request"
	"syamsf/learning-gin/src/modules/user/model"
	"github.com/go-playground/validator/v10"
	"github.com/gin-gonic/gin"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository 
	Validate 			 *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository, 
		Validate: validate,
	}
}

func (service *UserServiceImpl) FindAll() []model.User {
	return service.UserRepository.FindAll()
}

func (service *UserServiceImpl) FindById(id int) model.User {
	return service.UserRepository.FindById(id)
}

func (service *UserServiceImpl) Create(ctx *gin.Context) (*model.User, error) {
	var input request.CreateUserInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		return nil, err
	}

	err := service.Validate.Struct(input)

	if err != nil {
		return nil, err
	}

	user := model.User {
		Name: input.Name,
		Email: input.Email,
	}

	result, err := service.UserRepository.Save(user)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *UserServiceImpl) Update(ctx *gin.Context) (*model.User, error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var input request.UpdateUserInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		return nil, err
	}

	err := service.Validate.Struct(input)

	if err != nil {
		return nil, err
	}

	user := model.User {
		ID: int64(id),
		Name: input.Name,
		Email: input.Email,
	}

	result, err := service.UserRepository.Update(user)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *UserServiceImpl) Delete(ctx *gin.Context) (*model.User, error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	user := model.User{
		ID: int64(id),
	}

	result, err := service.UserRepository.Delete(user)
	if err != nil {
		return nil, err
	}

	return result, nil
}