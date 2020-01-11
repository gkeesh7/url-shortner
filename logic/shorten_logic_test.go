package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
	"url-shortner/api/dto"
	"url-shortner/storage/dao"
	"url-shortner/storage/dao/mocks"
)

func Test_ExtractShortenRequest_JsonDecodeFailure(t *testing.T) {

	//setup httpRequest With illegal body
	r, _ := http.NewRequest("POST", "http://0.0.0.0:8080/shorten", bytes.NewBuffer([]byte(`{`)))
	shortenRequest, err := ExtractShortenRequest(r)
	assert.NotNil(t, err)
	assert.Nil(t, shortenRequest)
	assert.Equal(t, dto.JsonError, err.Code)
}

func Test_ExtractShortenRequest_BadRequestEror(t *testing.T) {

	body, _ := json.Marshal(map[string]string{
		"url":        "",
		"request_id": "bar",
	})

	r, _ := http.NewRequest("POST", "http://0.0.0.0:8080/shorten", bytes.NewBuffer(body))
	shortenRequest, err := ExtractShortenRequest(r)
	assert.NotNil(t, err)
	assert.Nil(t, shortenRequest)
	assert.Equal(t, dto.IllegalRequest, err.Code)
}

func Test_ExtractShortenRequest_Succeeds(t *testing.T) {
	body, _ := json.Marshal(map[string]string{
		"url":        "http://www.google.com",
		"request_id": "bar",
	})
	r, _ := http.NewRequest("POST", "http://0.0.0.0:8080/shorten", bytes.NewBuffer(body))
	shortenRequest, err := ExtractShortenRequest(r)
	assert.Nil(t, err)
	assert.NotNil(t, shortenRequest)
	assert.Equal(t, "http://www.google.com", shortenRequest.URL)
	assert.Equal(t, "bar", shortenRequest.RequestID)
}

func Test_ShortenUrl_WhenDBErrors_ReturnsError(t *testing.T) {
	//Setup mock behaviour
	mockDao := mocks.RedirectionDaoMock{}
	mockDao.On("Save", mock.Anything, mock.Anything).Return(fmt.Errorf("DB Error While Saving!!!"))

	//save the old state
	save := dao.RedirectionGormImplObj

	//Assign mock behaviour to actual behaviour
	dao.RedirectionGormImplObj = &mockDao

	//Teardown or reset when done
	defer func() {
		dao.RedirectionGormImplObj = save
	}()

	shortenResponse, err := ShortenURL(context.Background(), dto.URLShortenRequest{
		RequestID: "jasdhfkjdshgasd",
		URL:       "www.google.com",
		Expiry:    nil,
	})

	assert.Nil(t, shortenResponse)
	assert.NotNil(t, err)
	assert.Equal(t, dto.DatabaseError, err.Code)
}

func Test_ShortenUrl_Successfully(t *testing.T) {
	//Setup mock behaviour
	mockDao := mocks.RedirectionDaoMock{}
	mockDao.On("Save", mock.Anything, mock.Anything).Return(nil)

	//save the old state
	save := dao.RedirectionGormImplObj

	//Assign mock behaviour to actual behaviour
	dao.RedirectionGormImplObj = &mockDao

	//Teardown or reset when done
	defer func() {
		dao.RedirectionGormImplObj = save
	}()

	//Call the actual function being tested
	shortenResponse, err := ShortenURL(context.Background(), dto.URLShortenRequest{
		RequestID: "jasdhfkjdshgasd",
		URL:       "www.google.com",
		Expiry:    nil,
	})

	//Assert the results
	assert.Nil(t, err)
	assert.NotNil(t, shortenResponse)
	assert.NotEmpty(t, shortenResponse.ShortUrl)
}
