package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	"kratos-cms/app/core/service/internal/biz"
	"kratos-cms/app/core/service/internal/data/ent"
	"kratos-cms/app/core/service/internal/data/ent/comment"

	v1 "kratos-cms/gen/api/go/comment/service/v1"
	"kratos-cms/gen/api/go/common/pagination"

	"github.com/tx7do/kratos-utils/entgo"
	paging "github.com/tx7do/kratos-utils/pagination"
	util "github.com/tx7do/kratos-utils/time"
)

var _ biz.CommentRepo = (*CommentRepo)(nil)

type CommentRepo struct {
	data *Data
	log  *log.Helper
}

func NewCommentRepo(data *Data, logger log.Logger) biz.CommentRepo {
	l := log.NewHelper(log.With(logger, "module", "comment/repo/core-service"))
	return &CommentRepo{
		data: data,
		log:  l,
	}
}

func (r *CommentRepo) convertEntToProto(in *ent.Comment) *v1.Comment {
	if in == nil {
		return nil
	}
	return &v1.Comment{
		Id:                in.ID,
		Author:            in.Author,
		Email:             in.Email,
		IpAddress:         in.IPAddress,
		AuthorUrl:         in.AuthorURL,
		GravatarMd5:       in.GravatarMd5,
		Content:           in.Content,
		Status:            in.Status,
		UserAgent:         in.UserAgent,
		ParentId:          in.ParentID,
		IsAdmin:           in.IsAdmin,
		AllowNotification: in.AllowNotification,
		Avatar:            in.Avatar,
		CreateTime:        util.UnixMilliToStringPtr(in.CreateTime),
		UpdateTime:        util.UnixMilliToStringPtr(in.UpdateTime),
	}
}

func (r *CommentRepo) Count(ctx context.Context, whereCond entgo.WhereConditions) (int, error) {
	builder := r.data.db.Comment.Query()
	if len(whereCond) != 0 {
		for _, cond := range whereCond {
			builder = builder.Where(cond)
		}
	}
	return builder.Count(ctx)
}

func (r *CommentRepo) List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListCommentResponse, error) {
	whereCond, orderCond := entgo.QueryCommandToSelector(req.GetQuery(), req.GetOrderBy())

	builder := r.data.db.Comment.Query()
	if len(whereCond) != 0 {
		for _, v := range whereCond {
			builder.Where(v)
		}
	}
	if len(orderCond) != 0 {
		for _, v := range orderCond {
			builder.Order(v)
		}
	} else {
		builder.Order(ent.Desc(comment.FieldCreateTime))
	}
	if req.GetPage() > 0 && req.GetPageSize() > 0 && !req.GetNopaging() {
		builder.
			Offset(paging.GetPageOffset(req.GetPage(), req.GetPageSize())).
			Limit(int(req.GetPageSize()))
	}
	comments, err := builder.All(ctx)
	if err != nil {
		return nil, err
	}

	items := make([]*v1.Comment, 0, len(comments))
	for _, po := range comments {
		item := r.convertEntToProto(po)
		items = append(items, item)
	}

	count, err := r.Count(ctx, whereCond)
	if err != nil {
		return nil, err
	}

	ret := v1.ListCommentResponse{
		Total: int32(count),
		Items: items,
	}

	return &ret, err
}

func (r *CommentRepo) Get(ctx context.Context, req *v1.GetCommentRequest) (*v1.Comment, error) {
	po, err := r.data.db.Comment.Get(ctx, req.GetId())
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	return r.convertEntToProto(po), err
}

func (r *CommentRepo) Create(ctx context.Context, req *v1.CreateCommentRequest) (*v1.Comment, error) {
	po, err := r.data.db.Comment.Create().
		SetNillableAuthor(req.Comment.Author).
		SetNillableEmail(req.Comment.Email).
		SetNillableIPAddress(req.Comment.IpAddress).
		SetNillableAuthorURL(req.Comment.AuthorUrl).
		SetNillableGravatarMd5(req.Comment.GravatarMd5).
		SetNillableContent(req.Comment.Content).
		SetNillableStatus(req.Comment.Status).
		SetNillableUserAgent(req.Comment.UserAgent).
		SetNillableParentID(req.Comment.ParentId).
		SetNillableIsAdmin(req.Comment.IsAdmin).
		SetNillableAllowNotification(req.Comment.AllowNotification).
		SetNillableAvatar(req.Comment.Avatar).
		SetCreateTime(time.Now().UnixMilli()).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.convertEntToProto(po), err
}

func (r *CommentRepo) Update(ctx context.Context, req *v1.UpdateCommentRequest) (*v1.Comment, error) {
	builder := r.data.db.Comment.UpdateOneID(req.Id).
		SetNillableAuthor(req.Comment.Author).
		SetNillableEmail(req.Comment.Email).
		SetNillableIPAddress(req.Comment.IpAddress).
		SetNillableAuthorURL(req.Comment.AuthorUrl).
		SetNillableGravatarMd5(req.Comment.GravatarMd5).
		SetNillableContent(req.Comment.Content).
		SetNillableStatus(req.Comment.Status).
		SetNillableUserAgent(req.Comment.UserAgent).
		SetNillableParentID(req.Comment.ParentId).
		SetNillableIsAdmin(req.Comment.IsAdmin).
		SetNillableAllowNotification(req.Comment.AllowNotification).
		SetNillableAvatar(req.Comment.Avatar).
		SetUpdateTime(time.Now().UnixMilli())

	po, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.convertEntToProto(po), err
}

func (r *CommentRepo) Delete(ctx context.Context, req *v1.DeleteCommentRequest) (bool, error) {
	err := r.data.db.Comment.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	return err != nil, err
}
