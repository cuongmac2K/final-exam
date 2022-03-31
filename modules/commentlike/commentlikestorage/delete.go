package commentlikestorage

import (
	"context"
	"finnal-exam/common"
	"finnal-exam/modules/commentlike/commentlikemodel"
)

func (s *sqlStore) DeleteLike(ctx context.Context,
	id, userId int) error {
	db := s.db
	if err := db.Table(commentlikemodel.CommentLike{}.TableName()).
		Where("id = ? and user_id = ?", id, userId).
		Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
