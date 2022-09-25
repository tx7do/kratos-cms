package service

import (
	"context"
	"errors"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtV4 "github.com/golang-jwt/jwt/v4"
	"google.golang.org/protobuf/types/known/emptypb"

	v1 "kratos-blog/api/user/service/v1"
)

type UserService struct {
	v1.UnimplementedUserServiceServer

	userClient v1.UserServiceClient
	log        *log.Helper
}

func NewUserService(logger log.Logger, userClient v1.UserServiceClient) *UserService {
	l := log.NewHelper(log.With(logger, "module", "service/user"))
	return &UserService{
		log:        l,
		userClient: userClient,
	}
}

// Login 登陆
func (s *UserService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	return s.userClient.Login(ctx, req)
}

// Logout 登出
func (s *UserService) Logout(ctx context.Context, req *v1.LogoutRequest) (*emptypb.Empty, error) {
	return s.userClient.Logout(ctx, req)
}

// Register 注册
func (s *UserService) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.RegisterResponse, error) {
	return s.userClient.Register(ctx, req)
}

func (s *UserService) ParseFromContext(ctx context.Context) (uint32, string, error) {
	claims, ok := jwt.FromContext(ctx)
	if !ok {
		return 0, "", errors.New("no jwt token in context")
	}

	mc, ok := claims.(jwtV4.MapClaims)
	if !ok {
		return 0, "", errors.New("claims is not map claims")
	}

	var strUserName string
	userName, ok := mc["sub"]
	if ok {
		strUserName = userName.(string)
	}

	var userId uint32
	strUserId, ok := mc["userId"]
	if ok {
		userId_, err := strconv.ParseUint(strUserId.(string), 10, 64)
		if err != nil {
			return 0, "", err
		}
		userId = uint32(userId_)
	}

	return userId, strUserName, nil
}

func (s *UserService) GetMe(ctx context.Context, req *v1.GetMeRequest) (*v1.User, error) {
	userId, _, err := s.ParseFromContext(ctx)
	if err != nil {
		return nil, v1.ErrorRequestNotSupport("%d 权限信息不存在", userId)
	}

	req.Id = &userId

	return s.userClient.GetMe(ctx, req)
}
