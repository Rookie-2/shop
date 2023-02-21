package biz

import (
	"context"
	userpb "shop/user/api-gen/v1"
	"shop/user/internal/shop/store"
)

type UserBiz interface {
	Create(ctx context.Context, r *userpb.CreateUserRequest) error
	List(ctx context.Context, offset, limit int) (*userpb.UserListResponse, error)
	Update(ctx context.Context, username string, r *userpb.UpdateUserRequest) error
	ListByMobile(ctx context.Context, mobile string) (*userpb.UserInfo, error)
	ListByUserId(ctx context.Context, userId string) (*userpb.UserInfo, error)
}

type userBiz struct {
	store store.IStore
}

func NewUserBiz(store store.IStore) *userBiz {
	return &userBiz{store: store}
}

func (ub *userBiz) Create(ctx context.Context, r *userpb.CreateUserRequest) error {
	return nil
}
func (ub *userBiz) List(ctx context.Context, offset, limit int) (*userpb.UserListResponse, error) {
	return nil, nil
}
func (ub *userBiz) Update(ctx context.Context, username string, r *userpb.UpdateUserRequest) error {
	return nil
}
func (ub *userBiz) ListByMobile(ctx context.Context, mobile string) (*userpb.UserInfo, error) {
	return nil, nil
}
func (ub *userBiz) ListByUserId(ctx context.Context, userId string) (*userpb.UserInfo, error) {
	return nil, nil
}
