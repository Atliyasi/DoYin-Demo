package Controller

import (
	"github.com/gin-gonic/gin"
	"go-crud-demo/Service/message"
	"net/http"
	"strconv"
	"time"
)

type MessageResponse struct {
	Response
	MessageList []message.MessageYes `json:"message_list"`
}

func SendMessage(c *gin.Context) {
	fromUserIdA, _ := c.Get("uid")
	toUserIdS := c.Query("to_user_id")
	fromUserId := fromUserIdA.(int)
	toUserId, _ := strconv.Atoi(toUserIdS)
	actionType := c.Query("action_type")
	content := c.Query("content")
	now := time.Now().Unix()
	if actionType == "1" {
		if err := message.SendMessage(fromUserId, toUserId, content, now); err != nil {
			c.JSON(http.StatusBadRequest, Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, Response{
			StatusCode: 1,
			StatusMsg:  "失败",
		})
	}
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  "发送成功",
	})
}

func GetMessage(c *gin.Context) {
	userIdA, _ := c.Get("uid")
	toUserIdS := c.Query("to_user_id")
	userId := userIdA.(int)
	toUserId, _ := strconv.Atoi(toUserIdS)
	messageYesList, err := message.GetMessageList(userId, toUserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, MessageResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
			MessageList: nil,
		})
	}
	c.JSON(http.StatusOK, MessageResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "成功",
		},
		MessageList: messageYesList,
	})
}
