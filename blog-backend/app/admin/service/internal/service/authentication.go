package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-blog/app/admin/service/internal/biz"

	v1 "kratos-blog/gen/api/go/admin/service/v1"
	userV1 "kratos-blog/gen/api/go/user/service/v1"
	"kratos-blog/pkg/util/auth"
)

type AuthenticationService struct {
	v1.AuthenticationServiceHTTPServer

	uc   userV1.UserServiceClient
	utuc *biz.UserTokenUseCase

	log *log.Helper
}

func NewAuthenticationService(logger log.Logger, uc userV1.UserServiceClient, utuc *biz.UserTokenUseCase) *AuthenticationService {
	l := log.NewHelper(log.With(logger, "module", "user/service/admin-service"))
	return &AuthenticationService{
		log:  l,
		uc:   uc,
		utuc: utuc,
	}
}

// Login 登陆
func (s *AuthenticationService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	if _, err := s.uc.VerifyPassword(ctx, &userV1.VerifyPasswordRequest{
		UserName: req.GetUserName(),
		Password: req.GetPassword(),
	}); err != nil {
		return &v1.LoginResponse{}, err
	}

	user, err := s.uc.GetUserByUserName(ctx, &userV1.GetUserByUserNameRequest{UserName: req.GetUserName()})
	if err != nil {
		return &v1.LoginResponse{}, err
	}

	token, err := s.utuc.GenerateToken(ctx, user)
	if err != nil {
		return &v1.LoginResponse{}, err
	}

	return &v1.LoginResponse{
		Token:    token,
		Id:       user.GetId(),
		UserName: user.GetUserName(),
	}, nil
}

// Logout 登出
func (s *AuthenticationService) Logout(ctx context.Context, req *v1.LogoutRequest) (*emptypb.Empty, error) {
	err := s.utuc.RemoveToken(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *AuthenticationService) GetMe(ctx context.Context, req *v1.GetMeRequest) (*userV1.User, error) {
	userId, _, err := auth.ParseFromContext(ctx)
	if err != nil {
		return nil, v1.ErrorRequestNotSupport("%d 权限信息不存在", userId)
	}

	req.Id = userId

	return s.uc.GetUser(ctx, &userV1.GetUserRequest{
		Id: req.GetId(),
	})
}
