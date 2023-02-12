package store

import "gorm.io/gorm"

type UserStore interface {
	Create()
}

type users struct {
	db *gorm.DB
}

func newUsers(db *gorm.DB) *users {
	return &users{db: db}
}

func (u *users) Create() {}
