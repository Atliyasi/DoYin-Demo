package user

import (
	"errors"
	"go-crud-demo/Dao"
	"go-crud-demo/Util"
	"regexp"
)

type UserRegisterFlow struct {
	username string
	password string
}

func NewUserRegisterFlow(username string, password string) *UserRegisterFlow {
	return &UserRegisterFlow{
		username: username,
		password: password,
	}
}
func Register(username string, password string) (int, error) {
	return NewUserRegisterFlow(username, password).Do()
}
func (u *UserRegisterFlow) Do() (int, error) {
	if err := u.CheckParam(); err != nil {
		return 0, err
	}
	id, err := u.Register()
	if err != nil {
		return 0, err
	}
	return id, err
}

func (u *UserRegisterFlow) CheckParam() error {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	if reg.MatchString(u.username) {
		return nil
	}
	return errors.New("请输入正确的邮箱")
}

func (u *UserRegisterFlow) Register() (int, error) {
	//用户密码加密
	PasswordHash, err := Util.PwdHash(u.password)
	if err != nil {
		return 0, err
	}
	//用户是否存在
	userquery, err := Dao.NewUserDao().FindUserByName(u.username)
	if err != nil {
		return 0, err
	} else if err == nil && userquery.ID != 0 {
		return 0, errors.New("用户已经存在")
	}
	//创建用户
	userinfo := &Dao.User{
		Username:  u.username,
		Password:  PasswordHash,
		IsDeleted: false,
	}
	userId, err := Dao.NewUserDao().RegisterUser(userinfo)
	if err != nil {
		return 0, err
	}
	return userId, nil
}
