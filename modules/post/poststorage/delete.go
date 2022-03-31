package poststorage

import (
	"context"
	"finnal-exam/common"
	"finnal-exam/modules/post/postmodel"
)

func (s *sqlStore) DeleteData(
	ctx context.Context,
	id int,
) error {
	db := s.db

	if err := db.Table(postmodel.Post{}.TableName()).Where("id = ?", id).Update("status", 0).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
