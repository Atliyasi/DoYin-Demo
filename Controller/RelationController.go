package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-crud-demo/Dao"
	"go-crud-demo/Service/relation"
	"net/http"
	"strconv"
)

type RelationResponse struct {
	Response
	UserList []Dao.VideoUser `json:"user_list"`
}

func Relation(c *gin.Context) {
	toUserIdS := c.Query("to_user_id")
	userIdA, _ := c.Get("uid")
	actionType := c.Query("action_type")
	toUserId, _ := strconv.Atoi(toUserIdS)
	userId := userIdA.(int)
	if actionType == "1" {
		if err := relation.NewRelation(userId, toUserId).Relation(); err != nil {
			fmt.Println("relation.NewRelation(userId, toUserId).Relation()")
			c.JSON(http.StatusBadRequest, Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			})
		}
	} else if actionType == "2" {
		if err := relation.NewRelation(userId, toUserId).CancelRelation(); err != nil {
			c.JSON(http.StatusBadRequest, Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			})
		}
	} else {
		fmt.Println("else")
		c.JSON(http.StatusBadRequest, Response{
			StatusCode: 1,
			StatusMsg:  "失败",
		})
	}
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  "关注成功",
	})
}

func Follow(c *gin.Context) {
	userIdS := c.Query("user_id")
	userId, _ := strconv.Atoi(userIdS)
	videoUserList, err := relation.FollowList(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, RelationResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
			UserList: nil,
		})
	}
	c.JSON(http.StatusOK, RelationResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "成功",
		},
		UserList: videoUserList,
	})
}

func Follower(c *gin.Context) {
	userIdS := c.Query("user_id")
	userId, _ := strconv.Atoi(userIdS)
	videoUserList, err := relation.FollowerList(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, RelationResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
			UserList: nil,
		})
	}
	c.JSON(http.StatusOK, RelationResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "成功",
		},
		UserList: videoUserList,
	})
}
