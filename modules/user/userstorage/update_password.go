package userstorage

import (
	"context"
	"finnal-exam/common"
	"finnal-exam/modules/user/usermodel"
)

func (s *sqlStore) UpdatePassword(
	ctx context.Context,
	id int,
	user *usermodel.User,
) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(user).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
