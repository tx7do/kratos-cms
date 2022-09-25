package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	v1 "kratos-blog/api/file/service/v1"
	"kratos-blog/third_party/pagination"
)

type AttachmentRepo interface {
	List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListAttachmentResponse, error)
	Get(ctx context.Context, req *v1.GetAttachmentRequest) (*v1.Attachment, error)
	Create(ctx context.Context, req *v1.CreateAttachmentRequest) (*v1.Attachment, error)
	Update(ctx context.Context, req *v1.UpdateAttachmentRequest) (*v1.Attachment, error)
	Delete(ctx context.Context, req *v1.DeleteAttachmentRequest) (bool, error)
}

type AttachmentUseCase struct {
	repo AttachmentRepo
	log  *log.Helper
}

func NewAttachmentUseCase(repo AttachmentRepo, logger log.Logger) *AttachmentUseCase {
	l := log.NewHelper(log.With(logger, "module", "attachment/usecase"))
	return &AttachmentUseCase{
		repo: repo,
		log:  l,
	}
}

func (uc *AttachmentUseCase) List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListAttachmentResponse, error) {
	return uc.repo.List(ctx, req)
}

func (uc *AttachmentUseCase) Get(ctx context.Context, req *v1.GetAttachmentRequest) (*v1.Attachment, error) {
	return uc.repo.Get(ctx, req)
}

func (uc *AttachmentUseCase) Create(ctx context.Context, req *v1.CreateAttachmentRequest) (*v1.Attachment, error) {
	return uc.repo.Create(ctx, req)
}

func (uc *AttachmentUseCase) Update(ctx context.Context, req *v1.UpdateAttachmentRequest) (*v1.Attachment, error) {
	return uc.repo.Update(ctx, req)
}

func (uc *AttachmentUseCase) Delete(ctx context.Context, req *v1.DeleteAttachmentRequest) (bool, error) {
	return uc.repo.Delete(ctx, req)
}
