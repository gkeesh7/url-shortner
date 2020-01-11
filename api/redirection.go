package api

import (
	"context"
	"net/http"
	"url-shortner/api/dto"
	"url-shortner/logic"
)

func RedirectRequest(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ShortURLId := logic.ExtractURLIdFromRequest(r)
	if len(ShortURLId) == 0 {
		ErrorHandler(w, *dto.NewError(dto.RecordNotFound, "url_id not sent in request"))
		return
	}
	redirectURL, errDataStore := logic.FindLongURLFromDataStore(ctx, ShortURLId)
	if errDataStore != nil {
		ErrorHandler(w, *errDataStore)
		return
	}
	if len(redirectURL) == 0 {
		ErrorHandler(w, *dto.NewError(dto.RecordNotFound, "redirect URL is empty"))
		return
	}

	//Asynchronously push stats to the data store
	go logic.PushStatsToDataStore(ctx, ShortURLId)

	http.Redirect(w, r, redirectURL, http.StatusTemporaryRedirect)
}
