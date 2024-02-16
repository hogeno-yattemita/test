package usecase
import (
	"context"
	"go_docker/domain"
)



type UserUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(u domain.UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepo: u,
	}
}

func (u *UserUsecase) GetById(ctx context.Context, id int64) (res domain.User, err error) {
	return u.userRepo.GetById(ctx, id)
}