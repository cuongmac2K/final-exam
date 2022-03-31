package poststorage

import (
	"context"
	"finnal-exam/common"
	"finnal-exam/modules/post/postmodel"
	"gorm.io/gorm"
)

func (s *sqlStore) UpdateData(
	ctx context.Context,
	id int,
	data *postmodel.PostUpdate,
) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) IncreaseLikeCount(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(postmodel.Post{}.TableName()).Where("id = ?", id).
		Update("likes_count", gorm.Expr("likes_count + ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
func (s *sqlStore) DecreaseLikeCount(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(postmodel.Post{}.TableName()).
		Where("id = ?", id).Update("likes_count", gorm.Expr("likes_count - ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
