package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-utils/entgo"
	paging "github.com/tx7do/kratos-utils/pagination"
	util "github.com/tx7do/kratos-utils/time"

	"kratos-cms/app/core/service/internal/biz"
	"kratos-cms/app/core/service/internal/data/ent"
	"kratos-cms/app/core/service/internal/data/ent/tag"

	"kratos-cms/gen/api/go/common/pagination"
	"kratos-cms/gen/api/go/content/service/v1"
)

var _ biz.TagRepo = (*TagRepo)(nil)

type TagRepo struct {
	data *Data
	log  *log.Helper
}

func NewTagRepo(data *Data, logger log.Logger) biz.TagRepo {
	l := log.NewHelper(log.With(logger, "module", "tag/repo/core-service"))
	return &TagRepo{
		data: data,
		log:  l,
	}
}

func (r *TagRepo) convertEntToProto(in *ent.Tag) *v1.Tag {
	if in == nil {
		return nil
	}
	return &v1.Tag{
		Id:         in.ID,
		Name:       in.Name,
		Slug:       in.Slug,
		Color:      in.Color,
		Thumbnail:  in.Thumbnail,
		SlugName:   in.SlugName,
		PostCount:  in.PostCount,
		CreateTime: util.UnixMilliToStringPtr(in.CreateTime),
		UpdateTime: util.UnixMilliToStringPtr(in.UpdateTime),
	}
}

func (r *TagRepo) Count(ctx context.Context, whereCond entgo.WhereConditions) (int, error) {
	builder := r.data.db.Tag.Query()
	if len(whereCond) != 0 {
		for _, cond := range whereCond {
			builder = builder.Where(cond)
		}
	}
	return builder.Count(ctx)
}

func (r *TagRepo) List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListTagResponse, error) {
	whereCond, orderCond := entgo.QueryCommandToSelector(req.GetQuery(), req.GetOrderBy())

	builder := r.data.db.Tag.Query()
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
		builder.Order(ent.Desc(tag.FieldCreateTime))
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

	items := make([]*v1.Tag, 0, len(results))
	for _, res := range results {
		item := r.convertEntToProto(res)
		items = append(items, item)
	}

	count, err := r.Count(ctx, whereCond)
	if err != nil {
		return nil, err
	}

	return &v1.ListTagResponse{
		Total: int32(count),
		Items: items,
	}, nil
}

func (r *TagRepo) Get(ctx context.Context, req *v1.GetTagRequest) (*v1.Tag, error) {
	res, err := r.data.db.Tag.Get(ctx, req.GetId())
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	return r.convertEntToProto(res), err
}

func (r *TagRepo) Create(ctx context.Context, req *v1.CreateTagRequest) (*v1.Tag, error) {
	res, err := r.data.db.Tag.Create().
		SetNillableName(req.Tag.Name).
		SetNillableSlug(req.Tag.Slug).
		SetNillableColor(req.Tag.Color).
		SetNillableThumbnail(req.Tag.Thumbnail).
		SetNillableSlugName(req.Tag.SlugName).
		SetNillablePostCount(req.Tag.PostCount).
		SetCreateTime(time.Now().UnixMilli()).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.convertEntToProto(res), err
}

func (r *TagRepo) Update(ctx context.Context, req *v1.UpdateTagRequest) (*v1.Tag, error) {
	builder := r.data.db.Tag.UpdateOneID(req.Id).
		SetNillableName(req.Tag.Name).
		SetNillableSlug(req.Tag.Slug).
		SetNillableColor(req.Tag.Color).
		SetNillableThumbnail(req.Tag.Thumbnail).
		SetNillableSlugName(req.Tag.SlugName).
		SetNillablePostCount(req.Tag.PostCount).
		SetUpdateTime(time.Now().UnixMilli())

	res, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.convertEntToProto(res), err
}

func (r *TagRepo) Delete(ctx context.Context, req *v1.DeleteTagRequest) (bool, error) {
	err := r.data.db.Link.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	return err != nil, err
}
