package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-utils/trans"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-cms/app/core/service/internal/data"

	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	v1 "kratos-cms/gen/api/go/user/service/v1"
)

type UserService struct {
	v1.UnimplementedUserServiceServer

	uc  *data.UserRepo
	log *log.Helper
}

func NewUserService(logger log.Logger, uc *data.UserRepo) *UserService {
	l := log.NewHelper(log.With(logger, "module", "user/service/core-service"))
	return &UserService{
		log: l,
		uc:  uc,
	}
}

func (s *UserService) ListUser(ctx context.Context, req *pagination.PagingRequest) (*v1.ListUserResponse, error) {
	resp, err := s.uc.List(ctx, req)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(resp.Items); i++ {
		resp.Items[i].Password = trans.String("")
	}

	return resp, nil
}

func (s *UserService) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.User, error) {
	resp, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	resp.Password = trans.String("")

	return resp, nil
}

func (s *UserService) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*v1.User, error) {
	resp, err := s.uc.Create(ctx, req)
	if err != nil {
		// s.log.Info(err)
		return nil, err
	}

	resp.Password = trans.String("")

	return resp, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*v1.User, error) {
	resp, err := s.uc.Update(ctx, req)
	if err != nil {
		// s.log.Info(err)
		return nil, err
	}

	resp.Password = trans.String("")

	return resp, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *v1.DeleteUserRequest) (*emptypb.Empty, error) {
	_, err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
