package relation

import "go-crud-demo/Dao"

// FollowerList 实现查询粉丝列表
func FollowerList(id int) ([]Dao.VideoUser, error) {
	videoUserList, err := Dao.NewRelationDao().FollowerList(id)
	if err != nil {
		return nil, err
	}
	return videoUserList, nil
}
