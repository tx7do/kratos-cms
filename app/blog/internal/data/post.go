package data

import (
	"context"
	"kratos-blog/app/blog/internal/data/ent"
	"kratos-blog/app/blog/internal/data/ent/post"
	"kratos-blog/pkg/util/entgo"
	paging "kratos-blog/pkg/util/pagination"

	"github.com/go-kratos/kratos/v2/log"

	v1 "kratos-blog/api/blog/v1"
	"kratos-blog/app/blog/internal/biz"
	"kratos-blog/third_party/pagination"
)

var _ biz.PostRepo = (*PostRepo)(nil)

type PostRepo struct {
	data *Data
	log  *log.Helper
}

func NewPostRepo(data *Data, logger log.Logger) biz.PostRepo {
	l := log.NewHelper(log.With(logger, "module", "post/repo"))
	return &PostRepo{
		data: data,
		log:  l,
	}
}

func (r *PostRepo) convertEntToProto(in *ent.Post) *v1.Post {
	if in == nil {
		return nil
	}
	return &v1.Post{
		Id:      in.ID,
		Title:   in.Title,
		Summary: in.Summary,
		Content: in.Content,
		//Category: in.Category,
		//Tags:     in.Tags,
	}
}

func (r *PostRepo) List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListPostResponse, error) {
	whereCond, orderCond := entgo.QueryCommandToSelector(req.GetQuery(), req.GetOrderBy())

	builder1 := r.data.db.Post.Query()
	if len(whereCond) != 0 {
		for _, v := range whereCond {
			builder1.Where(v)
		}
	}
	if len(orderCond) != 0 {
		for _, v := range orderCond {
			builder1.Order(v)
		}
	} else {
		builder1.Order(ent.Desc(post.FieldCreateTime))
	}
	if req.Page != nil && req.PageSize != nil {
		builder1.
			Offset(paging.GetPageOffset(req.GetPage(), req.GetPageSize())).
			Limit(int(req.GetPageSize()))
	}
	posts, err := builder1.All(ctx)
	if err != nil {
		return nil, err
	}

	builder2 := r.data.db.Post.Query()
	if len(whereCond) != 0 {
		for _, v := range whereCond {
			builder2.Where(v)
		}
	}
	count, err := builder2.
		Select(post.FieldID).
		Count(ctx)
	if err != nil {
		return nil, err
	}

	items := make([]*v1.Post, 0, len(posts))
	for _, po := range posts {
		item := r.convertEntToProto(po)
		items = append(items, item)
	}

	ret := v1.ListPostResponse{
		Total: int32(count),
		Items: items,
	}

	return &ret, err
}

func (r *PostRepo) Get(ctx context.Context, req *v1.GetPostRequest) (*v1.Post, error) {
	po, err := r.data.db.Post.Get(ctx, req.GetId())
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	return r.convertEntToProto(po), err
}

func (r *PostRepo) Create(ctx context.Context, req *v1.CreatePostRequest) (*v1.Post, error) {
	po, err := r.data.db.Post.Create().
		SetNillableTitle(req.Post.Title).
		SetNillableSummary(req.Post.Summary).
		SetNillableContent(req.Post.Content).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.convertEntToProto(po), err
}

func (r *PostRepo) Update(ctx context.Context, req *v1.UpdatePostRequest) (*v1.Post, error) {
	builder := r.data.db.Post.UpdateOneID(req.Id).
		SetNillableTitle(req.Post.Title).
		SetNillableSummary(req.Post.Summary).
		SetNillableContent(req.Post.Content)

	po, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.convertEntToProto(po), err
}

func (r *PostRepo) Delete(ctx context.Context, req *v1.DeletePostRequest) (bool, error) {
	err := r.data.db.Post.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	return err != nil, err
}
