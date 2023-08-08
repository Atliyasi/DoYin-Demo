package user

import (
	"go-crud-demo/Dao"
)

func QueryUserById(id int) (*Dao.VideoUser, error) {
	user, err := Dao.NewUserDao().FindUserById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
