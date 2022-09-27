package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"kratos-blog/app/content/service/internal/biz"
	"kratos-blog/app/content/service/internal/data/ent"
	"kratos-blog/app/content/service/internal/data/ent/tag"

	v1 "kratos-blog/api/blog/v1"
	"kratos-blog/pkg/util/entgo"
	paging "kratos-blog/pkg/util/pagination"
	"kratos-blog/third_party/pagination"
)

var _ biz.TagRepo = (*TagRepo)(nil)

type TagRepo struct {
	data *Data
	log  *log.Helper
}

func NewTagRepo(data *Data, logger log.Logger) biz.TagRepo {
	l := log.NewHelper(log.With(logger, "module", "tag/repo"))
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
		CreateTime: entgo.UnixMilliToStringPtr(in.CreateTime),
		UpdateTime: entgo.UnixMilliToStringPtr(in.UpdateTime),
	}
}

func (r *TagRepo) List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListTagResponse, error) {
	whereCond, orderCond := entgo.QueryCommandToSelector(req.GetQuery(), req.GetOrderBy())

	builder1 := r.data.db.Tag.Query()
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
		builder1.Order(ent.Desc(tag.FieldCreateTime))
	}
	if req.GetPage() > 0 && req.GetPageSize() > 0 && !req.GetNopaging() {
		builder1.
			Offset(paging.GetPageOffset(req.GetPage(), req.GetPageSize())).
			Limit(int(req.GetPageSize()))
	}
	tags, err := builder1.All(ctx)
	if err != nil {
		return nil, err
	}

	builder2 := r.data.db.Tag.Query()
	if len(whereCond) != 0 {
		for _, v := range whereCond {
			builder2.Where(v)
		}
	}
	count, err := builder2.
		Select(tag.FieldID).
		Count(ctx)
	if err != nil {
		return nil, err
	}

	items := make([]*v1.Tag, 0, len(tags))
	for _, po := range tags {
		item := r.convertEntToProto(po)
		items = append(items, item)
	}

	ret := v1.ListTagResponse{
		Total: int32(count),
		Items: items,
	}

	return &ret, err
}

func (r *TagRepo) Get(ctx context.Context, req *v1.GetTagRequest) (*v1.Tag, error) {
	po, err := r.data.db.Tag.Get(ctx, req.GetId())
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	return r.convertEntToProto(po), err
}

func (r *TagRepo) Create(ctx context.Context, req *v1.CreateTagRequest) (*v1.Tag, error) {
	po, err := r.data.db.Tag.Create().
		SetNillableName(req.Tag.Name).
		SetNillableSlug(req.Tag.Slug).
		SetNillableColor(req.Tag.Color).
		SetNillableThumbnail(req.Tag.Thumbnail).
		SetNillableSlugName(req.Tag.SlugName).
		SetNillablePostCount(req.Tag.PostCount).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.convertEntToProto(po), err
}

func (r *TagRepo) Update(ctx context.Context, req *v1.UpdateTagRequest) (*v1.Tag, error) {
	builder := r.data.db.Tag.UpdateOneID(req.Id).
		SetNillableName(req.Tag.Name).
		SetNillableSlug(req.Tag.Slug).
		SetNillableColor(req.Tag.Color).
		SetNillableThumbnail(req.Tag.Thumbnail).
		SetNillableSlugName(req.Tag.SlugName).
		SetNillablePostCount(req.Tag.PostCount)

	po, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.convertEntToProto(po), err
}

func (r *TagRepo) Delete(ctx context.Context, req *v1.DeleteTagRequest) (bool, error) {
	err := r.data.db.Link.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	return err != nil, err
}
