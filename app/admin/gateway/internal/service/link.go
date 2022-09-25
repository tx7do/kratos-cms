package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	v1 "kratos-blog/api/blog/v1"
	"kratos-blog/third_party/pagination"
)

type LinkService struct {
	v1.UnimplementedLinkServiceServer

	linkClient v1.LinkServiceClient
	log        *log.Helper
}

func NewLinkService(logger log.Logger, linkClient v1.LinkServiceClient) *LinkService {
	l := log.NewHelper(log.With(logger, "module", "service/link"))
	return &LinkService{
		log:        l,
		linkClient: linkClient,
	}
}

// ListLink 获取链接列表
func (s *LinkService) ListLink(ctx context.Context, req *pagination.PagingRequest) (*v1.ListLinkResponse, error) {
	return s.linkClient.ListLink(ctx, req)
}

// GetLink 获取链接数据
func (s *LinkService) GetLink(ctx context.Context, req *v1.GetLinkRequest) (*v1.Link, error) {
	return s.linkClient.GetLink(ctx, req)
}

// CreateLink 创建链接
func (s *LinkService) CreateLink(ctx context.Context, req *v1.CreateLinkRequest) (*v1.Link, error) {
	return s.linkClient.CreateLink(ctx, req)
}

// UpdateLink 更新链接
func (s *LinkService) UpdateLink(ctx context.Context, req *v1.UpdateLinkRequest) (*v1.Link, error) {
	return s.linkClient.UpdateLink(ctx, req)
}

// DeleteLink 删除链接
func (s *LinkService) DeleteLink(ctx context.Context, req *v1.DeleteLinkRequest) (*emptypb.Empty, error) {
	return s.linkClient.DeleteLink(ctx, req)
}
