package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"kratos-blog/gen/api/go/common/pagination"
	"kratos-blog/gen/api/go/content/service/v1"
)

type PostRepo interface {
	List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListPostResponse, error)
	Get(ctx context.Context, req *v1.GetPostRequest) (*v1.Post, error)
	Create(ctx context.Context, req *v1.CreatePostRequest) (*v1.Post, error)
	Update(ctx context.Context, req *v1.UpdatePostRequest) (*v1.Post, error)
	Delete(ctx context.Context, req *v1.DeletePostRequest) (bool, error)
}

type PostUseCase struct {
	repo PostRepo
	log  *log.Helper
}

func NewPostUseCase(repo PostRepo, logger log.Logger) *PostUseCase {
	l := log.NewHelper(log.With(logger, "module", "post/usecase/content-service"))
	return &PostUseCase{
		repo: repo,
		log:  l,
	}
}

func (uc *PostUseCase) List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListPostResponse, error) {
	return uc.repo.List(ctx, req)
}

func (uc *PostUseCase) Get(ctx context.Context, req *v1.GetPostRequest) (*v1.Post, error) {
	return uc.repo.Get(ctx, req)
}

func (uc *PostUseCase) Create(ctx context.Context, req *v1.CreatePostRequest) (*v1.Post, error) {
	return uc.repo.Create(ctx, req)
}

func (uc *PostUseCase) Update(ctx context.Context, req *v1.UpdatePostRequest) (*v1.Post, error) {
	return uc.repo.Update(ctx, req)
}

func (uc *PostUseCase) Delete(ctx context.Context, req *v1.DeletePostRequest) (bool, error) {
	return uc.repo.Delete(ctx, req)
}
