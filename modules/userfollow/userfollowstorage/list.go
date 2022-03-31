package userfollowstorage

import (
	"context"
	"finnal-exam/common"
	userfollowmodel "finnal-exam/modules/userfollow/userfollowmodel"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"time"
)

const timeLayout = "2006-01-02T15:04:05.999999"

func (s *sqlStore) GetUserFollow(ctx context.Context, ids []int) (map[int]int, error) {
	result := make(map[int]int)

	type sqlData struct {
		Id          int `gorm:"column:id;"`
		FollowCount int `gorm:"column:count;"`
	}

	var listFollow []sqlData

	if err := s.db.Table(userfollowmodel.UserFollow{}.TableName()).
		Select("id, count(id) as count").
		Where("id in (?)", ids).
		Group("id").Find(&listFollow).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for _, item := range listFollow {
		result[item.Id] = item.FollowCount
	}

	return result, nil
	//return nil, errors.New("cannot get likes")
}

func (s *sqlStore) GetUsersFollow(ctx context.Context,
	conditions map[string]interface{},
	filter *userfollowmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]common.SimpleUser, error) {
	var result []userfollowmodel.UserFollow

	db := s.db

	db = db.Table(userfollowmodel.UserFollow{}.TableName()).Where(conditions)

	if v := filter; v != nil {
		if v.Id > 0 {
			db = db.Where("id = ?", v.Id)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	//for i := range moreKeys {
	//	db = db.Preload(moreKeys[i])
	//}

	db = db.Preload("User")

	if v := paging.FakeCursor; v != "" {
		timeCreated, err := time.Parse(timeLayout, string(base58.Decode(v)))

		if err != nil {
			return nil, common.ErrDB(err)
		}

		db = db.Where("created_at < ?", timeCreated.Format("2006-01-02 15:04:05"))
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.
		Limit(paging.Limit).
		Order("created_at desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	users := make([]common.SimpleUser, len(result))

	for i, item := range result {
		result[i].User.CreatedAt = item.CreatedAt
		result[i].User.UpdatedAt = nil
		users[i] = *result[i].User

		if i == len(result)-1 {
			cursorStr := base58.Encode([]byte(fmt.Sprintf("%v", item.CreatedAt.Format(timeLayout))))
			paging.NextCursor = cursorStr
		}
	}

	return users, nil
}

func (s *sqlStore) GetUsersFollowing(ctx context.Context,
	conditions map[string]interface{},
	filter *userfollowmodel.Filter,
	paging *common.Paging,
	userid int,
	moreKeys ...string,
) ([]common.SimpleUser, error) {
	var result []userfollowmodel.UserFollowing

	db := s.db

	db = db.Table(userfollowmodel.UserFollowing{}.TableName()).Where("user_id = ?", userid)

	//if v := filter; v != nil {
	//	if v.Id > 0 {
	//		db = db.Where("id = ?", v.Id)
	//	}
	//}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	//for i := range moreKeys {
	//	db = db.Preload(moreKeys[i])
	//}
	db = db.Preload("User")

	if v := paging.FakeCursor; v != "" {
		timeCreated, err := time.Parse(timeLayout, string(base58.Decode(v)))

		if err != nil {
			return nil, common.ErrDB(err)
		}

		db = db.Where("created_at < ?", timeCreated.Format("2006-01-02 15:04:05"))
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.
		Limit(paging.Limit).
		Order("created_at desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	users := make([]common.SimpleUser, len(result))

	for i, item := range result {
		result[i].User.CreatedAt = item.CreatedAt
		result[i].User.UpdatedAt = nil
		users[i] = *result[i].User

		if i == len(result)-1 {
			cursorStr := base58.Encode([]byte(fmt.Sprintf("%v", item.CreatedAt.Format(timeLayout))))
			paging.NextCursor = cursorStr
		}
	}

	return users, nil
}
