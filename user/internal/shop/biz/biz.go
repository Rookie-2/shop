package biz

import "shop/user/internal/shop/store"

type IBiz interface {
	Users()
}

type biz struct {
	store *store.IStore
}

func newBiz(store *store.IStore) *biz {
	return &biz{store: store}
}

func (b *biz) Users() {

}
