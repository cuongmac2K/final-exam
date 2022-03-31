package poststorage

import (
	"context"
	"finnal-exam/common"
	"finnal-exam/modules/post/postmodel"
)

func (s *sqlStore) ListDataByCondition(ctx context.Context,
	conditions map[string]interface{},
	id int,
	filter *postmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]postmodel.Post, error) {
	var result []postmodel.Post

	db := s.db
	dbcomment := db

	db = db.Table(postmodel.Post{}.TableName()).Where(conditions).Where("posts.status = 1")
	if v := filter; v != nil {
		if v.Email != "" {
			db = db.Where("users.email = ?", v.Email)
			db = db.Joins("LEFT OUTER JOIN users ON users.id = posts.user_id")
		}
		if v.Username != "" {
			db = db.Where("CONCAT(users.first_name,' ',users.last_name) LIKE ?", "%"+v.Username+"%")
			db = db.Joins("LEFT OUTER JOIN users ON users.id = posts.user_id")
		}
		if v.Caption != "" {
			db = db.Where("posts.caption LIKE ?", "%"+v.Caption+"%")
		}
		if (v.DateFrom != "") && (v.DateTo != "") {
			db = db.Where("posts.created_at BETWEEN ? AND ?", v.DateFrom, v.DateTo)
		}
		if v.IsFollowing == true {
			db = db.Where("user_follows.user_id = ?", id)
			db = db.Joins("LEFT OUTER JOIN user_follows ON user_follows.id = posts.user_id")
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Order("created_at desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	dbcomment.Preload("SimpleUser")
	for i := range result {

		dbcomment.
			Table(common.SimpleComment{}.TableName()).
			Limit(3).
			Where("post_id = ?", result[i].Id).
			Where("status = 1").
			Order("created_at desc").
			Find(result[i].Comments)

		//for j := range *result[i].Comments {
		//	if err := dbuser.Table(common.SimpleUser{}.TableName()).
		//		Where("users.id = ?", (*result[i].Comments)[j].UserId).
		//		Find(&user).Error; err != nil {
		//		return nil, common.ErrDB(err)
		//	}
		//	*result[i].Comments)[j].SimpleUser = &user
		//}

		for j := range *result[i].Comments {
			var user common.SimpleUser
			dbuser := s.db
			if err := dbuser.Table(common.SimpleUser{}.TableName()).
				Where("id = ?", (*result[i].Comments)[j].UserId).
				Find(&user).Error; err != nil {
				return nil, common.ErrDB(err)
			}
			(*result[i].Comments)[j].SimpleUser = &user
		}
	}

	return result, nil
}
