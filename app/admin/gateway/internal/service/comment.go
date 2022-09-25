package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	v1 "kratos-blog/api/comment/service/v1"
	"kratos-blog/third_party/pagination"
)

type CommentService struct {
	v1.UnimplementedCommentServiceServer

	commentClient v1.CommentServiceClient
	log           *log.Helper
}

func NewCommentService(logger log.Logger, commentClient v1.CommentServiceClient) *CommentService {
	l := log.NewHelper(log.With(logger, "module", "service/comment"))
	return &CommentService{
		log:           l,
		commentClient: commentClient,
	}
}

// ListComment 获取留言列表
func (s *CommentService) ListComment(ctx context.Context, req *pagination.PagingRequest) (*v1.ListCommentResponse, error) {
	return s.commentClient.ListComment(ctx, req)
}

// GetComment 获取留言数据
func (s *CommentService) GetComment(ctx context.Context, req *v1.GetCommentRequest) (*v1.Comment, error) {
	return s.commentClient.GetComment(ctx, req)
}

// CreateComment 创建留言
func (s *CommentService) CreateComment(ctx context.Context, req *v1.CreateCommentRequest) (*v1.Comment, error) {
	return s.commentClient.CreateComment(ctx, req)
}

// UpdateComment 更新留言
func (s *CommentService) UpdateComment(ctx context.Context, req *v1.UpdateCommentRequest) (*v1.Comment, error) {
	return s.commentClient.UpdateComment(ctx, req)
}

// DeleteComment 删除留言
func (s *CommentService) DeleteComment(ctx context.Context, req *v1.DeleteCommentRequest) (*emptypb.Empty, error) {
	return s.commentClient.DeleteComment(ctx, req)
}
