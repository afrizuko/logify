package model

import (
	"errors"

	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		Username  string
		Password  string
		Email     string
		Contact   string
		Firstname string
		Lastname  string
		Gender    string
	}

	UserRepository interface {
		GetUsers(page, limit int) ([]User, error)
		GetUser(id uint) (*User, error)
		SaveUser(u *User) error
		UpdateUser(u *User) error
		DeleteUser(id uint) error
	}
)

type UserRepositoryImpl struct {
	*gorm.DB
}

func NewUserRepository() *UserRepositoryImpl {
	stub := new(UserRepositoryImpl)
	stub.DB = GetConnection()
	stub.DB.AutoMigrate(&User{})
	return stub
}

func (ur *UserRepositoryImpl) GetUsers(page, limit int) ([]User, error) {
	var users []User
	err := ur.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *UserRepositoryImpl) UpdateUser(u *User) error {
	if u == nil {
		return errors.New("Invalid user reference")
	}
	if err := ur.DB.Save(u).Error; err != nil {
		return err
	}
	return nil
}

func (ur *UserRepositoryImpl) SaveUser(u *User) error {
	if u == nil {
		return errors.New("Invalid user reference")
	}
	if err := ur.DB.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func (ur *UserRepositoryImpl) GetUser(id uint) (*User, error) {
	var user User
	if err := ur.DB.First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepositoryImpl) DeleteUser(id uint) error {
	if err := ur.DB.Delete(&User{}, id).Error; err != nil {
		return err
	}
	return nil
}
