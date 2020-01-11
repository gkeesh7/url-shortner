package dao

import (
	"context"
	"fmt"
	"url-shortner/config/database"
	"url-shortner/storage/gormmodel"
)

type UrlStatsAggregateDaoImpl struct {
}

var UrlStatsAggregateDaoImplObj UrlStatsAggregateDao

func init() {
	UrlStatsAggregateDaoImplObj = &UrlStatsAggregateDaoImpl{}
}

func (dao *UrlStatsAggregateDaoImpl) NativeQuery(context context.Context, query string, args ...interface{}) ([]*gormmodel.URLStatsAggregate, error) {
	db := database.GetDbHandle()
	if db == nil {
		return nil, fmt.Errorf("failed to cast gorm session")
	}
	var response = &[]*gormmodel.URLStatsAggregate{}
	res := db.Raw(query, args...).Find(response)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RecordNotFound() {
		return nil, nil
	}
	return *response, nil
}
