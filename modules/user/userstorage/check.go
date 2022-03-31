package userstorage

import (
	"context"
	"finnal-exam/modules/user/usermodel"
)

func (s *sqlStore) CheckUser(ctx context.Context, id int) bool {
	var user usermodel.User
	if err := s.db.Table(usermodel.User{}.TableName()).
		Where("id = ?", id).
		First(&user).Error; err != nil {
		return false
	}
	return true
}
