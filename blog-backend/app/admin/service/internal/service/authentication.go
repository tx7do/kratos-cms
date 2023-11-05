package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-cms/app/admin/service/internal/data"

	v1 "kratos-cms/gen/api/go/admin/service/v1"
	userV1 "kratos-cms/gen/api/go/user/service/v1"

	"kratos-cms/pkg/middleware/auth"
)

type AuthenticationService struct {
	v1.AuthenticationServiceHTTPServer

	uc   userV1.UserServiceClient
	utuc *data.UserTokenRepo

	log *log.Helper
}

func NewAuthenticationService(logger log.Logger, uc userV1.UserServiceClient, utuc *data.UserTokenRepo) *AuthenticationService {
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
		UserName: req.GetUsername(),
		Password: req.GetPassword(),
	}); err != nil {
		return &v1.LoginResponse{}, err
	}

	user, err := s.uc.GetUserByUserName(ctx, &userV1.GetUserByUserNameRequest{UserName: req.GetUsername()})
	if err != nil {
		return &v1.LoginResponse{}, err
	}

	token, err := s.utuc.GenerateToken(ctx, user)
	if err != nil {
		return &v1.LoginResponse{}, err
	}

	return &v1.LoginResponse{
		TokenType:   "bearer",
		AccessToken: token,
		Id:          user.GetId(),
		Username:    user.GetUserName(),
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
	authInfo, err := auth.FromContext(ctx)
	if err != nil {
		s.log.Errorf("用户认证失败[%s]", err.Error())
		return nil, v1.ErrorAccessForbidden("用户认证失败")
	}

	req.Id = authInfo.UserId

	return s.uc.GetUser(ctx, &userV1.GetUserRequest{
		Id: req.GetId(),
	})
}
