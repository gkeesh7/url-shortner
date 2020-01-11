package dao

import (
	"context"
	"github.com/jinzhu/gorm"
	"url-shortner/storage/gormmodel"
)

type UrlStatsDao interface {
	Save(context context.Context, urlStats *gormmodel.URLStats) error
	Update(context context.Context, db *gorm.DB, urlStats *gormmodel.URLStats, updatedUrlStats *gormmodel.URLStats) error
	FindOne(context context.Context, condition interface{}) (*gormmodel.URLStats, error)
	FindAll(context context.Context, condition interface{}) (*[]gormmodel.URLStats, error)
}
