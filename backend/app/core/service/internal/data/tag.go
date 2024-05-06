package data

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"

	entgo "github.com/tx7do/go-utils/entgo/query"
	util "github.com/tx7do/go-utils/timeutil"

	"kratos-cms/app/core/service/internal/data/ent"
	"kratos-cms/app/core/service/internal/data/ent/tag"

	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	v1 "kratos-cms/gen/api/go/content/service/v1"
)

type TagRepo struct {
	data *Data
	log  *log.Helper
}

func NewTagRepo(data *Data, logger log.Logger) *TagRepo {
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

func (r *TagRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().Tag.Query()
	if len(whereCond) != 0 {
		builder.Modify(whereCond...)
	}

	count, err := builder.Count(ctx)
	if err != nil {
		r.log.Errorf("query count failed: %s", err.Error())
	}

	return count, err
}

func (r *TagRepo) List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListTagResponse, error) {
	builder := r.data.db.Client().Tag.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), tag.FieldCreateTime,
		req.GetFieldMask().GetPaths(),
	)
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

	items := make([]*v1.Tag, 0, len(results))
	for _, res := range results {
		item := r.convertEntToProto(res)
		items = append(items, item)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &v1.ListTagResponse{
		Total: int32(count),
		Items: items,
	}, nil
}

func (r *TagRepo) Get(ctx context.Context, req *v1.GetTagRequest) (*v1.Tag, error) {
	res, err := r.data.db.Client().Tag.Get(ctx, req.GetId())
	if err != nil && !ent.IsNotFound(err) {
		r.log.Errorf("query one data failed: %s", err.Error())
		return nil, err
	}

	return r.convertEntToProto(res), err
}

func (r *TagRepo) Create(ctx context.Context, req *v1.CreateTagRequest) (*v1.Tag, error) {
	res, err := r.data.db.Client().Tag.Create().
		SetNillableName(req.Tag.Name).
		SetNillableSlug(req.Tag.Slug).
		SetNillableColor(req.Tag.Color).
		SetNillableThumbnail(req.Tag.Thumbnail).
		SetNillableSlugName(req.Tag.SlugName).
		SetNillablePostCount(req.Tag.PostCount).
		SetCreateTime(time.Now().UnixMilli()).
		Save(ctx)
	if err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return nil, err
	}

	return r.convertEntToProto(res), err
}

func (r *TagRepo) Update(ctx context.Context, req *v1.UpdateTagRequest) (*v1.Tag, error) {
	builder := r.data.db.Client().Tag.UpdateOneID(req.Id).
		SetNillableName(req.Tag.Name).
		SetNillableSlug(req.Tag.Slug).
		SetNillableColor(req.Tag.Color).
		SetNillableThumbnail(req.Tag.Thumbnail).
		SetNillableSlugName(req.Tag.SlugName).
		SetNillablePostCount(req.Tag.PostCount).
		SetUpdateTime(time.Now().UnixMilli())

	res, err := builder.Save(ctx)
	if err != nil {
		r.log.Errorf("update one data failed: %s", err.Error())
		return nil, err
	}

	return r.convertEntToProto(res), err
}

func (r *TagRepo) Delete(ctx context.Context, req *v1.DeleteTagRequest) (bool, error) {
	err := r.data.db.Client().Tag.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	if err != nil {
		r.log.Errorf("delete one data failed: %s", err.Error())
	}

	return err == nil, err
}
