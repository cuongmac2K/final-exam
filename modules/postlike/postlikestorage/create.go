package postlikestorage

import (
	"context"
	"finnal-exam/common"
	"finnal-exam/modules/postlike/postlikemodel"
)

func (s *sqlStore) Create(ctx context.Context, data *postlikemodel.PostLike) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
