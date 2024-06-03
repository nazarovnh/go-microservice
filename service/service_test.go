package service

import (
	"testing"

	"github.com/bruc3mackenzi3/microservice-demo/model"
	"github.com/bruc3mackenzi3/microservice-demo/repository"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	var testID uint = 1234
	
	testUser := model.User{Name: "Abraham"}

	mockRepo := &repository.MockRepository{}
	userService := NewUserService(mockRepo)

	t.Run("successful response", func(t *testing.T) {
		mockRepo.On("SelectUser", testID).Return(&testUser, nil).Once()

		actualUser, err := userService.GetUser(testID)

		assert.Nil(t, err)
		assert.Equal(t, &testUser, actualUser)
	})

	t.Run("error response", func(t *testing.T) {
		// testError := model.NewNotFoundError(fmt.Sprintf("User %d not found", testID))
		mockRepo.On("SelectUser", testID).Return(nil, model.ErrUserNotFound).Once()

		actualUser, err := userService.GetUser(testID)

		assert.Equal(t, model.ErrUserNotFound, err)
		assert.Nil(t, actualUser)
	})
}
