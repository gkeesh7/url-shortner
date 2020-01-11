package logic

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"url-shortner/api/dto"
	"url-shortner/storage/dao"
	"url-shortner/storage/dao/mocks"
	"url-shortner/storage/gormmodel"
)

func Test_WhenPushStatsCalledAndDatabaseGivesError_ItErrors(t *testing.T) {
	//Setup mock behaviour
	mockDao := mocks.UrlStatsMock{}
	mockDao.On("Save", mock.Anything, mock.Anything).Return(fmt.Errorf("dB error!!!"))

	//Save the existing behaviour
	save := dao.UrlStatsGormImplObj

	//Replace the mock behaviour
	dao.UrlStatsGormImplObj = &mockDao

	//Reset or TearDown after the test is executed
	defer func() {
		dao.UrlStatsGormImplObj = save
	}()

	//Call the function to be testted
	err := PushStatsToDataStore(context.Background(), "abcdef")

	//Assert that the response is not nil
	assert.NotNil(t, err)
}

func Test_WhenPushStatsCalled_ItSucceeds(t *testing.T) {
	//Setup mock behaviour
	mockDao := mocks.UrlStatsMock{}
	mockDao.On("Save", mock.Anything, mock.Anything).Return(nil)

	//Save the existing behaviour
	save := dao.UrlStatsGormImplObj

	//Replace the mock behaviour
	dao.UrlStatsGormImplObj = &mockDao

	//Reset or TearDown after the test is executed
	defer func() {
		dao.UrlStatsGormImplObj = save
	}()

	//Call the function to be tested
	err := PushStatsToDataStore(context.Background(), "abcdef")

	//Assert that the response is not nil
	assert.Nil(t, err)
}

func Test_GetUrlOpenStats_WhenDBErrors_ReturnsAnError(t *testing.T) {
	//Setup mock behaviour
	mockDao := mocks.UrlStatsAggregateMockDao{}
	mockDao.On("NativeQuery", mock.Anything, mock.Anything).Return(nil, fmt.Errorf("DB Error!!!"))

	//Save the existing behaviour
	save := dao.UrlStatsAggregateDaoImplObj

	//Replace the mock Behaviour
	dao.UrlStatsAggregateDaoImplObj = &mockDao

	//Reset or tearDown after the test is executed
	defer func() {
		dao.UrlStatsAggregateDaoImplObj = save
	}()

	//Call the function to be tested
	statsResponse, err := GetUrlOpenStats(context.Background())

	//Assert the responses
	assert.Nil(t, statsResponse)
	assert.NotNil(t, err)
	assert.Equal(t, dto.DatabaseError, err.Code)

}

func Test_GetUrlOpenStats_WhenSucceeds(t *testing.T) {

	//Setup mock behaviour
	mockDao := mocks.UrlStatsAggregateMockDao{}
	mockDao.On("NativeQuery", mock.Anything, mock.Anything).Return([]*gormmodel.URLStatsAggregate{
		{
			UrlId:    "abcdef",
			SumCount: 4,
		},
		{
			UrlId:    "foobar",
			SumCount: 2,
		},
	}, nil)

	//Save the existing behaviour
	save := dao.UrlStatsAggregateDaoImplObj

	//Replace the mock Behaviour
	dao.UrlStatsAggregateDaoImplObj = &mockDao

	//Reset or tearDown after the test is executed
	defer func() {
		dao.UrlStatsAggregateDaoImplObj = save
	}()

	//Call the function to be tested
	statsResponse, err := GetUrlOpenStats(context.Background())

	//Assert the responses
	assert.Nil(t, err)
	assert.NotNil(t, statsResponse)
	assert.Equal(t, 2, len(statsResponse.URLStats))
	assert.True(t, statsResponse.URLStats[0].Count >= statsResponse.URLStats[1].Count)
}
