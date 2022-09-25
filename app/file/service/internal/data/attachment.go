package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"kratos-blog/app/file/service/internal/biz"

	v1 "kratos-blog/api/file/service/v1"
	"kratos-blog/third_party/pagination"
)

var _ biz.AttachmentRepo = (*AttachmentRepo)(nil)

type AttachmentRepo struct {
	data *Data
	log  *log.Helper
}

func NewAttachmentRepo(data *Data, logger log.Logger) biz.AttachmentRepo {
	l := log.NewHelper(log.With(logger, "module", "attachment/repo"))
	return &AttachmentRepo{
		data: data,
		log:  l,
	}
}

func (r *AttachmentRepo) List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListAttachmentResponse, error) {
	return nil, nil
}

func (r *AttachmentRepo) Get(ctx context.Context, req *v1.GetAttachmentRequest) (*v1.Attachment, error) {
	return nil, nil
}

func (r *AttachmentRepo) Create(ctx context.Context, req *v1.CreateAttachmentRequest) (*v1.Attachment, error) {
	return nil, nil
}

func (r *AttachmentRepo) Update(ctx context.Context, req *v1.UpdateAttachmentRequest) (*v1.Attachment, error) {
	return nil, nil
}

func (r *AttachmentRepo) Delete(ctx context.Context, req *v1.DeleteAttachmentRequest) (bool, error) {
	return false, nil
}
