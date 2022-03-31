package commentstorage

import (
	"context"
	"finnal-exam/common"
	"finnal-exam/modules/comment/commentmodel"
	"gorm.io/gorm"
)

func (s *sqlStore) FindDataByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*commentmodel.Comment, error) {
	var result commentmodel.Comment

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
