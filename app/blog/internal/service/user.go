package service

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtV4 "github.com/golang-jwt/jwt/v4"
	"google.golang.org/protobuf/types/known/emptypb"
	v1 "kratos-blog/api/blog/v1"
	"kratos-blog/app/blog/internal/biz"
	"strconv"
)

type UserService struct {
	v1.UnimplementedUserServiceServer

	uc  *biz.UserUseCase
	tuc *biz.UserTokenUseCase
	log *log.Helper
}

func NewUserService(logger log.Logger, uc *biz.UserUseCase, tuc *biz.UserTokenUseCase) *UserService {
	l := log.NewHelper(log.With(logger, "module", "service/user"))
	return &UserService{
		log: l,
		uc:  uc,
		tuc: tuc,
	}
}

// Login 登陆
func (s *UserService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	_, err := s.uc.VerifyPassword(ctx, req)
	if err != nil {
		return &v1.LoginResponse{}, err
	}

	user, err := s.uc.GetUserByUserName(ctx, req.GetUserName())
	if err != nil {
		return &v1.LoginResponse{}, err
	}

	token, err := s.tuc.GenerateToken(ctx, user)
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
func (s *UserService) Logout(ctx context.Context, req *v1.LogoutRequest) (*emptypb.Empty, error) {
	err := s.tuc.RemoveToken(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// Register 注册
func (s *UserService) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.RegisterResponse, error) {
	user, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	token, err := s.tuc.GenerateToken(ctx, user)
	if err != nil {
		return &v1.RegisterResponse{}, err
	}

	return &v1.RegisterResponse{
		Id:       user.GetId(),
		UserName: user.GetUserName(),
		Token:    token,
	}, nil
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

func (s *UserService) GetMe(ctx context.Context, _ *emptypb.Empty) (*v1.User, error) {
	userId, _, err := s.ParseFromContext(ctx)
	if err != nil {
		return nil, v1.ErrorRequestNotSupport("权限信息不存在")
	}

	return s.uc.Get(ctx, userId)
}
