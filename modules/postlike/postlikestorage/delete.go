package postlikestorage

import (
	"context"
	"finnal-exam/common"
	"finnal-exam/modules/postlike/postlikemodel"
	"log"
)

func (s *sqlStore) Delete(ctx context.Context, userId, id int) error {
	db := s.db
	log.Println("========", userId, "==", id, "=====")
	if err := db.Table("post_likes").
		Where("user_id =? and id =?", userId, id).
		Delete(nil).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) CheckPostLike(ctx context.Context, userId, id int) bool {
	var data postlikemodel.PostLike
	if err := s.db.Table("post_likes").
		Where("user_id =? and id =?  ", userId, id).
		First(&data).Error; err != nil {
		return false
	}
	return true
}
