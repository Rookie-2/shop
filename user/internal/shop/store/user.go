package store

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"shop/known"
	"shop/user/internal/pkg/model"
)

type UserStore interface {
	Create(ctx context.Context, m *model.UserM) error
	ListByMobile(ctx context.Context, mobile string) (*model.UserM, error)
	ListByUserId(ctx context.Context, userId string) (*model.UserM, error)
	List(ctx context.Context, offset, limit int) (int64, []*model.UserM, error)
	Update(ctx context.Context, m *model.UserM) error
}

type users struct {
	db *gorm.DB
}

func newUsers(db *gorm.DB) *users {
	return &users{db: db}
}

func (u *users) Create(ctx context.Context, m *model.UserM) error {
	return u.db.WithContext(ctx).Create(&m).Error
}

func (u *users) ListByMobile(ctx context.Context, mobile string) (*model.UserM, error) {
	var user model.UserM
	err := u.db.WithContext(ctx).Model(new(model.UserM)).
		Where("mobile = ?", mobile).
		Where("is_del = ?", known.NotDelete).Find(&user).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return nil, err
}

func (u *users) ListByUserId(ctx context.Context, userId string) (*model.UserM, error) {
	var user model.UserM
	err := u.db.WithContext(ctx).Model(new(model.UserM)).
		Where("user_id = ?", userId).
		Where("is_del = ?", known.NotDelete).Find(&user).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return nil, err
}

func (u *users) List(ctx context.Context, offset, limit int) (int64, []*model.UserM, error) {
	var users []*model.UserM
	var count int64
	err := u.db.WithContext(ctx).Model(new(model.UserM)).
		Where("is_del = ?", known.NotDelete).Find(&users).
		Offset(offset).Limit(limit).Count(&count).Error

	if err != nil && errors.Is(err, gorm.ErrEmptySlice) {
		return 0, nil, err
	}
	return count, users, nil
}

func (u *users) Update(ctx context.Context, m *model.UserM) error {
	return u.db.WithContext(ctx).Save(m).Error
}
