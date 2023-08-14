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

//var (
//	lastMessageTimestamp map[int]int64
//	lock                 sync.Mutex
//)
//
//func CacheLastMessageTimestamp(timestamp int64, userId int) {
//	lock.Lock()
//	defer lock.Unlock()W
//	lastMessageTimestamp[userId] = timestamp
//}
//
//func ShouldReturnNewMessages(newTimestamp int64, userId int) bool {
//	lock.Lock()
//	defer lock.Unlock()
//	lastTimestamp := lastMessageTimestamp[userId]
//	return newTimestamp > lastTimestamp
//}

func SendMessage(c *gin.Context) {
	now := time.Now().Unix()
	fromUserIdA, _ := c.Get("uid")
	toUserIdS := c.Query("to_user_id")
	fromUserId := fromUserIdA.(int)
	toUserId, _ := strconv.Atoi(toUserIdS)
	actionType := c.Query("action_type")
	content := c.Query("content")
	if actionType == "1" {
		if err := message.SendMessage(fromUserId, toUserId, content, now); err != nil {
			c.JSON(http.StatusBadRequest, Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, Response{
			StatusCode: 1,
			StatusMsg:  "失败",
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  "发送成功",
	})
}

func GetMessage(c *gin.Context) {
	userIdA, _ := c.Get("uid")
	toUserIdS := c.Query("to_user_id")
	times := c.Query("pre_msg_time")
	userId := userIdA.(int)
	toUserId, _ := strconv.Atoi(toUserIdS)
	msgTime, _ := strconv.Atoi(times)
	newTime := int64(msgTime)
	//if newTime == 0 {
	//	return
	//}
	messageYesList, err := message.GetMessageList(userId, toUserId, newTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, MessageResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
			MessageList: nil,
		})
		return
	}
	c.JSON(http.StatusOK, MessageResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "成功",
		},
		MessageList: messageYesList,
	})
}
