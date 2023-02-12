package repo

import (
	"gorm.io/gorm"
	"shop/global"
	"shop/internal/pkg/model"
)

func GetUserList(page, pageSize int) ([]*model.Users, error) {
	var users []*model.Users
	result := global.DB.Scopes(Paginate(int(page), int(pageSize))).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// Paginate 分页, 感觉也没什么用...,接口开始前会检查request参数
func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
