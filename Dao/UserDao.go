package Dao

import (
	"gorm.io/gorm"
	"sync"
)

type User struct {
	gorm.Model
	Username  string `gorm:"column:username" gorm:"uniqueIndex"`
	Password  string `gorm:"column:password"`
	IsDeleted bool   `gorm:"column:is_deleted"`
}

func (user *User) SetUser(username string, password string) {
	user.Username = username
	user.Password = password
}

type UserDao struct{}

var userDao *UserDao
var userOnce sync.Once

func NewUserDao() *UserDao {
	userOnce.Do(func() {
		userDao = &UserDao{}
	})
	return userDao
}

func (*UserDao) FindUserById(id int) (*VideoUser, error) {
	var user VideoUser
	if err := GetDB().Where("id=?", id).Find(&user).Error; err != nil {
		return nil, err
	} else {
		return &user, err
	}
}

func (*UserDao) FindUserByName(username string) (*User, error) {
	var user User
	if err := GetDB().Where("username=?", username).Find(&user).Error; err != nil {
		return nil, err
	} else {
		return &user, err
	}
}

func (*UserDao) RegisterUser(user *User) (int, error) {
	if err := GetDB().Create(user).Error; err != nil {
		return int(user.ID), err
	}
	return int(user.ID), nil
}
