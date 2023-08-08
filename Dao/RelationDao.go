package Dao

import "gorm.io/gorm"

type relation struct {
	gorm.Model
	UserIdOne int  `json:"user_id_one"`
	UserIdTwo int  `json:"user_id_two"`
	Forward   bool `json:"forward"` // 正向关系 UserOne->UserTwo
	Reverse   bool `json:"reverse"` // 反向关系 UserOne<-UserTwo
}
