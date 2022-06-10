package service

import (
	"errors"
	"test-rod/mock/mock_repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_UserInvalidUserError(t *testing.T) {
	mockRepo := &mock_repository.Repository{}

	mockRepo.On("GetUser", mock.Anything).Return("root", nil)

	service := NewService(mockRepo)
	err := service.ValidUser("")

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "invalid user")
}

func Test_UserNotFoundError(t *testing.T) {
	mockRepo := &mock_repository.Repository{}

	mockRepo.On("GetUser", mock.Anything).Return("", errors.New("not found"))

	service := NewService(mockRepo)
	err := service.ValidUser("fp")

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "not found")
}

func Test_UserSuccess(t *testing.T) {
	mockRepo := &mock_repository.Repository{}

	mockRepo.On("GetUser", mock.Anything).Return("root", nil)

	service := NewService(mockRepo)
	err := service.ValidUser("fparente")

	assert.Nil(t, err)
}
