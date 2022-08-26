package data

import (
	"context"
	"kratos-blog/app/blog/internal/data/ent/category"
	"kratos-blog/pkg/util/entgo"

	"entgo.io/ent/entc/integration/edgeschema/ent/role"
	"github.com/go-kratos/kratos/v2/log"

	v1 "kratos-blog/api/blog/v1"
	"kratos-blog/app/blog/internal/biz"
	"kratos-blog/app/blog/internal/data/ent"
	paging "kratos-blog/pkg/util/pagination"
	"kratos-blog/third_party/pagination"
)

var _ biz.CategoryRepo = (*CategoryRepo)(nil)

type CategoryRepo struct {
	data *Data
	log  *log.Helper
}

func NewCategoryRepo(data *Data, logger log.Logger) biz.CategoryRepo {
	l := log.NewHelper(log.With(logger, "module", "category/repo"))
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
		Id:       in.ID,
		Name:     in.Name,
		SeoDesc:  in.SeoDesc,
		ParentId: in.ParentID,
	}
}

func (r *CategoryRepo) List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListCategoryResponse, error) {
	whereCond, orderCond := entgo.QueryCommandToSelector(req.GetQuery(), req.GetOrderBy())

	builder1 := r.data.db.Category.Query()
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
		builder1.Order(ent.Desc(category.FieldCreateTime))
	}
	if req.Page != nil && req.PageSize != nil {
		builder1.
			Offset(paging.GetPageOffset(req.GetPage(), req.GetPageSize())).
			Limit(int(req.GetPageSize()))
	}
	categories, err := builder1.All(ctx)
	if err != nil {
		return nil, err
	}

	builder2 := r.data.db.Category.Query()
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

	items := make([]*v1.Category, 0, len(categories))
	for _, po := range categories {
		item := r.convertEntToProto(po)
		items = append(items, item)
	}

	ret := v1.ListCategoryResponse{
		Total: int32(count),
		Items: items,
	}

	return &ret, err
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
		SetNillableSeoDesc(req.Category.SeoDesc).
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
		SetNillableSeoDesc(req.Category.SeoDesc)

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
