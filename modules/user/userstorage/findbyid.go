package userstorage

import (
	"context"
	"finnal-exam/common"
	"finnal-exam/modules/user/usermodel"

	"gorm.io/gorm"
)

func (s *sqlStore) GetUserByID(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*usermodel.User, error) {
	var result usermodel.User

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
