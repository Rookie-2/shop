package biz

import "shop/user/internal/shop/store"

type UserBiz interface {
	Create()
}

type userBiz struct {
	store store.IStore
}

func NewUserBiz(store store.IStore) *userBiz {
	return &userBiz{store: store}
}

func (u *userBiz) Create() {

}