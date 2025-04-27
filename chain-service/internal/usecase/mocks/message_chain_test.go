package usecase_test

import (
	"errors"
	"testing"

	gomock "go.uber.org/mock/gomock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/szaluzhanskaya/Innopolis/chain-service/internal/entity"
	"github.com/szaluzhanskaya/Innopolis/chain-service/internal/usecase"
)

func TestMockMessageChainRepository_CreateMessageChain_Success(t *testing.T) {
	// Create controller for mock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock for MessageChainRepository
	mockRepo := NewMockMessageChainRepository(ctrl)

	// Prepare test data
	chain := &entity.MessageChain{
		ID:    1,
		Title: "Test Title",
	}

	// Setting expectations
	mockRepo.EXPECT().
		CreateMessageChain(chain).
		Return(nil)

	// Run CreateMessageChain
	err := mockRepo.CreateMessageChain(chain)

	// Checking the results
	assert.NoError(t, err)
}

func TestMockMessageChainRepository_CreateMessageChain_Error(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockMessageChainRepository(ctrl)

	chain := &entity.MessageChain{
		ID:    1,
		Title: "Test Title",
	}
	expectedError := errors.New("failed to create message chain")

	mockRepo.EXPECT().
		CreateMessageChain(chain).
		Return(expectedError)

	err := mockRepo.CreateMessageChain(chain)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
}

func TestMockMessageChainUsecase_CreateMessageChain(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := NewMockMessageChainUsecase(ctrl)

	userID := 1
	title := "Test Title"
	expectedChain := &entity.MessageChain{
		ID:    1,
		Title: title,
	}

	mockUsecase.EXPECT().
		CreateMessageChain(userID, title).
		Return(expectedChain, nil)

	result, err := mockUsecase.CreateMessageChain(userID, title)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedChain.ID, result.ID)
	assert.Equal(t, expectedChain.Title, result.Title)
}

func TestMockMessageChainUsecase_CreateMessageChain_Error(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := NewMockMessageChainUsecase(ctrl)

	userID := 1
	title := "Test Title"
	expectedError := errors.New("failed to create message chain")

	mockUsecase.EXPECT().
		CreateMessageChain(userID, title).
		Return(nil, expectedError)

	result, err := mockUsecase.CreateMessageChain(userID, title)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, expectedError, err)
}

func TestMockMessageChainRepository_DeleteMessageChain_Success(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepository := NewMockMessageChainRepository(ctrl)

	uuid := "test string"

	mockRepository.EXPECT().DeleteMessageChain(uuid).Return(nil)

	actual := mockRepository.DeleteMessageChain(uuid)
	require.NoError(t, actual)
}

func TestMockMessageChainRepository_DeleteMessageChain_Error(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepository := NewMockMessageChainRepository(ctrl)

	errForTest := errors.New("test error")
	uuid := "test string"

	mockRepository.EXPECT().DeleteMessageChain(uuid).Return(errForTest)

	actual := mockRepository.DeleteMessageChain(uuid)
	require.ErrorIs(t, actual, errForTest)
}

func TestMockMessageChainUsecase_DeleteMessageChain_Success(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepository := NewMockMessageChainRepository(ctrl)
	service := usecase.New(mockRepository)

	uuid := "test string"

	mockRepository.EXPECT().DeleteMessageChain(uuid).Return(nil)

	actual := service.DeleteMessageChain(uuid)
	require.NoError(t, actual)
}

func TestMockMessageChainUsecase_DeleteMessageChain_Error(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepository := NewMockMessageChainRepository(ctrl)
	service := usecase.New(mockRepository)

	errForTest := errors.New("test error")
	uuid := "test string"

	mockRepository.EXPECT().DeleteMessageChain(uuid).Return(errForTest)

	actual := service.DeleteMessageChain(uuid)
	require.ErrorIs(t, actual, errForTest)
}
