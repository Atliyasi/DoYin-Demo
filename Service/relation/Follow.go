package relation

import "go-crud-demo/Dao"

// FollowList 实现查询关注列表
func FollowList(id int) ([]Dao.VideoUser, error) {
	videoUserList, err := Dao.NewRelationDao().FollowList(id)
	if err != nil {
		return nil, err
	}
	return videoUserList, nil
}
