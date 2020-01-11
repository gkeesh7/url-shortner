package dao

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
	"url-shortner/config/database"
	"url-shortner/storage/gormmodel"
)

type UrlStatsGormImpl struct {
}

var UrlStatsGormImplObj UrlStatsDao

func init() {
	UrlStatsGormImplObj = &UrlStatsGormImpl{}
}

func (u *UrlStatsGormImpl) Save(context context.Context, urlStats *gormmodel.URLStats) error {
	db := database.GetDbHandle()
	if db == nil {
		return fmt.Errorf("failed to get database handle")
	}
	urlStats.CreatedAt = time.Now().UTC()
	urlStats.UpdatedAt = urlStats.CreatedAt
	if err := db.Create(urlStats).Error; err != nil {
		return err
	}
	return nil
}

func (u *UrlStatsGormImpl) Update(context context.Context, db *gorm.DB, urlStats *gormmodel.URLStats, updatedUrlStats *gormmodel.URLStats) error {
	updatedUrlStats.UpdatedAt = time.Now().UTC()
	err := db.Model(urlStats).Updates(updatedUrlStats).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *UrlStatsGormImpl) FindOne(context context.Context, condition interface{}) (*gormmodel.URLStats, error) {
	db := database.GetDbHandle()
	if db == nil {
		return nil, fmt.Errorf("failed to cast gorm session")
	}
	var response = &gormmodel.URLStats{}
	res := db.First(response, condition)
	if res.RecordNotFound() {
		return nil, nil
	}
	if res.Error != nil {
		return nil, res.Error
	}
	return response, nil
}

func (u *UrlStatsGormImpl) FindAll(context context.Context, condition interface{}) (*[]gormmodel.URLStats, error) {
	db := database.GetDbHandle()
	if db == nil {
		return nil, fmt.Errorf("failed to cast gorm session")
	}
	var response []gormmodel.URLStats
	res := db.Find(&response, condition)
	if res.RecordNotFound() {
		return nil, nil
	}
	if res.Error != nil {
		return nil, res.Error
	}
	return &response, nil
}
