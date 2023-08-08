package Dao

import (
	"errors"
	"gorm.io/gorm"
	"sync"
)

type FavoriteList struct {
	gorm.Model
	UserId     int  `gorm:"column:user_id"`
	VideoId    int  `gorm:"column:video_id"`
	IsFavorite bool `gorm:"column:is_favorite"`
}

type FavoriteDao struct{}

var favoriteDao *FavoriteDao
var FavoriteOnce sync.Once

func NewFavoriteDao() *FavoriteDao {
	FavoriteOnce.Do(func() {
		favoriteDao = &FavoriteDao{}
	})
	return favoriteDao
}

// Like 点赞 +1
func (f *FavoriteDao) Like(videoId int, userId int) error {
	db := GetDB()
	// 开始事务
	tx := db.Begin()
	var video Video
	if err := tx.First(&video, "id=?", videoId).Error; err != nil {
		tx.Rollback() // 回滚事务
		return err
	}
	video.FavoriteCount++
	if err := tx.Save(&video).Error; err != nil {
		tx.Rollback() // 回滚事务
		return err
	}
	var videoUser VideoUser
	if err := tx.Where("id=?", video.Author).First(&videoUser).Error; err != nil || videoUser.Id == int64(userId) {
		tx.Rollback() // 回滚事务
		return errors.New("有误")
	}
	//var videoUser1 VideoUser
	//if err := tx.Where("id=?", userId).First(&videoUser1).Error; err != nil {
	//	tx.Rollback() // 回滚事务
	//	return err
	//}
	videoUser.TotalFavorited++
	if err := tx.Save(&videoUser).Error; err != nil {
		tx.Rollback() // 回滚事务
		return err
	}
	// 提交事务
	var favoriteList FavoriteList
	if err := tx.Where("user_id=?", userId).Where("video_id=?", videoId).First(&favoriteList).Error; err != nil {
		//fmt.Println("tx.Where(\"user_id=?\", userId).Where(\"video_id=?\", videoId).First(&favoriteList).Error: ", err)
		if err := f.CreateFavoriteList(userId, videoId, true); err != nil {
			return err
		}
	} else {
		if err := f.UpdateFavoriteList(userId, videoId, true); err != nil {
			return err
		}
	}
	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

// Unlike 取消点赞 -1
func (f *FavoriteDao) Unlike(videoId int, userId int) error {
	db := GetDB()
	// 开始事务
	tx := db.Begin()
	var video Video
	if err := tx.First(&video, "id=?", videoId).Error; err != nil {
		tx.Rollback() // 回滚事务
		return err
	}
	video.FavoriteCount--
	if err := tx.Save(&video).Error; err != nil {
		tx.Rollback() // 回滚事务
		return err
	}
	var videoUser VideoUser
	if err := tx.Where("id=?", video.Author).First(&videoUser).Error; err != nil || videoUser.Id == int64(userId) {
		tx.Rollback() // 回滚事务
		return err
	}
	//var videoUser1 VideoUser
	//if err := tx.Where("id=?", userId).First(&videoUser1).Error; err != nil {
	//	tx.Rollback() // 回滚事务
	//	return err
	//}
	videoUser.TotalFavorited--
	if err := tx.Save(&videoUser).Error; err != nil {
		tx.Rollback() // 回滚事务
		return err
	}
	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return err
	}
	if err := f.UpdateFavoriteList(userId, videoId, false); err != nil {
		return err
	}
	return nil
}

func (*FavoriteDao) CreateFavoriteList(userId int, videoId int, isFavorite bool) error {
	favorite := &FavoriteList{
		Model:      gorm.Model{},
		UserId:     userId,
		VideoId:    videoId,
		IsFavorite: isFavorite,
	}
	if err := GetDB().Create(favorite).Error; err != nil {
		return err
	}
	return nil
}

func (*FavoriteDao) UpdateFavoriteList(userId int, videoId int, isFavorite bool) error {
	err := GetDB().Model(&FavoriteList{}).Where("user_id=?", userId).Where("video_id=?", videoId).Update("IsFavorite", isFavorite).Error
	if err != nil {
		return err
	}
	return nil
}

func (*FavoriteDao) FindFavoriteByUserIdByVideoId(userId int, videoId int) bool {
	var favoriteList FavoriteList
	if err := GetDB().Where("user_id=?", userId).Where("video_id=?", videoId).Find(&favoriteList).Error; err != nil {
		return false
	}
	return favoriteList.IsFavorite
}

func (*FavoriteDao) FindFavoriteList(userId int) ([]VideoList, error) {
	var favoriteList []FavoriteList
	if err := GetDB().Where("user_id=?", userId).Where("is_favorite=?", true).Find(&favoriteList).Error; err != nil {
		return nil, err
	}
	var videoList []VideoList
	for _, favorite := range favoriteList {
		videoUser := NewVideDao().GetVideoUser(int64(favorite.UserId))
		video := NewVideDao().GetVideoById(int64(favorite.VideoId))
		videoList = append(videoList, VideoList{
			Id:            int64(video.ID),
			Author:        *videoUser,
			PlayUrl:       "http://172.16.2.131:8080/public/mov/" + video.PlayUrl,
			CoverUrl:      "http://172.16.2.131:8080/public/pic/" + video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    true,
			Title:         video.Title,
		})
	}
	return videoList, nil
}
