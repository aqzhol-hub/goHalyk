package service

import (
	"errors"
	"home/custom"
	"home/models"
)

type UserService struct {
	cs custom.Custom
}

func NewUserService(cs custom.Custom) *UserService {
	return &UserService{cs: cs}
}

func (service *UserService) FindByID(userID uint) (error, *models.User) {
	rep := service.cs.GetRepository()
	var user models.User

	result := user.FindByID(rep, userID)
	if result == nil {
		return errors.New("User not found"), nil
	}

	return nil, result
}

func (service *UserService) FindByUsername(username string) (error, *models.User) {
	rep := service.cs.GetRepository()
	var user models.User

	result := user.FindByUsername(rep, username)
	if result == nil {
		return errors.New("User not found"), nil
	}

	return nil, result
}

func (service *UserService) FindAll() (error, []models.User) {
	rep := service.cs.GetRepository()
	var user models.User

	result := user.FindAll(rep)
	if result == nil {
		return errors.New("Users not found"), nil
	}

	return nil, result
}
