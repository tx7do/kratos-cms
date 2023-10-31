package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-utils/trans"
	"google.golang.org/protobuf/types/known/emptypb"

	v1 "kratos-cms/gen/api/go/admin/service/v1"
	"kratos-cms/gen/api/go/common/pagination"
	userV1 "kratos-cms/gen/api/go/user/service/v1"

	"kratos-cms/pkg/middleware/auth"
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
	authInfo, err := auth.FromContext(ctx)
	if err != nil {
		s.log.Errorf("用户认证失败[%s]", err.Error())
		return nil, v1.ErrorAccessForbidden("用户认证失败")
	}

	if req.User == nil {
		return nil, v1.ErrorBadRequest("错误的参数")
	}

	req.OperatorId = authInfo.UserId
	req.User.CreatorId = trans.Uint32(authInfo.UserId)
	if req.User.Authority == nil {
		req.User.Authority = userV1.UserAuthority_CUSTOMER_USER.Enum()
	}

	return s.uc.CreateUser(ctx, req)
}

func (s *UserService) UpdateUser(ctx context.Context, req *userV1.UpdateUserRequest) (*userV1.User, error) {
	authInfo, err := auth.FromContext(ctx)
	if err != nil {
		s.log.Errorf("用户认证失败[%s]", err.Error())
		return nil, v1.ErrorAccessForbidden("用户认证失败")
	}

	if req.User == nil {
		return nil, v1.ErrorBadRequest("错误的参数")
	}

	req.OperatorId = authInfo.UserId

	return s.uc.UpdateUser(ctx, req)
}

func (s *UserService) DeleteUser(ctx context.Context, req *userV1.DeleteUserRequest) (*emptypb.Empty, error) {
	authInfo, err := auth.FromContext(ctx)
	if err != nil {
		s.log.Errorf("用户认证失败[%s]", err.Error())
		return nil, v1.ErrorAccessForbidden("用户认证失败")
	}

	req.OperatorId = authInfo.UserId

	return s.uc.DeleteUser(ctx, req)
}

func (s *UserService) UserExists(ctx context.Context, req *userV1.UserExistsRequest) (*userV1.UserExistsResponse, error) {
	return s.uc.UserExists(ctx, req)
}
