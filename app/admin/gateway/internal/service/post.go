package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	v1 "kratos-blog/api/blog/v1"
	"kratos-blog/third_party/pagination"
)

type PostService struct {
	v1.UnimplementedPostServiceServer

	postClient v1.PostServiceClient
	log        *log.Helper
}

func NewPostService(logger log.Logger, postClient v1.PostServiceClient) *PostService {
	l := log.NewHelper(log.With(logger, "module", "service/post"))
	return &PostService{
		log:        l,
		postClient: postClient,
	}
}

// ListPost 获取帖子列表
func (s *PostService) ListPost(ctx context.Context, req *pagination.PagingRequest) (*v1.ListPostResponse, error) {
	return s.postClient.ListPost(ctx, req)
}

// GetPost 获取帖子数据
func (s *PostService) GetPost(ctx context.Context, req *v1.GetPostRequest) (*v1.Post, error) {
	return s.postClient.GetPost(ctx, req)
}

// CreatePost 创建帖子
func (s *PostService) CreatePost(ctx context.Context, req *v1.CreatePostRequest) (*v1.Post, error) {
	return s.postClient.CreatePost(ctx, req)
}

// UpdatePost 更新帖子
func (s *PostService) UpdatePost(ctx context.Context, req *v1.UpdatePostRequest) (*v1.Post, error) {
	return s.postClient.UpdatePost(ctx, req)
}

// DeletePost 删除帖子
func (s *PostService) DeletePost(ctx context.Context, req *v1.DeletePostRequest) (*emptypb.Empty, error) {
	return s.postClient.DeletePost(ctx, req)
}
