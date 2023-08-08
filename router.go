package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-crud-demo/Controller"
	middleware "go-crud-demo/Middleware"
)

func InitRouter() {
	r := gin.Default()
	r.Static("/public/mov/", "./public/mov")
	r.Static("/public/pic/", "./public/pic")
	douYin := r.Group("/douyin")
	{
		douYin.POST("/user/register/", Controller.Register)
		douYin.POST("/user/login/", Controller.Login)
		douYin.GET("/user/", middleware.AuthMiddleWare(), Controller.UserInfo)
	}
	{
		douYin.GET("/feed/", Controller.Feed)
	}
	{
		douYin.POST("/publish/action/", middleware.AuthMiddleWare(), Controller.PublishAction)
		douYin.GET("/publish/list/", middleware.AuthMiddleWare(), Controller.PublishList)
	}
	{
		douYin.POST("/favorite/action/", middleware.AuthMiddleWare(), Controller.Upvote)
		douYin.GET("/favorite/list/", middleware.AuthMiddleWare(), Controller.FavoriteList)
	}
	{
		douYin.POST("/comment/action/", middleware.AuthMiddleWare(), Controller.HandleComment)
		douYin.GET("/comment/list/", Controller.GetCommentList)
	}
	{
		douYin.POST("/relation/action/", middleware.AuthMiddleWare())
		douYin.GET("/relation/follow/list/", middleware.AuthMiddleWare())
		douYin.GET("/relation/follower/list/", middleware.AuthMiddleWare())
	}
	PORT := "8080"
	//启动服务
	err := r.Run(":" + PORT)
	if err != nil {
		fmt.Println("Run, err: ", err)
	}
}
