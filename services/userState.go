package service

import (
	"errors"
	"home/custom"
	"home/models"
)

type UserStateService struct {
	cs custom.Custom
}

func NewUserStateService(cs custom.Custom) *UserStateService {
	return &UserStateService{cs: cs}
}

func (service *UserStateService) FindByName(name string) (error, *models.UserState) {
	rep := service.cs.GetRepository()
	var userState models.UserState

	result := userState.FindByName(rep, name)
	if result == nil {
		return errors.New("userState not found"), nil
	}

	return nil, result
}

func (service *UserStateService) FindAll() (error, []models.UserState) {
	rep := service.cs.GetRepository()
	var userState models.UserState

	result := userState.FindAll(rep)
	if len(result) == 0 {
		return errors.New("userState not found"), nil
	}
	return nil, result
}
