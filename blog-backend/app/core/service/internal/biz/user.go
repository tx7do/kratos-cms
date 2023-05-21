package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"kratos-cms/gen/api/go/common/pagination"
	v1 "kratos-cms/gen/api/go/user/service/v1"
)

type UserRepo interface {
	Create(ctx context.Context, req *v1.CreateUserRequest) (*v1.User, error)
	Update(ctx context.Context, req *v1.UpdateUserRequest) (*v1.User, error)
	Delete(ctx context.Context, req *v1.DeleteUserRequest) (bool, error)
	List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListUserResponse, error)
	Get(ctx context.Context, req *v1.GetUserRequest) (*v1.User, error)

	GetUserByUserName(ctx context.Context, userName string) (*v1.User, error)

	VerifyPassword(ctx context.Context, req *v1.VerifyPasswordRequest) (bool, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	l := log.NewHelper(log.With(logger, "module", "user/usecase/core-service"))
	return &UserUseCase{
		repo: repo,
		log:  l,
	}
}

func (uc *UserUseCase) Get(ctx context.Context, req *v1.GetUserRequest) (*v1.User, error) {
	user, err := uc.repo.Get(ctx, req)
	if user != nil {
		user.Password = nil
	}
	return user, err
}

func (uc *UserUseCase) List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListUserResponse, error) {
	return uc.repo.List(ctx, req)
}

func (uc *UserUseCase) Create(ctx context.Context, req *v1.CreateUserRequest) (*v1.User, error) {
	return uc.repo.Create(ctx, req)
}

func (uc *UserUseCase) Update(ctx context.Context, req *v1.UpdateUserRequest) (*v1.User, error) {
	return uc.repo.Update(ctx, req)
}

func (uc *UserUseCase) Delete(ctx context.Context, req *v1.DeleteUserRequest) (bool, error) {
	return uc.repo.Delete(ctx, req)
}

func (uc *UserUseCase) GetUserByUserName(ctx context.Context, userName string) (*v1.User, error) {
	return uc.repo.GetUserByUserName(ctx, userName)
}

func (uc *UserUseCase) VerifyPassword(ctx context.Context, req *v1.VerifyPasswordRequest) (bool, error) {
	return uc.repo.VerifyPassword(ctx, req)
}
