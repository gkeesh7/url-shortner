package dao

import (
	"context"
	"url-shortner/storage/gormmodel"
)

type UrlStatsAggregateDao interface {
	NativeQuery(context context.Context, query string, args ...interface{}) ([]*gormmodel.URLStatsAggregate, error)
}
