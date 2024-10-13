package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-cms/app/core/service/internal/data"

	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	contentV1 "kratos-cms/api/gen/go/content/service/v1"
)

type LinkService struct {
	contentV1.UnimplementedLinkServiceServer

	uc  *data.LinkRepo
	log *log.Helper
}

func NewLinkService(logger log.Logger, uc *data.LinkRepo) *LinkService {
	l := log.NewHelper(log.With(logger, "module", "link/service/core-service"))
	return &LinkService{
		log: l,
		uc:  uc,
	}
}

// ListLink 获取链接列表
func (s *LinkService) ListLink(ctx context.Context, req *pagination.PagingRequest) (*contentV1.ListLinkResponse, error) {
	return s.uc.List(ctx, req)
}

// GetLink 获取链接数据
func (s *LinkService) GetLink(ctx context.Context, req *contentV1.GetLinkRequest) (*contentV1.Link, error) {
	return s.uc.Get(ctx, req)
}

// CreateLink 创建链接
func (s *LinkService) CreateLink(ctx context.Context, req *contentV1.CreateLinkRequest) (*contentV1.Link, error) {
	return s.uc.Create(ctx, req)
}

// UpdateLink 更新链接
func (s *LinkService) UpdateLink(ctx context.Context, req *contentV1.UpdateLinkRequest) (*contentV1.Link, error) {
	return s.uc.Update(ctx, req)
}

// DeleteLink 删除链接
func (s *LinkService) DeleteLink(ctx context.Context, req *contentV1.DeleteLinkRequest) (*emptypb.Empty, error) {
	_, err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
