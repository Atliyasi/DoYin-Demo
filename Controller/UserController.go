package Controller

import (
	"github.com/gin-gonic/gin"
	"go-crud-demo/Dao"
	middleware "go-crud-demo/Middleware"
	"go-crud-demo/Service/user"
	"net/http"
	"strconv"
)

type UserResponse struct {
	Response
	UserId int    `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserInfoResponse struct {
	Response
	User *Dao.VideoUser `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	userId, err := user.Register(username, password)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
	}
	token, err := middleware.GenToken(userId)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
	}
	c.JSON(http.StatusOK, UserResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "注册成功",
		},
		UserId: userId,
		Token:  token,
	})
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	userId, err := user.Login(username, password)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
	}
	token, err := middleware.GenToken(userId)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})

	}
	c.JSON(http.StatusOK, UserResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "登录成功",
		},
		UserId: userId,
		Token:  token,
	})
}

func UserInfo(c *gin.Context) {
	userIdString := c.Query("user_id")
	userid, err := strconv.Atoi(userIdString)
	if err != nil {
		c.JSON(http.StatusOK, UserInfoResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
	}
	userinfo, err := user.QueryUserById(userid)
	if err != nil {
		c.JSON(http.StatusOK, UserInfoResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
	}
	c.JSON(http.StatusOK, UserInfoResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "信息返回成功",
		},
		User: userinfo,
	})
}
