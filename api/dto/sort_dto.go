package dto

//SortRequest is the JSON object whose array in the request will be sorted
type SortRequest struct {
	RequestId string `json:"request_id"`
	Array     []int  `json:"array"`
	SortOrder string `json:"sort_order"`
}

//SortResponse is the JSON object whose array will  be sorted
type SortResponse struct {
	RequestId string `json:"request_id"`
	Array     []int  `json:"array"`
}
