package favorite

import "go-crud-demo/Dao"

type FavoriteFlow struct {
	videoId    int
	actionType int
}

func NewFavoriteFlow(videoId int, actionType int) *FavoriteFlow {
	return &FavoriteFlow{
		videoId:    videoId,
		actionType: actionType,
	}
}

// Upvote 实现点赞或取消点赞逻辑
func (f *FavoriteFlow) Upvote(userId int) error {
	if f.actionType == 1 {
		err := Dao.NewFavoriteDao().Like(f.videoId, userId)
		if err != nil {
			return err
		}
	}
	if f.actionType == 2 {
		err := Dao.NewFavoriteDao().Unlike(f.videoId, userId)
		if err != nil {
			return err
		}
	}
	return nil
}

// FavoriteList 实现喜欢列表的显示
func FavoriteList(userId int) ([]Dao.VideoList, error) {
	videoList, err := Dao.NewFavoriteDao().FindFavoriteList(userId)
	if err != nil {
		return nil, err
	}
	return videoList, nil
}
