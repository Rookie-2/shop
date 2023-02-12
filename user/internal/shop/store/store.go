package store

import (
	"gorm.io/gorm"
	"sync"
)

var (
	once sync.Once
	S    *dataStore
)

// IStore Store 层需要实现的方法
type IStore interface {
	DB() *gorm.DB
	Users() UserStore
}

type dataStore struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *dataStore {
	once.Do(func() {
		S = &dataStore{
			db: db,
		}
	})
	return S
}

func (ds *dataStore) DB() *gorm.DB {
	return ds.db
}

func (ds *dataStore) Users() UserStore {
	return newUsers(ds.db)
}
