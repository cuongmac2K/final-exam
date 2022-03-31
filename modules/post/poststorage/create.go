package poststorage

import (
	"context"
	"finnal-exam/common"
	"finnal-exam/modules/post/postmodel"
)

func (s *sqlStore) CreateData(ctx context.Context, data *postmodel.PostCreate) error {
	db := s.db
	//for i := range *data.Images {
	//	(*data.Images)[i].Id = i + 1
	//}
	if err := db.Create(data).Error; err != nil {
		return common.ErrCannotCreateEntity(postmodel.EntityName, err)
	}
	return nil
}
