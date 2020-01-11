package api

import (
	"context"
	"log"
	"net/http"
	"url-shortner/logic"
)

func Stats(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	aggreatedStatsResponse, statsErr := logic.GetUrlOpenStats(ctx)
	if statsErr != nil {
		log.Printf("received stats error %v", statsErr.Error())
		ErrorHandler(w, *statsErr)
	}
	SuccessHandler(w, *aggreatedStatsResponse)
}
