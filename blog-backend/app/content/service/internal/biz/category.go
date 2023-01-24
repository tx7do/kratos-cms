package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/tx7do/kratos-blog/blog-backend/gen/api/go/content/service/v1"
	"github.com/tx7do/kratos-blog/blog-backend/gen/api/go/pagination"
)

type CategoryRepo interface {
	List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListCategoryResponse, error)
	Get(ctx context.Context, req *v1.GetCategoryRequest) (*v1.Category, error)
	Create(ctx context.Context, req *v1.CreateCategoryRequest) (*v1.Category, error)
	Update(ctx context.Context, req *v1.UpdateCategoryRequest) (*v1.Category, error)
	Delete(ctx context.Context, req *v1.DeleteCategoryRequest) (bool, error)
}

type CategoryUseCase struct {
	repo CategoryRepo
	log  *log.Helper
}

func NewCategoryUseCase(repo CategoryRepo, logger log.Logger) *CategoryUseCase {
	l := log.NewHelper(log.With(logger, "module", "category/usecase/content-service"))
	return &CategoryUseCase{
		repo: repo,
		log:  l,
	}
}

func (uc *CategoryUseCase) List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListCategoryResponse, error) {
	return uc.repo.List(ctx, req)
}

func (uc *CategoryUseCase) Get(ctx context.Context, req *v1.GetCategoryRequest) (*v1.Category, error) {
	return uc.repo.Get(ctx, req)
}

func (uc *CategoryUseCase) Create(ctx context.Context, req *v1.CreateCategoryRequest) (*v1.Category, error) {
	return uc.repo.Create(ctx, req)
}

func (uc *CategoryUseCase) Update(ctx context.Context, req *v1.UpdateCategoryRequest) (*v1.Category, error) {
	return uc.repo.Update(ctx, req)
}

func (uc *CategoryUseCase) Delete(ctx context.Context, req *v1.DeleteCategoryRequest) (bool, error) {
	return uc.repo.Delete(ctx, req)
}
