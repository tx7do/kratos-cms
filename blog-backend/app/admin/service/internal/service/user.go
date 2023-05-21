package service

import (
	"context"
	"kratos-cms/pkg/util/authn"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
	"kratos-cms/gen/api/go/common/pagination"

	v1 "kratos-cms/gen/api/go/admin/service/v1"
	userV1 "kratos-cms/gen/api/go/user/service/v1"
)

type UserService struct {
	v1.UserServiceHTTPServer

	uc  userV1.UserServiceClient
	log *log.Helper
}

func NewUserService(logger log.Logger, uc userV1.UserServiceClient) *UserService {
	l := log.NewHelper(log.With(logger, "module", "user/service/admin-service"))
	return &UserService{
		log: l,
		uc:  uc,
	}
}

func (s *UserService) ListUser(ctx context.Context, req *pagination.PagingRequest) (*userV1.ListUserResponse, error) {
	return s.uc.ListUser(ctx, req)
}

func (s *UserService) GetUser(ctx context.Context, req *userV1.GetUserRequest) (*userV1.User, error) {
	return s.uc.GetUser(ctx, req)
}

func (s *UserService) GetUserByUserName(ctx context.Context, req *userV1.GetUserByUserNameRequest) (*userV1.User, error) {
	return s.uc.GetUserByUserName(ctx, req)
}

func (s *UserService) CreateUser(ctx context.Context, req *userV1.CreateUserRequest) (*userV1.User, error) {
	userId, _, err := authn.ParseFromContext(ctx)
	if err != nil {
		return nil, v1.ErrorRequestNotSupport("%d 权限信息不存在", userId)
	}

	if req.User == nil {
		return nil, v1.ErrorBadRequest("错误的参数")
	}

	authority := "CUSTOMER_USER"

	req.OperatorId = userId
	req.User.CreatorId = &userId
	if req.User.Authority == nil {
		req.User.Authority = &authority
	}

	return s.uc.CreateUser(ctx, req)
}

func (s *UserService) UpdateUser(ctx context.Context, req *userV1.UpdateUserRequest) (*userV1.User, error) {
	userId, _, err := authn.ParseFromContext(ctx)
	if err != nil {
		return nil, v1.ErrorRequestNotSupport("%d 权限信息不存在", userId)
	}

	if req.User == nil {
		return nil, v1.ErrorBadRequest("错误的参数")
	}

	req.OperatorId = userId

	return s.uc.UpdateUser(ctx, req)
}

func (s *UserService) DeleteUser(ctx context.Context, req *userV1.DeleteUserRequest) (*emptypb.Empty, error) {
	userId, _, err := authn.ParseFromContext(ctx)
	if err != nil {
		return nil, v1.ErrorRequestNotSupport("%d 权限信息不存在", userId)
	}

	req.OperatorId = userId

	return s.uc.DeleteUser(ctx, req)
}

func (s *UserService) UserExists(ctx context.Context, req *userV1.UserExistsRequest) (*userV1.UserExistsResponse, error) {
	return s.uc.UserExists(ctx, req)
}
