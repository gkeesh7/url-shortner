package logic

import (
	"context"
	"encoding/json"
	"net/http"
	"sort"
	"url-shortner/api/dto"
	"url-shortner/common"
)

func ExtractSortRequest(r *http.Request) (dto.SortRequest, *dto.Error) {
	decoder := json.NewDecoder(r.Body)
	var sortRequest dto.SortRequest
	err := decoder.Decode(&sortRequest)
	if err != nil {
		return dto.SortRequest{}, dto.NewError(dto.JsonError, err.Error())
	}
	if (sortRequest.SortOrder != common.ASCENDING) && (sortRequest.SortOrder != common.DESCENDING) {
		return dto.SortRequest{}, dto.NewError(dto.IllegalRequest, "no sorting order specified")
	}
	return sortRequest, nil
}

func SortArray(ctx context.Context, req dto.SortRequest) (dto.SortResponse, *dto.Error) {
	arr := req.Array
	switch req.SortOrder {
	case common.ASCENDING:
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})
		break
	case common.DESCENDING:
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] > arr[j]
		})
		break
	}
	return dto.SortResponse{
		RequestId: req.RequestId,
		Array:     arr,
	}, nil
}
