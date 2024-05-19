package datastorer

import (
	"fmt"
	
	"smuggr/net-work/common/logger"
)

func GetUser(login string) (*User, *logger.MessageWrapper) {
	var user User
	if result := DB.Where("login = ?", login).First(&user); result.Error != nil {
		return nil, logger.ErrFetchingResourceFromDB.Format(login, logger.ResourceUser)
	}

	return &user, nil
}

func GetLimitedUsers(limit int) ([]*User, *logger.MessageWrapper) {
	var users []*User
	if err := DB.Limit(limit).Find(&users).Error; err != nil {
		return nil, logger.ErrFetchingResourceFromDB.Format(fmt.Sprintf("limit %d", limit), logger.ResourceUser)
	}

	return users, nil
}

func GetPaginatedUsers(page int, pageSize int) ([]*User, *logger.MessageWrapper) {
	var users []*User
	if err := DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&users).Error; err != nil {
		return nil, logger.ErrFetchingResourceFromDB.Format(fmt.Sprintf("page %d, pageSize %d", page, pageSize), logger.ResourceUser)
	}

	return users, nil
}