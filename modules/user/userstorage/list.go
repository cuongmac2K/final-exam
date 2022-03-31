package userstorage

import (
	"context"
	"finnal-exam/common"
	"finnal-exam/modules/user/usermodel"
)

func (s *sqlStore) ListDataByCondition(ctx context.Context,
	conditions map[string]interface{},
	filter *usermodel.FindUser,
	paging *common.Paging,
	moreKeys ...string,
) ([]usermodel.UserInfo, error) {
	var result []usermodel.UserInfo

	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	db = db.Table("users").Where(conditions).Where("status in (1)")

	if v := filter; v != nil {
		if v.Email != "" {
			db = db.Where("email LIKE ?", "%"+v.Email+"%")
		}
		if v.LastName != "" {
			db = db.Where("last_name LIKE ?", "%"+v.LastName+"%")
		}
		if v.FirstName != "" {
			db = db.Where("first_name LIKE ?", "%"+v.FirstName+"%")
		}
		if v.Phone != "" {
			db = db.Where("phone LIKE ?", "%"+v.Phone+"%")
		}

	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if err := db.
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Order("email desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
