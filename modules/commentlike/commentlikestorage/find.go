package commentlikestorage

import (
	"context"
	"finnal-exam/modules/commentlike/commentlikemodel"
)

func (s *sqlStore) Find(ctx context.Context, id, userId int) bool {
	var data commentlikemodel.CommentLike
	if err := s.db.Table(commentlikemodel.CommentLike{}.TableName()).
		Where("id = ? and user_id = ?", id, userId).
		First(&data).Error; err != nil {
		return false
	}
	return true
}
