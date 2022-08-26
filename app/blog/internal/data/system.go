package data

import (
	"context"
	"entgo.io/ent/entc/integration/edgeschema/ent/role"
	"github.com/go-kratos/kratos/v2/log"
	v1 "kratos-blog/api/blog/v1"
	"kratos-blog/app/blog/internal/biz"
	"kratos-blog/app/blog/internal/data/ent"
	"kratos-blog/app/blog/internal/data/ent/system"
	"kratos-blog/pkg/util/entgo"
	paging "kratos-blog/pkg/util/pagination"
	"kratos-blog/third_party/pagination"
)

var _ biz.SystemRepo = (*SystemRepo)(nil)

type SystemRepo struct {
	data *Data
	log  *log.Helper
}

func NewSystemRepo(data *Data, logger log.Logger) biz.SystemRepo {
	l := log.NewHelper(log.With(logger, "module", "system/repo"))
	return &SystemRepo{
		data: data,
		log:  l,
	}
}

func (r *SystemRepo) convertEntToProto(in *ent.System) *v1.System {
	if in == nil {
		return nil
	}
	return &v1.System{
		Id:           in.ID,
		Theme:        in.Theme,
		Title:        in.Title,
		Keywords:     in.Keywords,
		Description:  in.Description,
		RecordNumber: in.RecordNumber,
	}
}

func (r *SystemRepo) List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListSystemResponse, error) {
	whereCond, orderCond := entgo.QueryCommandToSelector(req.GetQuery(), req.GetOrderBy())

	builder1 := r.data.db.System.Query()
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
		builder1.Order(ent.Desc(system.FieldCreateTime))
	}
	if req.Page != nil && req.PageSize != nil {
		builder1.
			Offset(paging.GetPageOffset(req.GetPage(), req.GetPageSize())).
			Limit(int(req.GetPageSize()))
	}
	systems, err := builder1.All(ctx)
	if err != nil {
		return nil, err
	}

	builder2 := r.data.db.System.Query()
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

	items := make([]*v1.System, 0, len(systems))
	for _, po := range systems {
		item := r.convertEntToProto(po)
		items = append(items, item)
	}

	ret := v1.ListSystemResponse{
		Total: int32(count),
		Items: items,
	}

	return &ret, err
}

func (r *SystemRepo) Get(ctx context.Context, req *v1.GetSystemRequest) (*v1.System, error) {
	po, err := r.data.db.System.Get(ctx, req.GetId())
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	return r.convertEntToProto(po), err
}

func (r *SystemRepo) Create(ctx context.Context, req *v1.CreateSystemRequest) (*v1.System, error) {
	po, err := r.data.db.System.Create().
		SetTheme(*req.System.Theme).
		SetNillableTitle(req.System.Title).
		SetNillableKeywords(req.System.Keywords).
		SetNillableDescription(req.System.Description).
		SetNillableRecordNumber(req.System.RecordNumber).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.convertEntToProto(po), err
}

func (r *SystemRepo) Update(ctx context.Context, req *v1.UpdateSystemRequest) (*v1.System, error) {
	builder := r.data.db.System.UpdateOneID(req.Id).
		SetNillableTitle(req.System.Title).
		SetNillableKeywords(req.System.Keywords).
		SetNillableDescription(req.System.Description).
		SetNillableRecordNumber(req.System.RecordNumber)

	if req.System.Theme != nil {
		builder.SetTheme(*req.System.Theme)
	}

	po, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.convertEntToProto(po), err
}

func (r *SystemRepo) Delete(ctx context.Context, req *v1.DeleteSystemRequest) (bool, error) {
	err := r.data.db.Link.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	return err != nil, err
}
