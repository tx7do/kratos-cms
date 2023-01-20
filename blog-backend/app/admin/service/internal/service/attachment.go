package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	v1 "kratos-blog/gen/api/go/admin/service/v1"
	fileV1 "kratos-blog/gen/api/go/file/service/v1"
	"kratos-blog/gen/api/go/pagination"
)

type AttachmentService struct {
	v1.AttachmentServiceHTTPServer

	attachmentClient fileV1.AttachmentServiceClient
	log              *log.Helper
}

func NewAttachmentService(logger log.Logger, attachmentClient fileV1.AttachmentServiceClient) *AttachmentService {
	l := log.NewHelper(log.With(logger, "module", "attachment/service/admin-service"))
	return &AttachmentService{
		log:              l,
		attachmentClient: attachmentClient,
	}
}

// ListAttachment 获取附件列表
func (s *AttachmentService) ListAttachment(ctx context.Context, req *pagination.PagingRequest) (*fileV1.ListAttachmentResponse, error) {
	return s.attachmentClient.ListAttachment(ctx, req)
}

// GetAttachment 获取附件数据
func (s *AttachmentService) GetAttachment(ctx context.Context, req *fileV1.GetAttachmentRequest) (*fileV1.Attachment, error) {
	return s.attachmentClient.GetAttachment(ctx, req)
}

// CreateAttachment 创建附件
func (s *AttachmentService) CreateAttachment(ctx context.Context, req *fileV1.CreateAttachmentRequest) (*fileV1.Attachment, error) {
	return s.attachmentClient.CreateAttachment(ctx, req)
}

// UpdateAttachment 更新附件
func (s *AttachmentService) UpdateAttachment(ctx context.Context, req *fileV1.UpdateAttachmentRequest) (*fileV1.Attachment, error) {
	return s.attachmentClient.UpdateAttachment(ctx, req)
}

// DeleteAttachment 删除附件
func (s *AttachmentService) DeleteAttachment(ctx context.Context, req *fileV1.DeleteAttachmentRequest) (*emptypb.Empty, error) {
	return s.attachmentClient.DeleteAttachment(ctx, req)
}
