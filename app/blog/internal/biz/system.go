package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	v1 "kratos-blog/api/blog/v1"
	"kratos-blog/third_party/pagination"
)

type SystemRepo interface {
	List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListSystemResponse, error)
	Get(ctx context.Context, req *v1.GetSystemRequest) (*v1.System, error)
	Create(ctx context.Context, req *v1.CreateSystemRequest) (*v1.System, error)
	Update(ctx context.Context, req *v1.UpdateSystemRequest) (*v1.System, error)
	Delete(ctx context.Context, req *v1.DeleteSystemRequest) (bool, error)
}

type SystemUseCase struct {
	repo SystemRepo
	log  *log.Helper
}

func NewSystemUseCase(repo SystemRepo, logger log.Logger) *SystemUseCase {
	l := log.NewHelper(log.With(logger, "module", "system/usecase"))
	return &SystemUseCase{
		repo: repo,
		log:  l,
	}
}

func (uc *SystemUseCase) List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListSystemResponse, error) {
	return uc.repo.List(ctx, req)
}

func (uc *SystemUseCase) Get(ctx context.Context, req *v1.GetSystemRequest) (*v1.System, error) {
	return uc.repo.Get(ctx, req)
}

func (uc *SystemUseCase) Create(ctx context.Context, req *v1.CreateSystemRequest) (*v1.System, error) {
	return uc.repo.Create(ctx, req)
}

func (uc *SystemUseCase) Update(ctx context.Context, req *v1.UpdateSystemRequest) (*v1.System, error) {
	return uc.repo.Update(ctx, req)
}

func (uc *SystemUseCase) Delete(ctx context.Context, req *v1.DeleteSystemRequest) (bool, error) {
	return uc.repo.Delete(ctx, req)
}
