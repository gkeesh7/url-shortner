package dao

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
	"url-shortner/config/database"
	"url-shortner/storage/gormmodel"
)

type RedirectionGormImpl struct {
}

var RedirectionGormImplObj RedirectionDao

func init() {
	RedirectionGormImplObj = &RedirectionGormImpl{}
}

func (dao *RedirectionGormImpl) Save(context context.Context, redirect *gormmodel.Redirection) error {
	db := database.GetDbHandle()
	if db == nil {
		return fmt.Errorf("failed to get database handle")
	}
	redirect.CreatedAt = time.Now().UTC()
	redirect.UpdatedAt = redirect.CreatedAt
	if err := db.Create(redirect).Error; err != nil {
		return err
	}
	return nil
}

func (dao *RedirectionGormImpl) Update(context context.Context, db *gorm.DB, redirect *gormmodel.Redirection, updatedRedirect *gormmodel.Redirection) error {
	updatedRedirect.UpdatedAt = time.Now().UTC()
	err := db.Model(redirect).Updates(updatedRedirect).Error
	if err != nil {
		return err
	}
	return nil
}

func (dao *RedirectionGormImpl) FindOne(context context.Context, condition interface{}) (*gormmodel.Redirection, error) {
	db := database.GetDbHandle()
	if db == nil {
		return nil, fmt.Errorf("failed to cast gorm session")
	}
	var response = &gormmodel.Redirection{}
	res := db.First(response, condition)
	if res.RecordNotFound() {
		return nil, nil
	}
	if res.Error != nil {
		return nil, res.Error
	}
	return response, nil
}

func (dao *RedirectionGormImpl) FindAll(context context.Context, condition interface{}) (*[]gormmodel.Redirection, error) {
	db := database.GetDbHandle()
	if db == nil {
		return nil, fmt.Errorf("failed to cast gorm session")
	}
	var response []gormmodel.Redirection
	res := db.Find(&response, condition)
	if res.RecordNotFound() {
		return nil, nil
	}
	if res.Error != nil {
		return nil, res.Error
	}
	return &response, nil
}

func (dao *RedirectionGormImpl) NativeExec(ctx context.Context, query string) error {
	db := database.GetDbHandle()
	if db == nil {
		return fmt.Errorf("failed to cast gorm session")
	}
	res := db.Exec(query)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
