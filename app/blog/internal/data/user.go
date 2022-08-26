package data

import (
	"context"
	"kratos-blog/app/blog/internal/data/ent"
	"kratos-blog/app/blog/internal/data/ent/user"
	"kratos-blog/pkg/util/crypto"
	"kratos-blog/pkg/util/entgo"

	"github.com/go-kratos/kratos/v2/log"

	v1 "kratos-blog/api/blog/v1"
	"kratos-blog/app/blog/internal/biz"
)

var _ biz.UserRepo = (*UserRepo)(nil)

type UserRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	l := log.NewHelper(log.With(logger, "module", "user/repo"))
	return &UserRepo{
		data: data,
		log:  l,
	}
}

func (r *UserRepo) convertEntToProto(in *ent.User) *v1.User {
	if in == nil {
		return nil
	}
	return &v1.User{
		Id:          in.ID,
		UserName:    in.Username,
		NickName:    in.Nickname,
		Email:       in.Email,
		Avatar:      in.Avatar,
		Description: in.Description,
		Password:    in.Password,
		CreateTime:  entgo.UnixMilliToStringPtr(in.CreateTime),
		UpdateTime:  entgo.UnixMilliToStringPtr(in.UpdateTime),
	}
}

func (r *UserRepo) Create(ctx context.Context, req *v1.RegisterRequest) (*v1.User, error) {
	ph, err := crypto.HashPassword(req.User.GetPassword())
	if err != nil {
		return nil, err
	}

	po, err := r.data.db.User.Create().
		SetNillableUsername(req.User.UserName).
		SetNillableNickname(req.User.NickName).
		SetNillableEmail(req.User.Email).
		SetPassword(ph).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.convertEntToProto(po), err
}

func (r *UserRepo) Update(ctx context.Context, req *v1.UpdateUserRequest) (*v1.User, error) {
	ph, err := crypto.HashPassword(req.User.GetPassword())
	if err != nil {
		return nil, err
	}

	builder := r.data.db.User.UpdateOneID(req.Id).
		SetNillableNickname(req.User.NickName).
		SetNillableEmail(req.User.Email).
		SetPassword(ph)

	po, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.convertEntToProto(po), err
}

func (r *UserRepo) Delete(ctx context.Context, req *v1.DeleteUserRequest) (bool, error) {
	err := r.data.db.User.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	return err != nil, err
}

func (r *UserRepo) VerifyPassword(ctx context.Context, req *v1.LoginRequest) (bool, error) {
	ret, err := r.data.db.User.
		Query().
		Select(user.FieldID, user.FieldPassword).
		Where(user.UsernameEQ(req.GetUserName())).
		Only(ctx)
	if err != nil {
		return false, v1.ErrorUserNotFound("用户未找到")
	}

	bMatched := crypto.CheckPasswordHash(req.GetPassword(), *ret.Password)

	if !bMatched {
		return false, v1.ErrorUserNotFound("密码错误")
	}

	return true, nil
}

func (r *UserRepo) Get(ctx context.Context, req uint32) (*v1.User, error) {
	po, err := r.data.db.User.Get(ctx, req)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	return r.convertEntToProto(po), err
}

func (r *UserRepo) GetUserByUserName(ctx context.Context, userName string) (*v1.User, error) {
	u, err := r.data.db.User.Query().
		Where(user.UsernameEQ(userName)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return r.convertEntToProto(u), nil
}
