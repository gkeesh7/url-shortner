package api

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"net/http"
	"url-shortner/api/dto"
	"url-shortner/common"
	"url-shortner/logic"
)

// RedirectRequest accepts a shortened URL and redirects it to the Long URL saved in the DataStore
func RedirectRequest(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ShortURLId := logic.ExtractURLIdFromRequest(r)
	if len(ShortURLId) == 0 {
		ErrorHandler(w, *dto.NewError(dto.RecordNotFound, "url_id not sent in request"))
		return
	}

	errCB := hystrix.Do(common.READ_FROM_DATABASE, func() error {
		redirectURL, errDataStore := logic.FindLongURLFromDataStore(ctx, ShortURLId)
		if errDataStore != nil {
			ErrorHandler(w, *errDataStore)
			return errDataStore
		}
		if len(redirectURL) == 0 {
			ErrorHandler(w, *dto.NewError(dto.RecordNotFound, "redirect URL is empty"))
			return nil
		}
		//Asynchronously push stats to the data store
		go logic.PushStatsToDataStore(ctx, ShortURLId)

		http.Redirect(w, r, redirectURL, http.StatusTemporaryRedirect)
		return nil
	}, nil)

	if errCB != nil {
		ErrorHandler(w, dto.Error{
			Code:    dto.InternalServerError,
			Message: errCB.Error(),
		})
	}
}
