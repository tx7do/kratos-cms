package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"kratos-blog/gen/api/go/common/pagination"
	"kratos-blog/gen/api/go/content/service/v1"

	"kratos-blog/app/content/service/internal/biz"
	"kratos-blog/app/content/service/internal/data/ent"
	"kratos-blog/app/content/service/internal/data/ent/category"

	"kratos-blog/pkg/util/entgo"
	paging "kratos-blog/pkg/util/pagination"
)

var _ biz.CategoryRepo = (*CategoryRepo)(nil)

type CategoryRepo struct {
	data *Data
	log  *log.Helper
}

func NewCategoryRepo(data *Data, logger log.Logger) biz.CategoryRepo {
	l := log.NewHelper(log.With(logger, "module", "category/repo/content-service"))
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
		CreateTime:  entgo.UnixMilliToStringPtr(in.CreateTime),
		UpdateTime:  entgo.UnixMilliToStringPtr(in.UpdateTime),
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
	if req.GetPage() > 0 && req.GetPageSize() > 0 && !req.GetNopaging() {
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
		Select(category.FieldID).
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
		SetNillableSlug(req.Category.Slug).
		SetNillableDescription(req.Category.Description).
		SetNillableThumbnail(req.Category.Thumbnail).
		SetNillablePassword(req.Category.Password).
		SetNillableFullPath(req.Category.FullPath).
		SetNillablePriority(req.Category.Priority).
		SetNillablePostCount(req.Category.PostCount).
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
		SetNillablePostCount(req.Category.PostCount)

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
