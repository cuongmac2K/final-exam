package poststorage

import (
	"context"
	"finnal-exam/common"
	"finnal-exam/modules/comment/commentmodel"
	"finnal-exam/modules/post/postmodel"
	"gorm.io/gorm"
)

func (s *sqlStore) FindDataByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*postmodel.Post, error) {
	var result postmodel.Post

	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Where(conditions).
		First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &result, nil
}

func (s *sqlStore) FindDataByConditionForList(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*postmodel.Post, error) {
	var result postmodel.Post

	db := s.db
	dbcomment := db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Where(conditions).
		First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	dbcomment.
		Table(commentmodel.Comment{}.TableName()).
		Where("status = 1").
		Where("post_id = ?", result.Id).
		Order("created_at desc").
		Find(result.Comments)
	for j := range *result.Comments {
		var user common.SimpleUser
		dbuser := s.db
		if err := dbuser.Table(common.SimpleUser{}.TableName()).
			Where("id = ?", (*result.Comments)[j].UserId).
			Find(&user).Error; err != nil {
			return nil, common.ErrDB(err)
		}
		(*result.Comments)[j].SimpleUser = &user
	}
	return &result, nil
}

func (s *sqlStore) CheckPost(ctx context.Context, id int) bool {
	var data postmodel.Post
	if err := s.db.Table("posts").
		Where("id = ?", id).
		Where("status = 1").
		First(&data).Error; err != nil {
		return false
	}
	return true
}
