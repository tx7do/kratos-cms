package service

import (
	"context"
	"kratos-blog/api/pagination"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	v1 "kratos-blog/api/admin/service/v1"
	contentV1 "kratos-blog/api/content/service/v1"
)

type LinkService struct {
	v1.LinkServiceHTTPServer

	linkClient contentV1.LinkServiceClient
	log        *log.Helper
}

func NewLinkService(logger log.Logger, linkClient contentV1.LinkServiceClient) *LinkService {
	l := log.NewHelper(log.With(logger, "module", "link/service/admin-service"))
	return &LinkService{
		log:        l,
		linkClient: linkClient,
	}
}

// ListLink 获取链接列表
func (s *LinkService) ListLink(ctx context.Context, req *pagination.PagingRequest) (*contentV1.ListLinkResponse, error) {
	return s.linkClient.ListLink(ctx, req)
}

// GetLink 获取链接数据
func (s *LinkService) GetLink(ctx context.Context, req *contentV1.GetLinkRequest) (*contentV1.Link, error) {
	return s.linkClient.GetLink(ctx, req)
}

// CreateLink 创建链接
func (s *LinkService) CreateLink(ctx context.Context, req *contentV1.CreateLinkRequest) (*contentV1.Link, error) {
	return s.linkClient.CreateLink(ctx, req)
}

// UpdateLink 更新链接
func (s *LinkService) UpdateLink(ctx context.Context, req *contentV1.UpdateLinkRequest) (*contentV1.Link, error) {
	return s.linkClient.UpdateLink(ctx, req)
}

// DeleteLink 删除链接
func (s *LinkService) DeleteLink(ctx context.Context, req *contentV1.DeleteLinkRequest) (*emptypb.Empty, error) {
	return s.linkClient.DeleteLink(ctx, req)
}
