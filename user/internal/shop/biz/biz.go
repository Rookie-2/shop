package biz

import "shop/user/internal/shop/store"

type IBiz interface {
	Users() UserBiz
}

type biz struct {
	store store.IStore
}

func NewBiz(store store.IStore) *biz {
	return &biz{store: store}
}

func (b *biz) Users() UserBiz {
	return NewUserBiz(b.store)
}
