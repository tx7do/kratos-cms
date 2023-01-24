package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/tx7do/kratos-blog/blog-backend/gen/api/go/content/service/v1"
	"github.com/tx7do/kratos-blog/blog-backend/gen/api/go/pagination"
)

type LinkRepo interface {
	List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListLinkResponse, error)
	Get(ctx context.Context, req *v1.GetLinkRequest) (*v1.Link, error)
	Create(ctx context.Context, req *v1.CreateLinkRequest) (*v1.Link, error)
	Update(ctx context.Context, req *v1.UpdateLinkRequest) (*v1.Link, error)
	Delete(ctx context.Context, req *v1.DeleteLinkRequest) (bool, error)
}

type LinkUseCase struct {
	repo LinkRepo
	log  *log.Helper
}

func NewLinkUseCase(repo LinkRepo, logger log.Logger) *LinkUseCase {
	l := log.NewHelper(log.With(logger, "module", "link/usecase/content-service"))
	return &LinkUseCase{
		repo: repo,
		log:  l,
	}
}

func (uc *LinkUseCase) List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListLinkResponse, error) {
	return uc.repo.List(ctx, req)
}

func (uc *LinkUseCase) Get(ctx context.Context, req *v1.GetLinkRequest) (*v1.Link, error) {
	return uc.repo.Get(ctx, req)
}

func (uc *LinkUseCase) Create(ctx context.Context, req *v1.CreateLinkRequest) (*v1.Link, error) {
	return uc.repo.Create(ctx, req)
}

func (uc *LinkUseCase) Update(ctx context.Context, req *v1.UpdateLinkRequest) (*v1.Link, error) {
	return uc.repo.Update(ctx, req)
}

func (uc *LinkUseCase) Delete(ctx context.Context, req *v1.DeleteLinkRequest) (bool, error) {
	return uc.repo.Delete(ctx, req)
}
