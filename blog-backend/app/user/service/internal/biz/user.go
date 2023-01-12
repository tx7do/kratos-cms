package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "kratos-blog/api/user/service/v1"
)

type UserRepo interface {
	Create(ctx context.Context, req *v1.RegisterRequest) (*v1.User, error)
	Update(ctx context.Context, req *v1.UpdateUserRequest) (*v1.User, error)
	Delete(ctx context.Context, req *v1.DeleteUserRequest) (bool, error)

	Get(ctx context.Context, req uint32) (*v1.User, error)
	GetUserByUserName(ctx context.Context, userName string) (*v1.User, error)

	VerifyPassword(ctx context.Context, req *v1.LoginRequest) (bool, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	l := log.NewHelper(log.With(logger, "module", "user/usecase/user-service"))
	return &UserUseCase{
		repo: repo,
		log:  l,
	}
}

func (uc *UserUseCase) Get(ctx context.Context, req uint32) (*v1.User, error) {
	user, err := uc.repo.Get(ctx, req)
	user.Password = nil
	return user, err
}

func (uc *UserUseCase) GetUserByUserName(ctx context.Context, userName string) (*v1.User, error) {
	return uc.repo.GetUserByUserName(ctx, userName)
}

func (uc *UserUseCase) Create(ctx context.Context, req *v1.RegisterRequest) (*v1.User, error) {
	return uc.repo.Create(ctx, req)
}

func (uc *UserUseCase) Update(ctx context.Context, req *v1.UpdateUserRequest) (*v1.User, error) {
	return uc.repo.Update(ctx, req)
}

func (uc *UserUseCase) Delete(ctx context.Context, req *v1.DeleteUserRequest) (bool, error) {
	return uc.repo.Delete(ctx, req)
}

func (uc *UserUseCase) VerifyPassword(ctx context.Context, req *v1.LoginRequest) (bool, error) {
	return uc.repo.VerifyPassword(ctx, req)
}
