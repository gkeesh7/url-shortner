package dao

import (
	"context"
	"github.com/jinzhu/gorm"
	"url-shortner/storage/gormmodel"
)

type RedirectionDao interface {
	Save(context context.Context, redirect *gormmodel.Redirection) error
	Update(context context.Context, db *gorm.DB, redirect *gormmodel.Redirection, updatedRedirect *gormmodel.Redirection) error
	FindOne(context context.Context, condition interface{}) (*gormmodel.Redirection, error)
	FindAll(context context.Context, condition interface{}) (*[]gormmodel.Redirection, error)
	NativeExec(context context.Context, query string) error
}
