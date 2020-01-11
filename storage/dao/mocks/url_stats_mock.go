package mocks

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/mock"
	"url-shortner/storage/gormmodel"
)

type UrlStatsMock struct {
	mock.Mock
}

func (u *UrlStatsMock) Save(context context.Context, urlStats *gormmodel.URLStats) error {
	args := u.Called()
	return args.Error(0)
}

func (u *UrlStatsMock) Update(context context.Context, db *gorm.DB, urlStats *gormmodel.URLStats, updatedUrlStats *gormmodel.URLStats) error {
	args := u.Called()
	return args.Error(0)
}

func (u *UrlStatsMock) FindOne(context context.Context, condition interface{}) (*gormmodel.URLStats, error) {
	args := u.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*gormmodel.URLStats), args.Error(1)
}

func (u *UrlStatsMock) FindAll(context context.Context, condition interface{}) (*[]gormmodel.URLStats, error) {
	args := u.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*[]gormmodel.URLStats), args.Error(1)
}
