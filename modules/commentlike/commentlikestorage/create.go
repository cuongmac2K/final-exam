package commentlikestorage

import (
	"context"
	"finnal-exam/common"
	"finnal-exam/modules/comment/commentmodel"
	"finnal-exam/modules/commentlike/commentlikemodel"
)

func (s *sqlStore) Create(ctx context.Context, data *commentlikemodel.CommentLike) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (s *sqlStore) CheckComment(ctx context.Context, id int) bool {
	var data commentmodel.Comment
	if err := s.db.Table("comments").
		Where("id = ?", id).
		First(&data).Error; err != nil {
		return false
	}
	return true
}
