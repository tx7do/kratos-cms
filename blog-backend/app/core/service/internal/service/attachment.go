package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-cms/app/core/service/internal/biz"

	"kratos-cms/gen/api/go/common/pagination"
	v1 "kratos-cms/gen/api/go/file/service/v1"
)

type AttachmentService struct {
	v1.UnimplementedAttachmentServiceServer

	uc  *biz.AttachmentUseCase
	log *log.Helper
}

func NewAttachmentService(logger log.Logger, uc *biz.AttachmentUseCase) *AttachmentService {
	l := log.NewHelper(log.With(logger, "module", "attachment/service/core-service"))
	return &AttachmentService{
		log: l,
		uc:  uc,
	}
}

// ListAttachment 获取附件列表
func (s *AttachmentService) ListAttachment(ctx context.Context, req *pagination.PagingRequest) (*v1.ListAttachmentResponse, error) {
	return s.uc.List(ctx, req)
}

// GetAttachment 获取附件数据
func (s *AttachmentService) GetAttachment(ctx context.Context, req *v1.GetAttachmentRequest) (*v1.Attachment, error) {
	return s.uc.Get(ctx, req)
}

// CreateAttachment 创建附件
func (s *AttachmentService) CreateAttachment(ctx context.Context, req *v1.CreateAttachmentRequest) (*v1.Attachment, error) {
	return s.uc.Create(ctx, req)
}

// UpdateAttachment 更新附件
func (s *AttachmentService) UpdateAttachment(ctx context.Context, req *v1.UpdateAttachmentRequest) (*v1.Attachment, error) {
	return s.uc.Update(ctx, req)
}

// DeleteAttachment 删除附件
func (s *AttachmentService) DeleteAttachment(ctx context.Context, req *v1.DeleteAttachmentRequest) (*emptypb.Empty, error) {
	_, err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
