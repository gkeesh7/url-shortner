package dto

type StatsResponse struct {
	Message  string        `json:"message"`
	URLStats []UrlHitCount `json:"url_stats"`
}

type UrlHitCount struct {
	URL   string `json:"url"`
	Count int    `json:"count"`
}
