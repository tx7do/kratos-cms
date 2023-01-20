package data

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	jwtV4 "github.com/golang-jwt/jwt/v4"

	"kratos-blog/app/user/service/internal/biz"
	"kratos-blog/app/user/service/internal/conf"

	v1 "kratos-blog/gen/api/go/user/service/v1"
)

// 官方的7个字段
const (
	ClaimID        = "jti" // 编号
	ClaimSubject   = "sub" // 主题
	ClaimIssuer    = "iss" // 签发人
	ClaimAudience  = "aud" // 受众
	ClaimExpiresAt = "exp" // 过期时间
	ClaimNotBefore = "nbf" // 生效时间
	ClaimIssuedAt  = "iat" // 签发时间
)

const (
	ClaimScopes  = "scopes"
	ClaimUserId  = "userId"
	ClaimEnabled = "enabled"
)

var _ biz.UserTokenRepo = (*userTokenRepo)(nil)

type userTokenRepo struct {
	data *Data
	log  *log.Helper
	key  []byte
}

func NewUserTokenRepo(data *Data, conf *conf.Auth, logger log.Logger) biz.UserTokenRepo {
	l := log.NewHelper(log.With(logger, "module", "user-token/repo/user-service"))
	return &userTokenRepo{
		data: data,
		log:  l,
		key:  []byte(conf.ApiKey),
	}
}

func (r *userTokenRepo) createAccessJwtToken(username string, userId uint32) string {
	nowTime := time.Now()

	claims := jwtV4.NewWithClaims(jwtV4.SigningMethodHS256,
		jwtV4.MapClaims{
			ClaimSubject:  username,
			ClaimUserId:   strconv.FormatUint(uint64(userId), 10),
			ClaimIssuedAt: nowTime.Unix(),
		})

	signedToken, err := claims.SignedString(r.key)
	if err != nil {
		return ""
	}

	return signedToken
}

func (r *userTokenRepo) parseAccessJwtTokenFromString(token string) (string, uint32, error) {
	parseAuth, err := jwtV4.Parse(token, func(*jwtV4.Token) (interface{}, error) {
		return r.key, nil
	})
	if err != nil {
		return "", 0, err
	}

	claims, ok := parseAuth.Claims.(jwtV4.MapClaims)
	if !ok {
		return "", 0, errors.New("no jwt token in context")
	}

	userName, userId, err := r.parseAccessJwtToken(claims)
	if err != nil {
		return "", 0, err
	}

	return userName, userId, nil
}

func (r *userTokenRepo) parseAccessJwtToken(claims jwtV4.Claims) (string, uint32, error) {
	if claims == nil {
		return "", 0, errors.New("claims is nil")
	}

	mc, ok := claims.(jwtV4.MapClaims)
	if !ok {
		return "", 0, errors.New("claims is not map claims")
	}

	var strUserName string
	userName, ok := mc[ClaimSubject]
	if ok {
		strUserName = userName.(string)
	}

	var userId uint32
	strUserId, ok := mc[ClaimUserId]
	if ok {
		userId_, err := strconv.ParseUint(strUserId.(string), 10, 64)
		if err != nil {
			return "", 0, err
		}
		userId = uint32(userId_)
	}

	return strUserName, userId, nil
}

func (r *userTokenRepo) GenerateToken(ctx context.Context, user *v1.User) (string, error) {
	token := r.createAccessJwtToken(user.GetUserName(), user.GetId())
	if token == "" {
		return "", errors.New("create token failed")
	}

	err := r.setToken(ctx, user.GetId(), token)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *userTokenRepo) GetToken(ctx context.Context, userId uint32) string {
	return r.getToken(ctx, userId)
}

func (r *userTokenRepo) ValidateToken(ctx context.Context, token string) error {
	_, userId, err := r.parseAccessJwtTokenFromString(token)
	if err != nil {
		return v1.ErrorInvalidToken("非法的令牌")
	}

	validToken := r.getToken(ctx, userId)
	if validToken == "" {
		return v1.ErrorTokenNotExist("令牌不存在")
	}

	return nil
}

func (r *userTokenRepo) RemoveToken(ctx context.Context, userId uint32) error {
	validToken := r.getToken(ctx, userId)
	if validToken == "" {
		return v1.ErrorTokenNotExist("令牌不存在")
	}

	return r.deleteToken(ctx, userId)
}

func (r *userTokenRepo) RemoveUserToken(ctx context.Context, userId uint32) error {
	validToken := r.getToken(ctx, userId)
	if validToken == "" {
		return v1.ErrorTokenNotExist("令牌不存在")
	}

	return r.deleteToken(ctx, userId)
}

const userTokenKeyPrefix = "ut_"

func (r *userTokenRepo) setToken(ctx context.Context, userId uint32, token string) error {
	key := fmt.Sprintf("%s%d", userTokenKeyPrefix, userId)
	return r.data.rdb.Set(ctx, key, token, 0).Err()
}

func (r *userTokenRepo) getToken(ctx context.Context, userId uint32) string {
	key := fmt.Sprintf("%s%d", userTokenKeyPrefix, userId)
	result, err := r.data.rdb.Get(ctx, key).Result()
	if err != nil {
		if err != redis.Nil {
			r.log.Error("get redis user token failed: %s", err.Error())
		}
		return ""
	}
	return result
}

func (r *userTokenRepo) deleteToken(ctx context.Context, userId uint32) error {
	key := fmt.Sprintf("%s%d", userTokenKeyPrefix, userId)
	return r.data.rdb.Del(ctx, key).Err()
}
