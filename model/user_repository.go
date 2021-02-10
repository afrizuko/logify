package model

import "gorm.io/gorm"

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

	return nil, nil
}

func (ur *UserRepositoryImpl) UpdateUser(u *User) error {

	return nil
}

func (ur *UserRepositoryImpl) SaveUser(u *User) error {
	return nil
}

func (ur *UserRepositoryImpl) GetUser(id uint) (*User, error) {
	return nil, nil
}

func (ur *UserRepositoryImpl) DeleteUser(id uint) error {
	return nil
}
