package biz

import "shop/user/internal/shop/store"

type IBiz interface {
	Users(store store.IStore) UserBiz
}

type biz struct {
	store store.IStore
}

func NewBiz(store store.IStore) *biz {
	return &biz{store: store}
}

func (b *biz) Users(store store.IStore) UserBiz {
	return NewUserBiz(store)
}
