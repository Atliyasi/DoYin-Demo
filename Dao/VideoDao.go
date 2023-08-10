package Dao

import (
	"fmt"
	"gorm.io/gorm"
	"sync"
	"time"
)

type VideoUser struct {
	Id              int64  `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	FollowCount     int64  `json:"follow_count,omitempty"`
	FollowerCount   int64  `json:"follower_count,omitempty"`
	IsFollow        bool   `json:"is_follow,omitempty"`
	Avatar          string `json:"avatar,omitempty"`
	BackgroundImage string `json:"background_image,omitempty"`
	Signature       string `json:"signature,omitempty"`
	TotalFavorited  int    `json:"total_favorited,omitempty"`
	WorkCount       int    `json:"work_count,omitempty"`
	FavoriteCount   int    `json:"favorite_count,omitempty"`
}

type Video struct {
	gorm.Model
	Author        int64  `json:"author"`
	PlayUrl       string `json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
	Title         string `json:"title,omitempty"`
}
type VideoList struct {
	Id            int64     `json:"id"`
	Author        VideoUser `json:"author"`
	PlayUrl       string    `json:"play_url,omitempty"`
	CoverUrl      string    `json:"cover_url,omitempty"`
	FavoriteCount int64     `json:"favorite_count,omitempty"`
	CommentCount  int64     `json:"comment_count,omitempty"`
	IsFavorite    bool      `json:"is_favorite,omitempty"`
	Title         string    `json:"title,omitempty"`
}

type VideoDao struct{}

var videoDao *VideoDao
var VideoOnce sync.Once

func NewVideDao() *VideoDao {
	VideoOnce.Do(func() {
		videoDao = &VideoDao{}
	})
	return videoDao
}

func (*VideoDao) GetVideoUser(id int64) *VideoUser {
	var videoUser VideoUser
	if err := GetDB().Where("id=?", id).Find(&videoUser).Error; err != nil {
		return nil
	}
	return &videoUser
}

func (this *VideoDao) GetVideoList(userId int) []VideoList {
	var video []Video
	err := GetDB().Find(&video).Error
	if err != nil {
		return nil
	}
	var videos []VideoList
	favoriteList, err := NewFavoriteDao().FindFavoritesByUserId(userId)
	if err != nil {
		return nil
	}
	favorites := make(map[int]bool)
	for _, fav := range favoriteList {
		favorites[fav.VideoId] = true
	}
	follows := make(map[int]bool)
	relations := NewRelationDao().FollowById(userId)
	var lock sync.Mutex
	var wg sync.WaitGroup
	wg.Add(len(relations))
	for _, relation := range relations {
		go func(relation Relation) {
			defer wg.Done()
			var userToFetch int
			if relation.UserIdOne == userId {
				userToFetch = relation.UserIdTwo
			} else {
				userToFetch = relation.UserIdOne
			}
			lock.Lock()
			follows[userToFetch] = true
			lock.Unlock()
		}(relation)
	}
	wg.Wait()
	for _, videoInfo := range video {
		videoUser := this.GetVideoUser(videoInfo.Author)
		videoUser.IsFollow = follows[int(videoUser.Id)]
		videos = append(videos, VideoList{
			Id:            int64(videoInfo.ID),
			Author:        *videoUser,
			PlayUrl:       fmt.Sprintf("http://172.16.2.131:8080/public/mov/%s", videoInfo.PlayUrl),
			CoverUrl:      fmt.Sprintf("http://172.16.2.131:8080/public/pic/%s", videoInfo.CoverUrl),
			FavoriteCount: videoInfo.FavoriteCount,
			CommentCount:  videoInfo.CommentCount,
			IsFavorite:    favorites[int(videoInfo.ID)],
			Title:         videoInfo.Title,
		})
	}
	return videos
}

func (this *VideoDao) GetVideoListById(id int) []VideoList {
	var video []Video
	err := GetDB().Find(&video, "author=?", id).Error
	//fmt.Println("video：", video)
	if err != nil {
		return nil
	}
	var videos []VideoList
	for _, videoInfo := range video {
		videoUser := this.GetVideoUser(videoInfo.Author)
		videos = append(videos, VideoList{
			Id:            int64(videoInfo.ID),
			Author:        *videoUser,
			PlayUrl:       fmt.Sprintf("http://172.16.2.131:8080/public/mov/%s", videoInfo.PlayUrl),
			CoverUrl:      fmt.Sprintf("http://172.16.2.131:8080/public/pic/%s", videoInfo.CoverUrl),
			FavoriteCount: videoInfo.FavoriteCount,
			CommentCount:  videoInfo.CommentCount,
			IsFavorite:    videoInfo.IsFavorite,
			Title:         videoInfo.Title,
		})
	}
	return videos
}

func (*VideoDao) GetVideoById(id int64) *Video {
	var videoList Video
	err := GetDB().Where("id=?", id).First(&videoList).Error
	if err != nil {
		return nil
	}
	return &videoList
}

func (*VideoDao) SetVideo(video *Video) error {
	if err := GetDB().Create(video).Error; err != nil {
		return err
	}
	return nil
}
func (*VideoDao) QueryVideoListByTime(lastTime time.Time, len int) (*[]Video, error) {
	var videos []Video
	if err := GetDB().Where("updated_time < ?", lastTime).Order("updated_time desc").Limit(len).Find(&videos).Error; err != nil {
		return nil, err
	}
	return &videos, nil
}

func (*VideoDao) UpdateVideoUserByWork(id int) error {
	var videoUser VideoUser
	err := GetDB().Where("id=?", id).Find(&videoUser).Error
	if err != nil {
		return err
	}
	videoUser.WorkCount += 1
	user := &VideoUser{}
	err = GetDB().Model(&user).Where("id=?", id).Update("WorkCount", videoUser.WorkCount).Error
	if err != nil {
		return err
	}
	return nil
}
