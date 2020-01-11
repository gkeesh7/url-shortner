package logic

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"url-shortner/api/dto"
	"url-shortner/storage/dao"
)

func ExtractURLIdFromRequest(r *http.Request) string {
	vars := mux.Vars(r)
	ShortURLId := vars["url_id"]
	log.Printf("Received URL redirection request for Id %v", ShortURLId)
	return ShortURLId
}

func FindLongURLFromDataStore(ctx context.Context, ShortURLId string) (string, *dto.Error) {
	conditions := createFilters(ShortURLId)
	redirect, err := dao.RedirectionGormImplObj.FindOne(ctx, conditions)
	if err != nil {
		log.Printf("Database Error received %v", err.Error())
		return "", dto.GenerateError(dto.DatabaseError, err.Error())
	}
	if redirect == nil {
		log.Print("No data found for url_id = " + ShortURLId)
		return "", dto.GenerateError(dto.RecordNotFound, "No data found for url_id = "+ShortURLId)
	}
	return redirect.URL, nil
}

func createFilters(ShortURLId string) map[string]interface{} {
	conditions := make(map[string]interface{})
	conditions["url_id"] = ShortURLId
	return conditions
}
