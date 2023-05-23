package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-utils/crypto"
	"github.com/tx7do/kratos-utils/entgo"
	paging "github.com/tx7do/kratos-utils/pagination"
	util "github.com/tx7do/kratos-utils/time"

	"kratos-cms/app/core/service/internal/biz"
	"kratos-cms/app/core/service/internal/data/ent"
	"kratos-cms/app/core/service/internal/data/ent/user"

	"kratos-cms/gen/api/go/common/pagination"
	v1 "kratos-cms/gen/api/go/user/service/v1"
)

var _ biz.UserRepo = (*UserRepo)(nil)

type UserRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	l := log.NewHelper(log.With(logger, "module", "user/repo/core-service"))
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
		CreateTime:  util.UnixMilliToStringPtr(in.CreateTime),
		UpdateTime:  util.UnixMilliToStringPtr(in.UpdateTime),
	}
}

func (r *UserRepo) Count(ctx context.Context, whereCond entgo.WhereConditions) (int, error) {
	builder := r.data.db.User.Query()
	if len(whereCond) != 0 {
		for _, cond := range whereCond {
			builder = builder.Where(cond)
		}
	}
	return builder.Count(ctx)
}

func (r *UserRepo) List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListUserResponse, error) {
	whereCond, orderCond := entgo.QueryCommandToSelector(req.GetQuery(), req.GetOrderBy())

	builder := r.data.db.User.Query()

	if len(whereCond) != 0 {
		for _, cond := range whereCond {
			builder = builder.Where(cond)
		}
	}
	if len(orderCond) != 0 {
		for _, cond := range orderCond {
			builder = builder.Order(cond)
		}
	} else {
		builder.Order(ent.Desc(user.FieldCreateTime))
	}
	if req.GetPage() > 0 && req.GetPageSize() > 0 && !req.GetNopaging() {
		builder.
			Offset(paging.GetPageOffset(req.GetPage(), req.GetPageSize())).
			Limit(int(req.GetPageSize()))
	}
	results, err := builder.All(ctx)
	if err != nil {
		return nil, err
	}

	items := make([]*v1.User, 0, len(results))
	for _, res := range results {
		item := r.convertEntToProto(res)
		items = append(items, item)
	}

	count, err := r.Count(ctx, whereCond)
	if err != nil {
		return nil, err
	}

	return &v1.ListUserResponse{
		Total: int32(count),
		Items: items,
	}, nil
}

func (r *UserRepo) Get(ctx context.Context, req *v1.GetUserRequest) (*v1.User, error) {
	res, err := r.data.db.User.Get(ctx, req.GetId())
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	return r.convertEntToProto(res), err
}

func (r *UserRepo) Create(ctx context.Context, req *v1.CreateUserRequest) (*v1.User, error) {
	ph, err := crypto.HashPassword(req.User.GetPassword())
	if err != nil {
		return nil, err
	}

	res, err := r.data.db.User.Create().
		SetNillableUsername(req.User.UserName).
		SetNillableNickname(req.User.NickName).
		SetNillableEmail(req.User.Email).
		SetPassword(ph).
		SetCreateTime(time.Now().UnixMilli()).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.convertEntToProto(res), err
}

func (r *UserRepo) Update(ctx context.Context, req *v1.UpdateUserRequest) (*v1.User, error) {
	cryptoPassword, err := crypto.HashPassword(req.User.GetPassword())
	if err != nil {
		return nil, err
	}

	builder := r.data.db.User.UpdateOneID(req.Id).
		SetNillableNickname(req.User.NickName).
		SetNillableEmail(req.User.Email).
		SetPassword(cryptoPassword).
		SetUpdateTime(time.Now().UnixMilli())

	res, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.convertEntToProto(res), err
}

func (r *UserRepo) Delete(ctx context.Context, req *v1.DeleteUserRequest) (bool, error) {
	err := r.data.db.User.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	return err != nil, err
}

func (r *UserRepo) GetUserByUserName(ctx context.Context, userName string) (*v1.User, error) {
	res, err := r.data.db.User.Query().
		Where(user.UsernameEQ(userName)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return r.convertEntToProto(res), nil
}

func (r *UserRepo) VerifyPassword(ctx context.Context, req *v1.VerifyPasswordRequest) (bool, error) {
	res, err := r.data.db.User.
		Query().
		Select(user.FieldID, user.FieldPassword).
		Where(user.UsernameEQ(req.GetUserName())).
		Only(ctx)
	if err != nil {
		return false, v1.ErrorUserNotFound("用户未找到")
	}

	bMatched := crypto.CheckPasswordHash(req.GetPassword(), *res.Password)

	if !bMatched {
		return false, v1.ErrorUserNotFound("密码错误")
	}

	return true, nil
}
