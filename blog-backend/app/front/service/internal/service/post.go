package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-cms/gen/api/go/common/pagination"
	contentV1 "kratos-cms/gen/api/go/content/service/v1"
	v1 "kratos-cms/gen/api/go/front/service/v1"
)

type PostService struct {
	v1.PostServiceHTTPServer

	postClient contentV1.PostServiceClient
	log        *log.Helper
}

func NewPostService(logger log.Logger, postClient contentV1.PostServiceClient) *PostService {
	l := log.NewHelper(log.With(logger, "module", "post/service/front-service"))
	return &PostService{
		log:        l,
		postClient: postClient,
	}
}

// ListPost 获取帖子列表
func (s *PostService) ListPost(ctx context.Context, req *pagination.PagingRequest) (*contentV1.ListPostResponse, error) {
	return s.postClient.ListPost(ctx, req)
}

// GetPost 获取帖子数据
func (s *PostService) GetPost(ctx context.Context, req *contentV1.GetPostRequest) (*contentV1.Post, error) {
	return s.postClient.GetPost(ctx, req)
}

// CreatePost 创建帖子
func (s *PostService) CreatePost(ctx context.Context, req *contentV1.CreatePostRequest) (*contentV1.Post, error) {
	return s.postClient.CreatePost(ctx, req)
}

// UpdatePost 更新帖子
func (s *PostService) UpdatePost(ctx context.Context, req *contentV1.UpdatePostRequest) (*contentV1.Post, error) {
	return s.postClient.UpdatePost(ctx, req)
}

// DeletePost 删除帖子
func (s *PostService) DeletePost(ctx context.Context, req *contentV1.DeletePostRequest) (*emptypb.Empty, error) {
	return s.postClient.DeletePost(ctx, req)
}
