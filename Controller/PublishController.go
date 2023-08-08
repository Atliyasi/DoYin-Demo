package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-crud-demo/Dao"
	middleware "go-crud-demo/Middleware"
	"go-crud-demo/Util"
	"gorm.io/gorm"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
)

func PublishAction(c *gin.Context) {
	//userId, _ := c.Get("uid")
	//id, _ := userId.(int64)
	token := c.PostForm("token")
	_, claims, err := middleware.ParseToken(token)
	id := claims.UserId
	data, err := c.FormFile("data")
	if err != nil {
		fmt.Println("c.FormFile(\"data\"), err: ", err)
		c.JSON(http.StatusBadRequest, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s", id, filename)
	saveFile := filepath.Join("./public/mov/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		fmt.Println("c.SaveUploadedFile(data, saveFile), err: ", err)
		c.JSON(http.StatusBadRequest, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	snapshotName := Util.GetSnapshot(saveFile, finalName)
	fmt.Println(snapshotName)
	title := c.PostForm("title")
	video := &Dao.Video{
		Model:         gorm.Model{},
		Author:        int64(id),
		PlayUrl:       finalName,
		CoverUrl:      snapshotName,
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		Title:         title,
	}
	err = Dao.NewVideDao().SetVideo(video)
	if err != nil {
		fmt.Println("c.SaveUploadedFile(data, saveFile), err: ", err)
		c.JSON(http.StatusBadRequest, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	err = Dao.NewVideDao().UpdateVideoUserByWork(id)
	if err != nil {
		fmt.Println("Dao.NewVideDao().UpdateVideoUserByWork(id), err: ", err)
		c.JSON(http.StatusBadRequest, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + "上传成功",
	})
	// publishResponse := &message.DouyinPublishActionResponse{}
	//userId, _ := c.Get("UserId")
	//userId, err := common.VerifyToken(token)
	//title := c.PostForm("title")
	//data, err := c.FormFile("data")
	//if err != nil {
	//	fmt.Println("c.FormFile(\"data\")")
	//	c.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	//	return
	//}
	//filename := filepath.Base(data.Filename)
	//
	//finalName := fmt.Sprintf("%s_%s", userId, filename)
	//saveFile := filepath.Join("./public/mov/", finalName)
	//fmt.Println("finalName: ", finalName)
	//fmt.Println("saveFile: ", saveFile)
	//if err := c.SaveUploadedFile(data, saveFile); err != nil {
	//	fmt.Println("c.SaveUploadedFile(data, saveFile), err", err)
	//	c.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	//	return
	//}
	//c.JSON(http.StatusOK, Response{
	//	StatusCode: 0,
	//	StatusMsg:  finalName + " uploaded successfully",
	//})
}

func PublishList(c *gin.Context) {
	userId := c.Query("user_id")
	id, _ := strconv.Atoi(userId)
	fmt.Println("id: ", id)
	if id == 0 {
		token := c.Query("token")
		_, claims, err := middleware.ParseToken(token)
		if err != nil {
			return
		}
		id = claims.UserId
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "请求成功",
		},
		VideoList: Dao.NewVideDao().GetVideoListById(id),
		NextTime:  time.Now().Unix(),
	})
}
