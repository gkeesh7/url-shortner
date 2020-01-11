package api

import (
	"context"
	"log"
	"net/http"
	"url-shortner/logic"
)

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	//Extract Request Dto from http Request
	requestDto, errExtracting := logic.ExtractShortenRequest(r)
	if errExtracting != nil {
		ErrorHandler(w, *errExtracting)
		return
	}
	log.Printf("The extracted Request Dto %v", *requestDto)
	//Store the URL to be redirected along with it's unique id in DB
	response, errShortening := logic.ShortenURL(ctx, *requestDto)
	if errShortening != nil {
		ErrorHandler(w, *errShortening)
		return
	}
	//Create Successful Response
	SuccessHandler(w, response)
}
