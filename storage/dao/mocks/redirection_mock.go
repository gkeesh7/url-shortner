package mocks

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/mock"
	"url-shortner/storage/gormmodel"
)

type RedirectionDaoMock struct {
	mock.Mock
}

func (r *RedirectionDaoMock) Save(context context.Context, redirect *gormmodel.Redirection) error {
	args := r.Called()
	return args.Error(0)
}

func (r *RedirectionDaoMock) Update(context context.Context, db *gorm.DB, redirect *gormmodel.Redirection, updatedRedirect *gormmodel.Redirection) error {
	args := r.Called()
	return args.Error(0)
}

func (r *RedirectionDaoMock) FindOne(context context.Context, condition interface{}) (*gormmodel.Redirection, error) {
	args := r.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*gormmodel.Redirection), args.Error(1)
}

func (r *RedirectionDaoMock) FindAll(context context.Context, condition interface{}) (*[]gormmodel.Redirection, error) {
	args := r.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*[]gormmodel.Redirection), args.Error(1)
}

func (r *RedirectionDaoMock) NativeExec(context context.Context, query string) error {
	args := r.Called()
	return args.Error(0)
}
