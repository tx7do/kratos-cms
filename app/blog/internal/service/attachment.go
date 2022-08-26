package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
	v1 "kratos-blog/api/blog/v1"
	"kratos-blog/app/blog/internal/biz"
	"kratos-blog/third_party/pagination"
)

type AttachmentService struct {
	v1.UnimplementedAttachmentServiceServer

	uc  *biz.AttachmentUseCase
	log *log.Helper
}

func NewAttachmentService(logger log.Logger, uc *biz.AttachmentUseCase) *AttachmentService {
	l := log.NewHelper(log.With(logger, "module", "service/attachment"))
	return &AttachmentService{
		log: l,
		uc:  uc,
	}
}

// ListAttachment 获取链接列表
func (s *AttachmentService) ListAttachment(ctx context.Context, req *pagination.PagingRequest) (*v1.ListAttachmentResponse, error) {
	return s.uc.List(ctx, req)
}

// GetAttachment 获取链接数据
func (s *AttachmentService) GetAttachment(ctx context.Context, req *v1.GetAttachmentRequest) (*v1.Attachment, error) {
	return s.uc.Get(ctx, req)
}

// CreateAttachment 创建链接
func (s *AttachmentService) CreateAttachment(ctx context.Context, req *v1.CreateAttachmentRequest) (*v1.Attachment, error) {
	return s.uc.Create(ctx, req)
}

// UpdateAttachment 更新链接
func (s *AttachmentService) UpdateAttachment(ctx context.Context, req *v1.UpdateAttachmentRequest) (*v1.Attachment, error) {
	return s.uc.Update(ctx, req)
}

// DeleteAttachment 删除链接
func (s *AttachmentService) DeleteAttachment(ctx context.Context, req *v1.DeleteAttachmentRequest) (*emptypb.Empty, error) {
	_, err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
