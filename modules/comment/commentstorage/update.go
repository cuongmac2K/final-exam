package commentstorage

import (
	"context"
	"finnal-exam/common"
	"finnal-exam/modules/comment/commentmodel"
	"gorm.io/gorm"
)

func (s *sqlStore) DecreaseLikeCount(ctx context.Context,
	id int) error {
	if err := s.db.
		Table(commentmodel.Comment{}.TableName()).
		Where("id = ?", id).
		Update("likes_count", gorm.Expr("likes_count - ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
func (s *sqlStore) IncreaseLikeCount(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table("comments").Where("id = ?", id).
		Update("likes_count", gorm.Expr("likes_count + ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
