package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	frontV1 "kratos-cms/api/gen/go/front/service/v1"
	userV1 "kratos-cms/api/gen/go/user/service/v1"

	"kratos-cms/pkg/cache"
	"kratos-cms/pkg/middleware/auth"
)

type AuthenticationService struct {
	frontV1.AuthenticationServiceHTTPServer

	sc   userV1.UserServiceClient
	utuc *cache.UserToken

	log *log.Helper
}

func NewAuthenticationService(logger log.Logger, sc userV1.UserServiceClient, utuc *cache.UserToken) *AuthenticationService {
	l := log.NewHelper(log.With(logger, "module", "authn/service/front-service"))
	return &AuthenticationService{
		log:  l,
		sc:   sc,
		utuc: utuc,
	}
}

// Login 登陆
func (s *AuthenticationService) Login(ctx context.Context, req *frontV1.LoginRequest) (*frontV1.LoginResponse, error) {
	if _, err := s.sc.VerifyPassword(ctx, &userV1.VerifyPasswordRequest{
		UserName: req.GetUsername(),
		Password: req.GetPassword(),
	}); err != nil {
		return &frontV1.LoginResponse{}, err
	}

	user, err := s.sc.GetUserByUserName(ctx, &userV1.GetUserByUserNameRequest{UserName: req.GetUsername()})
	if err != nil {
		return &frontV1.LoginResponse{}, err
	}

	accessToken, refreshToken, err := s.utuc.GenerateToken(ctx, user)
	if err != nil {
		return nil, err
	}

	return &frontV1.LoginResponse{
		TokenType:    frontV1.TokenType_bearer.String(),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// Logout 登出
func (s *AuthenticationService) Logout(ctx context.Context, req *frontV1.LogoutRequest) (*emptypb.Empty, error) {
	err := s.utuc.RemoveToken(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *AuthenticationService) GetMe(ctx context.Context, req *frontV1.GetMeRequest) (*userV1.User, error) {
	authInfo, err := auth.FromContext(ctx)
	if err != nil {
		s.log.Errorf("用户认证失败[%s]", err.Error())
		return nil, frontV1.ErrorAccessForbidden("用户认证失败")
	}

	req.Id = authInfo.UserId

	return s.sc.GetUser(ctx, &userV1.GetUserRequest{
		Id: req.GetId(),
	})
}
