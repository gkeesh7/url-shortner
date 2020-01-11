package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"url-shortner/storage/gormmodel"
)

type UrlStatsAggregateMockDao struct {
	mock.Mock
}

func (u *UrlStatsAggregateMockDao) NativeQuery(context context.Context, query string, arguments ...interface{}) ([]*gormmodel.URLStatsAggregate, error) {
	args := u.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*gormmodel.URLStatsAggregate), args.Error(1)
}
