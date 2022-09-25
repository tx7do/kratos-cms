package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
	v1 "kratos-blog/api/blog/v1"
	"kratos-blog/third_party/pagination"
)

type TagService struct {
	v1.UnimplementedTagServiceServer

	tagClient v1.TagServiceClient
	log       *log.Helper
}

func NewTagService(logger log.Logger, tagClient v1.TagServiceClient) *TagService {
	l := log.NewHelper(log.With(logger, "module", "service/tag"))
	return &TagService{
		log:       l,
		tagClient: tagClient,
	}
}

// ListTag 获取标签列表
func (s *TagService) ListTag(ctx context.Context, req *pagination.PagingRequest) (*v1.ListTagResponse, error) {
	return s.tagClient.ListTag(ctx, req)
}

// GetTag 获取标签数据
func (s *TagService) GetTag(ctx context.Context, req *v1.GetTagRequest) (*v1.Tag, error) {
	return s.tagClient.GetTag(ctx, req)
}

// CreateTag 创建标签
func (s *TagService) CreateTag(ctx context.Context, req *v1.CreateTagRequest) (*v1.Tag, error) {
	return s.tagClient.CreateTag(ctx, req)
}

// UpdateTag 更新标签
func (s *TagService) UpdateTag(ctx context.Context, req *v1.UpdateTagRequest) (*v1.Tag, error) {
	return s.tagClient.UpdateTag(ctx, req)
}

// DeleteTag 删除标签
func (s *TagService) DeleteTag(ctx context.Context, req *v1.DeleteTagRequest) (*emptypb.Empty, error) {
	return s.tagClient.DeleteTag(ctx, req)
}
