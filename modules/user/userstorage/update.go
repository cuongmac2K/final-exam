package userstorage

import (
	"context"
	"finnal-exam/common"
	"finnal-exam/modules/user/usermodel"
	"gorm.io/gorm"
)

func (s *sqlStore) UpdateData(
	ctx context.Context,
	id int,
	data *usermodel.UserUpdate,
) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
func (s *sqlStore) IncreaseFollowingCount(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(usermodel.User{}.TableName()).
		Where("id = ?", id).
		Update("following_count", gorm.Expr("following_count + ?", 1)).
		Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
func (s *sqlStore) IncreaseFollowerCount(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(usermodel.User{}.TableName()).
		Where("id = ?", id).
		Update("follower_count", gorm.Expr("follower_count + ?", 1)).
		Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
func (s *sqlStore) DeIncreaseFollowingCount(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(usermodel.User{}.TableName()).
		Where("id = ? ", id).
		Update("following_count", gorm.Expr("following_count - ?", 1)).
		Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
func (s *sqlStore) DeIncreaseFollowerCount(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(usermodel.User{}.TableName()).
		Where("id = ? ", id).
		Update("follower_count", gorm.Expr("follower_count - ?", 1)).
		Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
