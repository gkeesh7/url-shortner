package logic

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"url-shortner/storage/dao"
	"url-shortner/storage/dao/mocks"
)

func Test_DeleteExpiredUrls_WhenDatabaseErrors_ReturnsAnError(t *testing.T) {
	//Setup mock behaviour
	mockDao := mocks.RedirectionDaoMock{}
	mockDao.On("NativeExec", mock.Anything, mock.Anything).Return(fmt.Errorf("there was a database error!!!"))

	//Save existing behaviour
	save := dao.RedirectionGormImplObj

	//Assign mock behaviour to Actual behaviour
	dao.RedirectionGormImplObj = &mockDao

	//TearDown Reset to old behaviour
	defer func() {
		dao.RedirectionGormImplObj = save
	}()

	//Call the function to be tested
	err := DeleteExpiredURLs(context.Background())

	//Write assertions
	assert.NotNil(t, err)
	assert.Equal(t, "there was a database error!!!", err.Error())
}

func Test_DeleteExpiredUrls_WhenQueryExecutesSuccesfully_ReturnsNil(t *testing.T) {
	//Setup mock behaviour
	mockDao := mocks.RedirectionDaoMock{}
	mockDao.On("NativeExec", mock.Anything, mock.Anything).Return(nil)

	//Save existing behaviour
	save := dao.RedirectionGormImplObj

	//Assign mock behaviour to Actual behaviour
	dao.RedirectionGormImplObj = &mockDao

	//TearDown Reset to old behaviour
	defer func() {
		dao.RedirectionGormImplObj = save
	}()

	//Call the function to be tested
	err := DeleteExpiredURLs(context.Background())

	//Write assertions
	assert.Nil(t, err)
}
