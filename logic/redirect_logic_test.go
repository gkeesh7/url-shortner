package logic

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
	"url-shortner/api/dto"
	"url-shortner/storage/dao"
	"url-shortner/storage/dao/mocks"
	"url-shortner/storage/gormmodel"
)

func Test_CreateFilters_ReturnsASingularFilter(t *testing.T) {
	mp := createFilters("abcdef")
	assert.Equal(t, 1, len(mp))
	assert.Equal(t, "abcdef", mp["url_id"])
}

func Test_FindLongURL_WhenDatabaseErrors_ReturnAnError(t *testing.T) {

	//set up mock behaviour
	mockDao := mocks.RedirectionDaoMock{}
	mockDao.On("FindOne", mock.Anything, mock.Anything).Return(nil, fmt.Errorf("DB error!!!"))

	//Save existing behaviour
	save := dao.RedirectionGormImplObj

	//Assign mock behaviour to Actual behaviour
	dao.RedirectionGormImplObj = &mockDao

	//TearDown Reset to old behaviour after the test
	defer func() {
		dao.RedirectionGormImplObj = save
	}()

	//call the function to be tested
	str, err := FindLongURLFromDataStore(context.Background(), "abcdef")

	//assert the results
	assert.Equal(t, dto.DatabaseError, err.Code)
	assert.Empty(t, str)
}

func Test_FindLongURL_WhenDatabaseReturnsNil_ReturnsRecordNotFound(t *testing.T) {

	//set up mock behaviour
	mockDao := mocks.RedirectionDaoMock{}
	mockDao.On("FindOne", mock.Anything, mock.Anything).Return(nil, nil)

	//Save existing behaviour
	save := dao.RedirectionGormImplObj

	//Assign mock behaviour to Actual behaviour
	dao.RedirectionGormImplObj = &mockDao

	//TearDown Reset to old behaviour after the test
	defer func() {
		dao.RedirectionGormImplObj = save
	}()

	//Call the function to be tested
	str, err := FindLongURLFromDataStore(context.Background(), "abcdef")

	//assert the results
	assert.Equal(t, dto.RecordNotFound, err.Code)
	assert.Empty(t, str)
}

func Test_FindLongURL_WhenEntryFound_ReturnsTheURL(t *testing.T) {

	//set up mock behaviour
	mockDao := mocks.RedirectionDaoMock{}
	mockDao.On("FindOne", mock.Anything, mock.Anything).Return(&gormmodel.Redirection{
		ID:        int32(1),
		UrlId:     "abcdef",
		URL:       "www.google.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Expiry:    time.Now(),
	}, nil)

	//Save existing behaviour
	save := dao.RedirectionGormImplObj

	//Assign mock behaviour to Actual behaviour
	dao.RedirectionGormImplObj = &mockDao

	//TearDown Reset to old behaviour after the test
	defer func() {
		dao.RedirectionGormImplObj = save
	}()

	//Call the function to be tested
	str, err := FindLongURLFromDataStore(context.Background(), "abcdef")

	//assert the results
	assert.Nil(t, err)
	assert.Equal(t, str, "www.google.com")
}
