package dto

import "time"

// URLShortenRequest request for the Shorten API
type URLShortenRequest struct {
	RequestID string     `json:"request_id"`
	URL       string     `json:"url"`
	Expiry    *time.Time `json:"expiry"`
}


// URLShortenResponse response for the Shorten API
type URLShortenResponse struct {
	RequestID   string    `json:"request_id"`
	ShortUrl    string    `json:"short_url"`
	RedirectUrl string    `json:"redirect_url"`
	Expiry      time.Time `json:"expiry"`
}
