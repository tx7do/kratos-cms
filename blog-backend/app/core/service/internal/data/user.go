package data

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-utils/crypto"
	entgo "github.com/tx7do/go-utils/entgo/query"
	util "github.com/tx7do/go-utils/time"
	"github.com/tx7do/go-utils/trans"

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

	var authority *v1.UserAuthority
	if in.Authority != nil {
		authority = (*v1.UserAuthority)(trans.Int32(v1.UserAuthority_value[string(*in.Authority)]))
	}

	return &v1.User{
		Id:          in.ID,
		UserName:    in.Username,
		NickName:    in.Nickname,
		Email:       in.Email,
		Avatar:      in.Avatar,
		Description: in.Description,
		Password:    in.Password,
		Authority:   authority,
		CreateTime:  util.UnixMilliToStringPtr(in.CreateTime),
		UpdateTime:  util.UnixMilliToStringPtr(in.UpdateTime),
	}
}

func (r *UserRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().User.Query()
	if len(whereCond) != 0 {
		builder.Modify(whereCond...)
	}

	count, err := builder.Count(ctx)
	if err != nil {
		r.log.Errorf("query count failed: %s", err.Error())
	}

	return count, err
}

func (r *UserRepo) List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListUserResponse, error) {
	builder := r.data.db.Client().User.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(r.data.db.Driver().Dialect(),
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), user.FieldCreateTime)
	if err != nil {
		r.log.Errorf("解析条件发生错误[%s]", err.Error())
		return nil, err
	}

	if querySelectors != nil {
		builder.Modify(querySelectors...)
	}

	results, err := builder.All(ctx)
	if err != nil {
		r.log.Errorf("query list failed: %s", err.Error())
		return nil, err
	}

	items := make([]*v1.User, 0, len(results))
	for _, res := range results {
		item := r.convertEntToProto(res)
		if item != nil {
			item.Password = nil
		}
		items = append(items, item)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &v1.ListUserResponse{
		Total: int32(count),
		Items: items,
	}, nil
}

func (r *UserRepo) Get(ctx context.Context, req *v1.GetUserRequest) (*v1.User, error) {
	ret, err := r.data.db.Client().User.Get(ctx, req.GetId())
	if err != nil && !ent.IsNotFound(err) {
		r.log.Errorf("query one data failed: %s", err.Error())
		return nil, err
	}

	u := r.convertEntToProto(ret)
	if u != nil {
		u.Password = nil
	}

	return u, err
}

func (r *UserRepo) Create(ctx context.Context, req *v1.CreateUserRequest) (*v1.User, error) {
	ph, err := crypto.HashPassword(req.User.GetPassword())
	if err != nil {
		return nil, err
	}

	builder := r.data.db.Client().User.Create().
		SetNillableUsername(req.User.UserName).
		SetNillableNickname(req.User.NickName).
		SetNillableEmail(req.User.Email).
		SetPassword(ph).
		SetCreateTime(time.Now().UnixMilli())

	if req.User.Authority != nil {
		builder.SetAuthority((user.Authority)(req.User.Authority.String()))
	}

	ret, err := builder.Save(ctx)
	if err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return nil, err
	}

	u := r.convertEntToProto(ret)
	if u != nil {
		u.Password = nil
	}

	return u, err
}

func (r *UserRepo) Update(ctx context.Context, req *v1.UpdateUserRequest) (*v1.User, error) {
	cryptoPassword, err := crypto.HashPassword(req.User.GetPassword())
	if err != nil {
		return nil, err
	}

	builder := r.data.db.Client().User.UpdateOneID(req.Id).
		SetNillableNickname(req.User.NickName).
		SetNillableEmail(req.User.Email).
		SetPassword(cryptoPassword).
		SetUpdateTime(time.Now().UnixMilli())

	if req.User.Authority != nil {
		builder.SetAuthority((user.Authority)(req.User.Authority.String()))
	}

	ret, err := builder.Save(ctx)
	if err != nil {
		r.log.Errorf("update one data failed: %s", err.Error())
		return nil, err
	}

	u := r.convertEntToProto(ret)
	if u != nil {
		u.Password = nil
	}

	return u, err
}

func (r *UserRepo) Delete(ctx context.Context, req *v1.DeleteUserRequest) (bool, error) {
	err := r.data.db.Client().User.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	if err != nil {
		r.log.Errorf("delete one data failed: %s", err.Error())
	}

	return err == nil, err
}

func (r *UserRepo) GetUserByUserName(ctx context.Context, userName string) (*v1.User, error) {
	ret, err := r.data.db.Client().User.Query().
		Where(user.UsernameEQ(userName)).
		Only(ctx)
	if err != nil {
		r.log.Errorf("query user data failed: %s", err.Error())
		return nil, err
	}

	u := r.convertEntToProto(ret)
	if u != nil {
		u.Password = nil
	}

	return u, err
}

func (r *UserRepo) VerifyPassword(ctx context.Context, req *v1.VerifyPasswordRequest) (bool, error) {
	res, err := r.data.db.Client().User.
		Query().
		Select(user.FieldID, user.FieldPassword).
		Where(user.UsernameEQ(req.GetUserName())).
		Only(ctx)
	if err != nil {
		r.log.Errorf("query user data failed: %s", err.Error())
		return false, v1.ErrorUserNotFound("用户未找到")
	}

	bMatched := crypto.CheckPasswordHash(req.GetPassword(), *res.Password)

	if !bMatched {
		return false, v1.ErrorUserNotFound("密码错误")
	}

	return true, nil
}
