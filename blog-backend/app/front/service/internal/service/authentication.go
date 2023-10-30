package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-cms/app/front/service/internal/biz"

	v1 "kratos-cms/gen/api/go/front/service/v1"
	userV1 "kratos-cms/gen/api/go/user/service/v1"

	"kratos-cms/pkg/middleware/auth"
)

type AuthenticationService struct {
	v1.AuthenticationServiceHTTPServer

	sc   userV1.UserServiceClient
	utuc *biz.UserTokenUseCase

	log *log.Helper
}

func NewAuthenticationService(logger log.Logger, sc userV1.UserServiceClient, utuc *biz.UserTokenUseCase) *AuthenticationService {
	l := log.NewHelper(log.With(logger, "module", "authn/service/front-service"))
	return &AuthenticationService{
		log:  l,
		sc:   sc,
		utuc: utuc,
	}
}

// Login 登陆
func (s *AuthenticationService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	if _, err := s.sc.VerifyPassword(ctx, &userV1.VerifyPasswordRequest{
		UserName: req.GetUsername(),
		Password: req.GetPassword(),
	}); err != nil {
		return &v1.LoginResponse{}, err
	}

	user, err := s.sc.GetUserByUserName(ctx, &userV1.GetUserByUserNameRequest{UserName: req.GetUsername()})
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

	return s.sc.GetUser(ctx, &userV1.GetUserRequest{
		Id: req.GetId(),
	})
}
