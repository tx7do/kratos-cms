package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"kratos-blog/app/content/service/internal/biz"
	"kratos-blog/app/content/service/internal/data/ent"
	"kratos-blog/app/content/service/internal/data/ent/link"

	v1 "kratos-blog/api/blog/v1"
	"kratos-blog/pkg/util/entgo"
	paging "kratos-blog/pkg/util/pagination"
	"kratos-blog/third_party/pagination"
)

var _ biz.LinkRepo = (*LinkRepo)(nil)

type LinkRepo struct {
	data *Data
	log  *log.Helper
}

func NewLinkRepo(data *Data, logger log.Logger) biz.LinkRepo {
	l := log.NewHelper(log.With(logger, "module", "link/repo"))
	return &LinkRepo{
		data: data,
		log:  l,
	}
}

func (r *LinkRepo) convertEntToProto(in *ent.Link) *v1.Link {
	if in == nil {
		return nil
	}
	return &v1.Link{
		Id:          in.ID,
		Name:        in.Name,
		Url:         in.URL,
		Logo:        in.Logo,
		Description: in.Description,
		Team:        in.Team,
		Priority:    in.Priority,
		CreateTime:  entgo.UnixMilliToStringPtr(in.CreateTime),
		UpdateTime:  entgo.UnixMilliToStringPtr(in.UpdateTime),
	}
}

func (r *LinkRepo) List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListLinkResponse, error) {
	whereCond, orderCond := entgo.QueryCommandToSelector(req.GetQuery(), req.GetOrderBy())

	builder1 := r.data.db.Link.Query()
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
		builder1.Order(ent.Desc(link.FieldCreateTime))
	}
	if req.GetPage() > 0 && req.GetPageSize() > 0 && !req.GetNopaging() {
		builder1.
			Offset(paging.GetPageOffset(req.GetPage(), req.GetPageSize())).
			Limit(int(req.GetPageSize()))
	}
	links, err := builder1.All(ctx)
	if err != nil {
		return nil, err
	}

	builder2 := r.data.db.Link.Query()
	if len(whereCond) != 0 {
		for _, v := range whereCond {
			builder2.Where(v)
		}
	}
	count, err := builder2.
		Select(link.FieldID).
		Count(ctx)
	if err != nil {
		return nil, err
	}

	items := make([]*v1.Link, 0, len(links))
	for _, po := range links {
		item := r.convertEntToProto(po)
		items = append(items, item)
	}

	ret := v1.ListLinkResponse{
		Total: int32(count),
		Items: items,
	}

	return &ret, err
}

func (r *LinkRepo) Get(ctx context.Context, req *v1.GetLinkRequest) (*v1.Link, error) {
	po, err := r.data.db.Link.Get(ctx, req.GetId())
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	return r.convertEntToProto(po), err
}

func (r *LinkRepo) Create(ctx context.Context, req *v1.CreateLinkRequest) (*v1.Link, error) {
	po, err := r.data.db.Link.Create().
		SetNillableName(req.Link.Name).
		SetNillableURL(req.Link.Url).
		SetNillableLogo(req.Link.Logo).
		SetNillableDescription(req.Link.Description).
		SetNillableTeam(req.Link.Team).
		SetNillablePriority(req.Link.Priority).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.convertEntToProto(po), err
}

func (r *LinkRepo) Update(ctx context.Context, req *v1.UpdateLinkRequest) (*v1.Link, error) {
	builder := r.data.db.Link.UpdateOneID(req.Id).
		SetNillableName(req.Link.Name).
		SetNillableURL(req.Link.Url).
		SetNillableLogo(req.Link.Logo).
		SetNillableDescription(req.Link.Description).
		SetNillableTeam(req.Link.Team).
		SetNillablePriority(req.Link.Priority)

	po, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.convertEntToProto(po), err
}

func (r *LinkRepo) Delete(ctx context.Context, req *v1.DeleteLinkRequest) (bool, error) {
	err := r.data.db.Link.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	return err != nil, err
}
