package mocks

import (
	"context"

	"go_docker/domain"

	mock "github.com/stretchr/testify/mock"
)



type UserRepository struct {
	mock.Mock
}

func (_m *UserRepository) GetById(ctx context.Context, id int64) (domain.User, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(context.Context, int64) domain.User); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.User)
	}
	return r0, ret.Error(1)
}