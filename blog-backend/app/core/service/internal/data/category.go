package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	"kratos-cms/gen/api/go/common/pagination"
	"kratos-cms/gen/api/go/content/service/v1"

	"kratos-cms/app/core/service/internal/biz"
	"kratos-cms/app/core/service/internal/data/ent"
	"kratos-cms/app/core/service/internal/data/ent/category"

	"github.com/tx7do/kratos-utils/entgo"
	paging "github.com/tx7do/kratos-utils/pagination"
	util "github.com/tx7do/kratos-utils/time"
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

func (r *CategoryRepo) Count(ctx context.Context, whereCond entgo.WhereConditions) (int, error) {
	builder := r.data.db.Category.Query()
	if len(whereCond) != 0 {
		for _, cond := range whereCond {
			builder = builder.Where(cond)
		}
	}
	return builder.Count(ctx)
}

func (r *CategoryRepo) List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListCategoryResponse, error) {
	whereCond, orderCond := entgo.QueryCommandToSelector(req.GetQuery(), req.GetOrderBy())

	builder := r.data.db.Category.Query()
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
		builder.Order(ent.Desc(category.FieldCreateTime))
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

	items := make([]*v1.Category, 0, len(results))
	for _, res := range results {
		item := r.convertEntToProto(res)
		items = append(items, item)
	}

	count, err := r.Count(ctx, whereCond)
	if err != nil {
		return nil, err
	}

	return &v1.ListCategoryResponse{
		Total: int32(count),
		Items: items,
	}, nil
}

func (r *CategoryRepo) Get(ctx context.Context, req *v1.GetCategoryRequest) (*v1.Category, error) {
	po, err := r.data.db.Category.Get(ctx, req.GetId())
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	return r.convertEntToProto(po), err
}

func (r *CategoryRepo) Create(ctx context.Context, req *v1.CreateCategoryRequest) (*v1.Category, error) {
	po, err := r.data.db.Category.Create().
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
		return nil, err
	}

	return r.convertEntToProto(po), err
}

func (r *CategoryRepo) Update(ctx context.Context, req *v1.UpdateCategoryRequest) (*v1.Category, error) {
	builder := r.data.db.Category.UpdateOneID(req.Id).
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

	po, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.convertEntToProto(po), err
}

func (r *CategoryRepo) Delete(ctx context.Context, req *v1.DeleteCategoryRequest) (bool, error) {
	err := r.data.db.Category.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	return err != nil, err
}
