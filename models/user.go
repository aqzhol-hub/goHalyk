package models

import (
	"errors"
	"fmt"
	"home/repository"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int        `gorm:"primary_key" json:"id"`
	FirstName string     `gorm:"firstname" json:"firstname"`
	LastName  string     `gorm:"lastname" json:"lastname"`
	Birth     string     `gorm:"birth" json:"birth"`
	Username  string     `gorm:"username,unique" json:"username" form:"username" binding:"required"`
	Password  string     `gorm:"password" json:"password" form:"password" binding:"required"`
	StateID   uint       `gorm:"state_id" json:"state_id"`
	State     *UserState `gorm:"state" json:"state"`
}

func (User) TableName() string {
	return "User"
}

func (user *User) CreateUserRecord(rep repository.Repository) error {

	result := repository.Repository.Create(rep, &user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (user *User) HashPassword() error {
	password := user.Password
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	user.Password = string(bytes)

	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}

	return nil
}

func (u *User) FindByUsername(rep repository.Repository, username string) *User {
	var users User
	rep.Find(&users, "username", username)
	if users.ID > 0 {
		return &users
	}
	return nil
}

func (u *User) FindByID(rep repository.Repository, ID uint) *User {
	var users User
	rep.Find(&users, ID)
	if users.ID > 0 {
		return &users
	}
	return nil
}

func (u *User) FindAll(rep repository.Repository) []User {
	var userList []User
	tablename := u.TableName()
	rep.Table(tablename).Select("*").Scan(&userList)
	return userList
}

func (u *User) UpdateState(rep repository.Repository, stateID uint) error {
	db := rep.Model(&User{}).Where("ID = ?", u.ID).Update("StateID", stateID)
	fmt.Println(db)
	if db == nil {
		return errors.New("can not update state")
	}
	return nil
}
