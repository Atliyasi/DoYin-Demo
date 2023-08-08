package Controller

import (
	"github.com/gin-gonic/gin"
	"go-crud-demo/Dao"
	"go-crud-demo/Service/favorite"
	"net/http"
	"strconv"
)

type FavoriteResponse struct {
	Response
	VideoList []Dao.VideoList `json:"video_list,omitempty"`
}

func Upvote(c *gin.Context) {
	videoIdS := c.Query("video_id")
	actionTypeS := c.Query("action_type")
	videoId, _ := strconv.Atoi(videoIdS)
	actionType, _ := strconv.Atoi(actionTypeS)
	userIdS, _ := c.Get("uid")
	userId, _ := userIdS.(int)
	err := favorite.NewFavoriteFlow(videoId, actionType).Upvote(userId)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
	}
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  "成功",
	})
}

func FavoriteList(c *gin.Context) {
	//userIdS, _ := c.Get("uid")
	//userId, _ := userIdS.(int)
	userIdS := c.Query("user_id")
	userId, _ := strconv.Atoi(userIdS)
	favoriteList, err := favorite.FavoriteList(userId)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
	}
	c.JSON(http.StatusOK, FavoriteResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "成功",
		},
		VideoList: favoriteList,
	})
}
