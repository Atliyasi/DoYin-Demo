package relation

import "go-crud-demo/Dao"

// FriendList 返回朋友列表
func FriendList(userId int) ([]Dao.VideoUser, error) {
	VideoUserList, err := Dao.NewRelationDao().FriendList(userId)
	if err != nil {
		return nil, err
	}
	return VideoUserList, nil
}
