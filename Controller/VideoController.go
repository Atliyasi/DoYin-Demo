package Controller

import (
	"github.com/gin-gonic/gin"
	"go-crud-demo/Dao"
	middleware "go-crud-demo/Middleware"
	"net/http"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []Dao.VideoList `json:"video_list,omitempty"`
	NextTime  int64           `json:"next_time,omitempty"`
}

// Feed same demo mov list for every request
func Feed(c *gin.Context) {
	token := c.Query("token")
	_, claims, _ := middleware.ParseToken(token)
	c.JSON(http.StatusOK, FeedResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "请求成功",
		},
		VideoList: Dao.NewVideDao().GetVideoList(claims.UserId),
		NextTime:  time.Now().Unix(),
	})
}
