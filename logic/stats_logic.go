package logic

import (
	"context"
	"log"
	"time"
	"url-shortner/api/dto"
	"url-shortner/common"
	"url-shortner/storage/dao"
	"url-shortner/storage/gormmodel"
)

func PushStatsToDataStore(ctx context.Context, ShortURLId string) error {
	statsObject := createStatsDBObject(ShortURLId)
	errSavingStat := dao.UrlStatsGormImplObj.Save(ctx, &statsObject)
	if errSavingStat != nil {
		log.Printf("Unable to push stats to dataStore %v", errSavingStat.Error())
	}
	return errSavingStat
}

func createStatsDBObject(ShortURLId string) gormmodel.URLStats {
	return gormmodel.URLStats{
		UrlId:     ShortURLId,
		Count:     1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func GetUrlOpenStats(ctx context.Context) (*dto.StatsResponse, *dto.Error) {
	resp := initStatsResponseObject()
	UrlStatsAggregatData, errStats := dao.UrlStatsAggregateDaoImplObj.NativeQuery(ctx, "SELECT url_id,sum(count) FROM url_stats WHERE (created_at>=(NOW()-INTERVAL 1 DAY)) GROUP BY(url_id) ORDER by sum(`count`) DESC LIMIT 10")
	if errStats != nil {
		log.Printf("Database error happened while finding stats " + errStats.Error())
		return nil, dto.NewError(dto.DatabaseError, "Database error while fetching stats")
	}
	assignStatstoResponse(UrlStatsAggregatData, &resp)
	return &resp, nil
}

func initStatsResponseObject() dto.StatsResponse {
	return dto.StatsResponse{
		Message:  "",
		URLStats: make([]dto.UrlHitCount, 0),
	}
}

func assignStatstoResponse(UrlStatsAggregateData []*gormmodel.URLStatsAggregate, response *dto.StatsResponse) {
	if len(UrlStatsAggregateData) == 0 {
		response.Message = common.NO_STATS_DATA
		return
	}
	response.Message = common.URL_STATS_MESSAGE
	for _, element := range UrlStatsAggregateData {
		response.URLStats = append(response.URLStats, dto.UrlHitCount{
			URL:   common.URL_REDIRECT_PREFIX + element.UrlId,
			Count: element.SumCount,
		})
	}
}
