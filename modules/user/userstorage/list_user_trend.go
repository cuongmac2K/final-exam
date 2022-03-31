package userstorage

import (
	"context"
	"finnal-exam/common"
	"finnal-exam/modules/user/usermodel"
)

func (s *sqlStore) ListUserByFollwer(ctx context.Context,
	conditions map[string]interface{},
	filter *usermodel.Trending,
	paging *common.Paging,
	moreKeys ...string,
) ([]usermodel.UserTren, error) {
	var result []usermodel.UserTren

	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	db = db.Table("users").Where(conditions).Where("status in (1)").Joins("LEFT OUTER JOIN user_follows ON user_follows.user_id = users.id")
	var limit int
	if v := filter; v != nil {
		if v.Day != "" {

			db = db.Where("DAY(user_follows.created_at) = ?", v.Day)
		}
		if v.Month != "" {
			db = db.Where("Month(user_follows.created_at) = ?", v.Month)
		}
		if v.Year != "" {
			db = db.Where("Year(user_follows.created_at) = ?", v.Year)
		}
		if v.Limit != 0 {
			limit = v.Limit
		} else {
			limit = 50
		}
	}
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if err := db.
		Offset((paging.Page - 1) * paging.Limit).
		Limit(limit).
		Order("follower_count desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
