package usecase_test

import (
	"context"
	"testing"

	"go_docker/domain"
	"go_docker/domain/mocks"
	usecase "go_docker/user/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetById(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockUser := domain.User{
		ID:   1,
		Name: "name1",
	}

	t.Run("success", func (t *testing.T) {
		mockUserRepo.On("GetById", mock.Anything, mock.AnythingOfType("int64")).Return(mockUser, nil)

		userUsecase := usecase.NewUserUsecase(mockUserRepo)
		user, err := userUsecase.GetById(context.Background(), 1)

		assert.NoError(t, err)
		assert.Equal(t, mockUser, user)
	})
}