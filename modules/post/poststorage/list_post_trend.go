package poststorage

import (
	"context"
	"finnal-exam/common"
	"finnal-exam/modules/post/postmodel"
	"fmt"
)

func (s *sqlStore) ListTrendByCondition(ctx context.Context,
	conditions map[string]interface{},
	filter *postmodel.FilterTrend,
	paging *common.Paging,
	moreKeys ...string,
) ([]postmodel.PostTrend, error) {
	var result []postmodel.PostTrend

	db := s.db

	db = db.Table(postmodel.Post{}.TableName()).Where(conditions).Where("posts.status = 1").Joins("LEFT OUTER JOIN post_likes ON post_likes.id = posts.id")

	if v := filter; v != nil {
		fmt.Println("date", v.Date)
		if v.Date != "" {
			db = db.Where("DAY(post_likes.created_at) =?", v.Date)
		}
		if v.Date != "" {
			db = db.Where("DAY(post_likes.created_at) BETWEEN ? AND ?", v.Date, v.Date+"7")
		}
		if v.Month != "" {
			db = db.Where(" Month(post_likes.created_at) = ?", v.Month)
		}
		if v.Year != "" {
			db = db.Where(" Year(post_likes.created_at) = ?", v.Year)
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
		Limit(3).
		Order("likes_count desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
