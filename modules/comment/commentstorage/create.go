package commentstorage

import (
	"context"
	"finnal-exam/common"
	"finnal-exam/modules/comment/commentmodel"
)

func (s *sqlStore) CreateData(ctx context.Context, data *commentmodel.CommentCreate) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		return common.ErrCannotCreateEntity(commentmodel.EntityName, err)
	}
	return nil
}
