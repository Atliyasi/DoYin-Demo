package Controller

import (
	"github.com/gin-gonic/gin"
	"go-crud-demo/Dao"
	"go-crud-demo/Service/comment"
	"net/http"
	"strconv"
)

type CommentResponse struct {
	Response
	Comment Dao.CommentList `json:"comment"`
}

type CommentListResponse struct {
	Response
	CommentList []Dao.CommentList `json:"comment_list"`
}

func HandleComment(c *gin.Context) {
	userIdA, _ := c.Get("uid")
	userId := userIdA.(int)
	videoIdS := c.Query("video_id")
	videoId, _ := strconv.Atoi(videoIdS)
	actionType := c.Query("action_type")
	commentText := c.Query("comment_text")
	commentId := c.Query("comment_id")
	if actionType == "1" {
		thisComment, err := comment.NewComment(videoId, actionType, commentText).AppendComment(userId)
		if err != nil {
			c.JSON(http.StatusOK, CommentResponse{
				Response: Response{
					StatusCode: 1,
					StatusMsg:  err.Error(),
				},
				Comment: Dao.CommentList{},
			})
		}
		c.JSON(http.StatusOK, CommentResponse{
			Response: Response{
				StatusCode: 0,
				StatusMsg:  "成功",
			},
			Comment: Dao.CommentList{
				Id:         thisComment.Id,
				User:       Dao.VideoUser{},
				Content:    thisComment.Content,
				CreateDate: thisComment.CreateDate,
			},
		})
	} else if actionType == "2" {
		if err := comment.NewComment(videoId, actionType, commentId).DeleteComment(); err != nil {
			c.JSON(http.StatusOK, CommentResponse{
				Response: Response{
					StatusCode: 1,
					StatusMsg:  err.Error(),
				},
				Comment: Dao.CommentList{},
			})
		}
		c.JSON(http.StatusOK, CommentResponse{
			Response: Response{
				StatusCode: 0,
				StatusMsg:  "删除成功",
			},
			Comment: Dao.CommentList{},
		})
	} else {
		c.JSON(http.StatusOK, CommentResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "失败",
			},
			Comment: Dao.CommentList{},
		})
	}
}

func GetCommentList(c *gin.Context) {
	videoIdS := c.Query("video_id")
	videoId, _ := strconv.Atoi(videoIdS)
	commentList, err := comment.NewComment(videoId, "", "").GetCommentList()
	if err != nil {
		c.JSON(http.StatusBadRequest, CommentListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
			CommentList: commentList,
		})
	}
	c.JSON(http.StatusOK, CommentListResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "成功",
		},
		CommentList: commentList,
	})
}
