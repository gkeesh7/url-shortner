package api

import (
	"context"
	"net/http"
	"url-shortner/logic"
)

func Sort(w http.ResponseWriter, r *http.Request) {
	requestDto, errExtracting := logic.ExtractSortRequest(r)
	if errExtracting != nil {
		ErrorHandler(w, *errExtracting)
		return
	}
	response, errSorting := logic.SortArray(context.Background(), requestDto)
	if errSorting != nil {
		ErrorHandler(w, *errSorting)
		return
	}
	SuccessHandler(w, response)
}
