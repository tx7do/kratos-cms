package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
	v1 "kratos-blog/api/comment/service/v1"
	"kratos-blog/api/pagination"
	"kratos-blog/app/comment/service/internal/biz"
)

type CommentService struct {
	v1.UnimplementedCommentServiceServer

	uc  *biz.CommentUseCase
	log *log.Helper
}

func NewCommentService(logger log.Logger, uc *biz.CommentUseCase) *CommentService {
	l := log.NewHelper(log.With(logger, "module", "service/comment/comment-service"))
	return &CommentService{
		log: l,
		uc:  uc,
	}
}

// ListComment 获取留言列表
func (s *CommentService) ListComment(ctx context.Context, req *pagination.PagingRequest) (*v1.ListCommentResponse, error) {
	return s.uc.List(ctx, req)
}

// GetComment 获取留言数据
func (s *CommentService) GetComment(ctx context.Context, req *v1.GetCommentRequest) (*v1.Comment, error) {
	return s.uc.Get(ctx, req)
}

// CreateComment 创建留言
func (s *CommentService) CreateComment(ctx context.Context, req *v1.CreateCommentRequest) (*v1.Comment, error) {
	return s.uc.Create(ctx, req)
}

// UpdateComment 更新留言
func (s *CommentService) UpdateComment(ctx context.Context, req *v1.UpdateCommentRequest) (*v1.Comment, error) {
	return s.uc.Update(ctx, req)
}

// DeleteComment 删除留言
func (s *CommentService) DeleteComment(ctx context.Context, req *v1.DeleteCommentRequest) (*emptypb.Empty, error) {
	_, err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
