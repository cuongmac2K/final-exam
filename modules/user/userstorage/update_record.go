package userstorage

import (
	"context"
	"finnal-exam/common"
	"finnal-exam/modules/user/usermodel"
)

func (s *sqlStore) UpdateRecord(
	ctx context.Context,
	id int,
	record *common.Record,
) error {
	db := s.db

	if err := db.Table(usermodel.User{}.TableName()).Where("id = ?", id).Update("record", record).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
