package Dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var _db *gorm.DB

func InitDB() {
	dsn := "root:726400sb@tcp(127.0.0.1:3306)/bytedemo?charset=utf8mb3&parseTime=True&loc=Local"
	var err error
	_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println("数据库连接有误")
	}
	sqlDB, err := _db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	err = _db.AutoMigrate(&User{}, &VideoUser{}, &Video{}, &FavoriteList{}, &Comment{}, &Relation{}, &Message{})
	if err != nil {
		fmt.Println("AutoMigrate, err: ", err)
	}
}

func GetDB() *gorm.DB {
	return _db
}
