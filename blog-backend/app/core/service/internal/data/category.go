package data

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	entgo "github.com/tx7do/go-utils/entgo/query"
	util "github.com/tx7do/go-utils/time"

	"kratos-cms/app/core/service/internal/biz"
	"kratos-cms/app/core/service/internal/data/ent"
	"kratos-cms/app/core/service/internal/data/ent/category"

	"kratos-cms/gen/api/go/common/pagination"
	"kratos-cms/gen/api/go/content/service/v1"
)

var _ biz.CategoryRepo = (*CategoryRepo)(nil)

type CategoryRepo struct {
	data *Data
	log  *log.Helper
}

func NewCategoryRepo(data *Data, logger log.Logger) biz.CategoryRepo {
	l := log.NewHelper(log.With(logger, "module", "category/repo/core-service"))
	return &CategoryRepo{
		data: data,
		log:  l,
	}
}

func (r *CategoryRepo) convertEntToProto(in *ent.Category) *v1.Category {
	if in == nil {
		return nil
	}
	return &v1.Category{
		Id:          in.ID,
		ParentId:    in.ParentID,
		Name:        in.Name,
		Slug:        in.Slug,
		Description: in.Description,
		Thumbnail:   in.Thumbnail,
		Password:    in.Password,
		FullPath:    in.FullPath,
		Priority:    in.Priority,
		PostCount:   in.PostCount,
		CreateTime:  util.UnixMilliToStringPtr(in.CreateTime),
		UpdateTime:  util.UnixMilliToStringPtr(in.UpdateTime),
	}
}

func (r *CategoryRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().Category.Query()
	if len(whereCond) != 0 {
		builder.Modify(whereCond...)
	}

	count, err := builder.Count(ctx)
	if err != nil {
		r.log.Errorf("query count failed: %s", err.Error())
	}

	return count, err
}

func (r *CategoryRepo) List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListCategoryResponse, error) {
	builder := r.data.db.Client().Category.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(r.data.db.Driver().Dialect(),
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), category.FieldCreateTime)
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

	items := make([]*v1.Category, 0, len(results))
	for _, res := range results {
		item := r.convertEntToProto(res)
		items = append(items, item)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &v1.ListCategoryResponse{
		Total: int32(count),
		Items: items,
	}, nil
}

func (r *CategoryRepo) Get(ctx context.Context, req *v1.GetCategoryRequest) (*v1.Category, error) {
	res, err := r.data.db.Client().Category.Get(ctx, req.GetId())
	if err != nil && !ent.IsNotFound(err) {
		r.log.Errorf("query one data failed: %s", err.Error())
		return nil, err
	}

	return r.convertEntToProto(res), err
}

func (r *CategoryRepo) Create(ctx context.Context, req *v1.CreateCategoryRequest) (*v1.Category, error) {
	res, err := r.data.db.Client().Category.Create().
		SetNillableName(req.Category.Name).
		SetNillableParentID(req.Category.ParentId).
		SetNillableSlug(req.Category.Slug).
		SetNillableDescription(req.Category.Description).
		SetNillableThumbnail(req.Category.Thumbnail).
		SetNillablePassword(req.Category.Password).
		SetNillableFullPath(req.Category.FullPath).
		SetNillablePriority(req.Category.Priority).
		SetNillablePostCount(req.Category.PostCount).
		SetCreateTime(time.Now().UnixMilli()).
		Save(ctx)
	if err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return nil, err
	}

	return r.convertEntToProto(res), err
}

func (r *CategoryRepo) Update(ctx context.Context, req *v1.UpdateCategoryRequest) (*v1.Category, error) {
	builder := r.data.db.Client().Category.UpdateOneID(req.Id).
		SetNillableName(req.Category.Name).
		SetNillableParentID(req.Category.ParentId).
		SetNillableSlug(req.Category.Slug).
		SetNillableDescription(req.Category.Description).
		SetNillableThumbnail(req.Category.Thumbnail).
		SetNillablePassword(req.Category.Password).
		SetNillableFullPath(req.Category.FullPath).
		SetNillablePriority(req.Category.Priority).
		SetNillablePostCount(req.Category.PostCount).
		SetUpdateTime(time.Now().UnixMilli())

	res, err := builder.Save(ctx)
	if err != nil {
		r.log.Errorf("update one data failed: %s", err.Error())
		return nil, err
	}

	return r.convertEntToProto(res), err
}

func (r *CategoryRepo) Delete(ctx context.Context, req *v1.DeleteCategoryRequest) (bool, error) {
	err := r.data.db.Client().Category.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	if err != nil {
		r.log.Errorf("delete one data failed: %s", err.Error())
	}

	return err == nil, err
}
