package apitest

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
	"url-shortner/api/dto"
)

//Tests if redirection request succeeds or not
func Test_GetLongURLFromShort(t *testing.T) {
	//Call the http API
	resp, err := http.Get("http://localhost:8080/redirect/test")
	//Assert the results
	assert.NotNil(t, resp)
	assert.Nil(t, err)
}

//Test if redirection succeeds the response is a non null object
func Test_GetLongURLFromShort_AssertOnResponseBody(t *testing.T) {
	//Call the Http API
	resp, err := http.Get("http://localhost:8080/redirect/test")
	body, errRead := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	//Assert the results
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.NotNil(t, body)
	assert.Nil(t, errRead)
}

//Test what happens when the URLid is not found in the database
func Test_GetLongURLFromShort_NotFoundInDatabase(t *testing.T) {
	//Call the endpoint
	resp, err := http.Get("http://localhost:8080/redirect/SurelyNotGoingToFindThis")
	body, errRead := ioutil.ReadAll(resp.Body)
	notFoundResponse := dto.Error{}
	errUnmarshal := json.Unmarshal(body, &notFoundResponse)
	defer resp.Body.Close()

	//Assert on the results
	assert.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
	assert.Nil(t, err)
	assert.NotNil(t, body)
	assert.Nil(t, errRead)
	assert.Nil(t, errUnmarshal)
	assert.Equal(t, dto.RecordNotFound, notFoundResponse.Code)
}

//Test when URL visit request succeeds
func Test_GetURLVisitStats(t *testing.T) {
	//Call the http Api
	resp, err := http.Get("http://localhost:8080/stats")
	body, errRead := ioutil.ReadAll(resp.Body)
	statsResponse := dto.StatsResponse{}
	errUnmarshal := json.Unmarshal(body, &statsResponse)
	defer resp.Body.Close()

	//Assert the results
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Nil(t, errRead)
	assert.Nil(t, errUnmarshal)
	assert.NotEmpty(t, statsResponse.Message)
}

//Test when URL shortening returns a successful response
func Test_ShortenURLsuccessfully(t *testing.T) {
	//Create the Post request
	requestBody, _ := json.Marshal(map[string]string{
		"url":        "https://www.google.com/search?q=do+a+barrel+roll",
		"request_id": "asdlfjhlaksdjffsajkflkjghjasfflkg",
	})
	//Call the HTTP endpoint
	resp, err := http.Post("http://0.0.0.0:8080/shorten", "application/json", bytes.NewBuffer(requestBody))
	body, errRead := ioutil.ReadAll(resp.Body)
	shortenResponse := dto.URLShortenResponse{}
	errUnmarshal := json.Unmarshal(body, &shortenResponse)
	defer resp.Body.Close()

	//Assert the results
	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Nil(t, errRead)
	assert.Nil(t, errUnmarshal)
	assert.NotEmpty(t, shortenResponse.ShortUrl)
	assert.Equal(t, shortenResponse.RedirectUrl, "https://www.google.com/search?q=do+a+barrel+roll")
	assert.Equal(t, "asdlfjhlaksdjffsajkflkjghjasfflkg", shortenResponse.RequestID)
}
