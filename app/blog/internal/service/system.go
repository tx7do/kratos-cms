package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
	v1 "kratos-blog/api/blog/v1"
	"kratos-blog/app/blog/internal/biz"
	"kratos-blog/third_party/pagination"
)

type SystemService struct {
	v1.UnimplementedSystemServiceServer

	uc  *biz.SystemUseCase
	log *log.Helper
}

func NewSystemService(logger log.Logger, uc *biz.SystemUseCase) *SystemService {
	l := log.NewHelper(log.With(logger, "module", "service/system"))
	return &SystemService{
		log: l,
		uc:  uc,
	}
}

// ListSystem 获取系统配置列表
func (s *SystemService) ListSystem(ctx context.Context, req *pagination.PagingRequest) (*v1.ListSystemResponse, error) {
	return s.uc.List(ctx, req)
}

// GetSystem 获取系统配置数据
func (s *SystemService) GetSystem(ctx context.Context, req *v1.GetSystemRequest) (*v1.System, error) {
	return s.uc.Get(ctx, req)
}

// CreateSystem 创建系统配置
func (s *SystemService) CreateSystem(ctx context.Context, req *v1.CreateSystemRequest) (*v1.System, error) {
	return s.uc.Create(ctx, req)
}

// UpdateSystem 更新系统配置
func (s *SystemService) UpdateSystem(ctx context.Context, req *v1.UpdateSystemRequest) (*v1.System, error) {
	return s.uc.Update(ctx, req)
}

// DeleteSystem 删除系统配置
func (s *SystemService) DeleteSystem(ctx context.Context, req *v1.DeleteSystemRequest) (*emptypb.Empty, error) {
	_, err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
