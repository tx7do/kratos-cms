package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	v1 "kratos-blog/gen/api/go/admin/service/v1"
	contentV1 "kratos-blog/gen/api/go/content/service/v1"
	"kratos-blog/gen/api/go/pagination"
)

type TagService struct {
	v1.TagServiceHTTPServer

	tagClient contentV1.TagServiceClient
	log       *log.Helper
}

func NewTagService(logger log.Logger, tagClient contentV1.TagServiceClient) *TagService {
	l := log.NewHelper(log.With(logger, "module", "tag/service/admin-service"))
	return &TagService{
		log:       l,
		tagClient: tagClient,
	}
}

// ListTag 获取标签列表
func (s *TagService) ListTag(ctx context.Context, req *pagination.PagingRequest) (*contentV1.ListTagResponse, error) {
	return s.tagClient.ListTag(ctx, req)
}

// GetTag 获取标签数据
func (s *TagService) GetTag(ctx context.Context, req *contentV1.GetTagRequest) (*contentV1.Tag, error) {
	return s.tagClient.GetTag(ctx, req)
}

// CreateTag 创建标签
func (s *TagService) CreateTag(ctx context.Context, req *contentV1.CreateTagRequest) (*contentV1.Tag, error) {
	return s.tagClient.CreateTag(ctx, req)
}

// UpdateTag 更新标签
func (s *TagService) UpdateTag(ctx context.Context, req *contentV1.UpdateTagRequest) (*contentV1.Tag, error) {
	return s.tagClient.UpdateTag(ctx, req)
}

// DeleteTag 删除标签
func (s *TagService) DeleteTag(ctx context.Context, req *contentV1.DeleteTagRequest) (*emptypb.Empty, error) {
	return s.tagClient.DeleteTag(ctx, req)
}
