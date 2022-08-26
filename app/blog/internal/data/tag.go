package data

import (
	"context"
	"entgo.io/ent/entc/integration/edgeschema/ent/role"
	"kratos-blog/app/blog/internal/data/ent"
	"kratos-blog/app/blog/internal/data/ent/tag"
	"kratos-blog/pkg/util/entgo"
	paging "kratos-blog/pkg/util/pagination"

	"github.com/go-kratos/kratos/v2/log"

	v1 "kratos-blog/api/blog/v1"
	"kratos-blog/app/blog/internal/biz"
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
		Id:          in.ID,
		Name:        in.Name,
		DisplayName: in.DisplayName,
		SeoDesc:     in.SeoDesc,
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
	if req.Page != nil && req.PageSize != nil {
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
		Select(role.FieldID).
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
		SetNillableDisplayName(req.Tag.DisplayName).
		SetNillableSeoDesc(req.Tag.SeoDesc).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.convertEntToProto(po), err
}

func (r *TagRepo) Update(ctx context.Context, req *v1.UpdateTagRequest) (*v1.Tag, error) {
	builder := r.data.db.Tag.UpdateOneID(req.Id).
		SetNillableName(req.Tag.Name).
		SetNillableDisplayName(req.Tag.DisplayName).
		SetNillableSeoDesc(req.Tag.SeoDesc)

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
