package userstorage

import (
	"context"
	"finnal-exam/common"
	"finnal-exam/modules/user/usermodel"
)

func (s *sqlStore) UpdateVerify(ctx context.Context, id int) error {
	if err := s.db.Table(usermodel.User{}.TableName()).
		Where("id = ?", id).
		Update("is_verified", 1).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
