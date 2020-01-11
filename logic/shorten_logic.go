package logic

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
	"url-shortner/api/dto"
	"url-shortner/common"
	"url-shortner/storage/dao"
	"url-shortner/storage/gormmodel"
	"url-shortner/utils/randomutils"
)

func ExtractShortenRequest(r *http.Request) (*dto.URLShortenRequest, *dto.Error) {
	decoder := json.NewDecoder(r.Body)
	var shortenRequest dto.URLShortenRequest
	errDecoding := decoder.Decode(&shortenRequest)
	if errDecoding != nil {
		log.Printf("Error while decoding the json %v", errDecoding.Error())
		return nil, dto.NewError(dto.JsonError, errDecoding.Error())
	}
	if len(shortenRequest.URL) == 0 {
		log.Print("No URL provided for shortening")
		return nil, dto.NewError(dto.IllegalRequest, "cannot shorten empty URL")
	}
	return &shortenRequest, nil
}

func ShortenURL(ctx context.Context, request dto.URLShortenRequest) (*dto.URLShortenResponse, *dto.Error) {
	expiryTime := time.Now().Add(24 * time.Hour)
	if request.Expiry != nil {
		expiryTime = *request.Expiry
	}
	redirectionDBObject := createObjectForStorage(request, expiryTime)
	errSaving := dao.RedirectionGormImplObj.Save(ctx, &redirectionDBObject)
	if errSaving != nil {
		log.Printf("Error saving the redirection object in DB %v", errSaving.Error())
		return nil, dto.NewError(dto.DatabaseError, "Cannot save data "+errSaving.Error())
	}
	resp := generateShortenResponseObject(request, redirectionDBObject.UrlId, expiryTime)
	return &resp, nil
}

func createObjectForStorage(request dto.URLShortenRequest, expiryTime time.Time) gormmodel.Redirection {
	return gormmodel.Redirection{
		UrlId:  randomutils.RandomString(10),
		URL:    request.URL,
		Expiry: expiryTime,
	}
}

func generateShortenResponseObject(request dto.URLShortenRequest, urlID string, expiryTime time.Time) dto.URLShortenResponse {
	resp := dto.URLShortenResponse{
		RequestID:   request.RequestID,
		ShortUrl:    common.URL_REDIRECT_PREFIX + urlID,
		RedirectUrl: request.URL,
		Expiry:      expiryTime,
	}
	return resp
}
