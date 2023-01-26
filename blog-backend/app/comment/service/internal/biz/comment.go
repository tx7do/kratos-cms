package biz

import (
	"context"
	"kratos-blog/gen/api/go/pagination"

	"github.com/go-kratos/kratos/v2/log"

	v1 "kratos-blog/gen/api/go/comment/service/v1"
)

type CommentRepo interface {
	List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListCommentResponse, error)
	Get(ctx context.Context, req *v1.GetCommentRequest) (*v1.Comment, error)
	Create(ctx context.Context, req *v1.CreateCommentRequest) (*v1.Comment, error)
	Update(ctx context.Context, req *v1.UpdateCommentRequest) (*v1.Comment, error)
	Delete(ctx context.Context, req *v1.DeleteCommentRequest) (bool, error)
}

type CommentUseCase struct {
	repo CommentRepo
	log  *log.Helper
}

func NewCommentUseCase(repo CommentRepo, logger log.Logger) *CommentUseCase {
	l := log.NewHelper(log.With(logger, "module", "comment/usecase/comment-service"))
	return &CommentUseCase{
		repo: repo,
		log:  l,
	}
}

func (uc *CommentUseCase) List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListCommentResponse, error) {
	return uc.repo.List(ctx, req)
}

func (uc *CommentUseCase) Get(ctx context.Context, req *v1.GetCommentRequest) (*v1.Comment, error) {
	return uc.repo.Get(ctx, req)
}

func (uc *CommentUseCase) Create(ctx context.Context, req *v1.CreateCommentRequest) (*v1.Comment, error) {
	return uc.repo.Create(ctx, req)
}

func (uc *CommentUseCase) Update(ctx context.Context, req *v1.UpdateCommentRequest) (*v1.Comment, error) {
	return uc.repo.Update(ctx, req)
}

func (uc *CommentUseCase) Delete(ctx context.Context, req *v1.DeleteCommentRequest) (bool, error) {
	return uc.repo.Delete(ctx, req)
}
