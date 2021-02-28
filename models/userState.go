package models

import "home/repository"

type UserState struct {
	ID          int    `gorm:"primary_key" json:"id"`
	Name        string `gorm:"name" json:"name" binding:"required"`
	Description string `gorm:"description" json:"description"`
}

func (UserState) TableName() string {
	return "UserState"
}

func (state *UserState) CreateUserStateRecord(rep repository.Repository) error {

	result := repository.Repository.Create(rep, &state)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *UserState) FindByName(rep repository.Repository, name string) *UserState {
	var userState UserState
	rep.Find(&userState, "name", name)
	if userState.ID > 0 {
		return &userState
	}
	return nil
}

func (u *UserState) FindAll(rep repository.Repository) []UserState {
	var stateList []UserState

	tablename := u.TableName()

	rep.Table(tablename).Select("*").Scan(&stateList)
	return stateList
}
