package model

import (
	"errors"
	"fmt"
)

type UserMockImpl struct {
	users   map[uint]User
	counter uint
}

func NewUserMockImpl(initCount int) *UserMockImpl {
	stub := new(UserMockImpl)
	stub.users = make(map[uint]User)
	for i := 1; i <= initCount; i++ {
		user := User{Username: fmt.Sprintf("%d", i)}
		user.ID = uint(i)
		stub.users[uint(i)] = user
		stub.counter++
	}
	return stub
}

func (ur *UserMockImpl) GetUsers(page, limit int) ([]User, error) {
	var users []User
	for _, v := range ur.users {
		users = append(users, v)
	}
	return users, nil
}

func (ur *UserMockImpl) UpdateUser(u *User) error {
	if _, ok := ur.users[u.ID]; ok {
		ur.users[u.ID] = *u
		return nil
	}
	return errors.New("user not found")
}

func (ur *UserMockImpl) SaveUser(u *User) error {
	ur.counter++
	u.ID = ur.counter
	ur.users[u.ID] = *u
	return nil
}

func (ur *UserMockImpl) GetUser(id uint) (*User, error) {
	if user, ok := ur.users[id]; ok {
		return &user, nil
	}
	return nil, errors.New("User not found")
}

func (ur *UserMockImpl) DeleteUser(id uint) error {
	if _, ok := ur.users[id]; ok {
		delete(ur.users, id)
		return nil
	}
	return errors.New("User not found")
}
