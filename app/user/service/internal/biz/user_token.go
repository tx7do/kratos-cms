package biz

import (
	"context"
	v1 "kratos-blog/api/user/service/v1"
)

type UserTokenRepo interface {
	GetToken(ctx context.Context, userId uint32) string
	ValidateToken(ctx context.Context, token string) error
	GenerateToken(ctx context.Context, user *v1.User) (string, error)
	RemoveToken(ctx context.Context, userId uint32) error
}

type UserTokenUseCase struct {
	repo UserTokenRepo
}

func NewUserAuthUseCase(repo UserTokenRepo) *UserTokenUseCase {
	return &UserTokenUseCase{
		repo: repo,
	}
}

func (uc *UserTokenUseCase) GenerateToken(ctx context.Context, user *v1.User) (string, error) {
	return uc.repo.GenerateToken(ctx, user)
}

func (uc *UserTokenUseCase) GetToken(ctx context.Context, userId uint32) string {
	return uc.repo.GetToken(ctx, userId)
}

func (uc *UserTokenUseCase) ValidateToken(ctx context.Context, token string) error {
	return uc.repo.ValidateToken(ctx, token)
}

func (uc *UserTokenUseCase) RemoveToken(ctx context.Context, userId uint32) error {
	return uc.repo.RemoveToken(ctx, userId)
}
