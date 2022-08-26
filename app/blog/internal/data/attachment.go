package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "kratos-blog/api/blog/v1"
	"kratos-blog/app/blog/internal/biz"
	"kratos-blog/app/blog/internal/data/ent"
	"kratos-blog/app/blog/internal/data/ent/link"
	"kratos-blog/pkg/util/entgo"
	paging "kratos-blog/pkg/util/pagination"
	"kratos-blog/third_party/pagination"
)

var _ biz.AttachmentRepo = (*AttachmentRepo)(nil)

type AttachmentRepo struct {
	data *Data
	log  *log.Helper
}

func NewAttachmentRepo(data *Data, logger log.Logger) biz.AttachmentRepo {
	l := log.NewHelper(log.With(logger, "module", "attachment/repo"))
	return &AttachmentRepo{
		data: data,
		log:  l,
	}
}

func (r *AttachmentRepo) convertEntToProto(in *ent.Attachment) *v1.Attachment {
	if in == nil {
		return nil
	}
	return &v1.Attachment{
		Id:         in.ID,
		Name:       in.Name,
		Path:       in.Path,
		FileKey:    in.FileKey,
		ThumbPath:  in.ThumbPath,
		MediaType:  in.MediaType,
		Suffix:     in.Suffix,
		Width:      in.Width,
		Height:     in.Height,
		Size:       in.Size,
		Type:       in.Type,
		CreateTime: entgo.UnixMilliToStringPtr(in.CreateTime),
	}
}

func (r *AttachmentRepo) List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListAttachmentResponse, error) {
	whereCond, orderCond := entgo.QueryCommandToSelector(req.GetQuery(), req.GetOrderBy())

	builder1 := r.data.db.Attachment.Query()
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
	if req.Page != nil && req.PageSize != nil {
		builder1.
			Offset(paging.GetPageOffset(req.GetPage(), req.GetPageSize())).
			Limit(int(req.GetPageSize()))
	}
	links, err := builder1.All(ctx)
	if err != nil {
		return nil, err
	}

	builder2 := r.data.db.Attachment.Query()
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

	items := make([]*v1.Attachment, 0, len(links))
	for _, po := range links {
		item := r.convertEntToProto(po)
		items = append(items, item)
	}

	ret := v1.ListAttachmentResponse{
		Total: int32(count),
		Items: items,
	}

	return &ret, err
}

func (r *AttachmentRepo) Get(ctx context.Context, req *v1.GetAttachmentRequest) (*v1.Attachment, error) {
	po, err := r.data.db.Attachment.Get(ctx, req.GetId())
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	return r.convertEntToProto(po), err
}

func (r *AttachmentRepo) Create(ctx context.Context, req *v1.CreateAttachmentRequest) (*v1.Attachment, error) {
	po, err := r.data.db.Attachment.Create().
		SetNillableName(req.Attachment.Name).
		SetNillablePath(req.Attachment.Path).
		SetNillableFileKey(req.Attachment.FileKey).
		SetNillableThumbPath(req.Attachment.ThumbPath).
		SetNillableMediaType(req.Attachment.MediaType).
		SetNillableSuffix(req.Attachment.Suffix).
		SetNillableWidth(req.Attachment.Width).
		SetNillableHeight(req.Attachment.Height).
		SetNillableSize(req.Attachment.Size).
		SetNillableType(req.Attachment.Type).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.convertEntToProto(po), err
}

func (r *AttachmentRepo) Update(ctx context.Context, req *v1.UpdateAttachmentRequest) (*v1.Attachment, error) {
	builder := r.data.db.Attachment.UpdateOneID(req.Id).
		SetNillableName(req.Attachment.Name).
		SetNillablePath(req.Attachment.Path).
		SetNillableFileKey(req.Attachment.FileKey).
		SetNillableThumbPath(req.Attachment.ThumbPath).
		SetNillableMediaType(req.Attachment.MediaType).
		SetNillableSuffix(req.Attachment.Suffix).
		SetNillableWidth(req.Attachment.Width).
		SetNillableHeight(req.Attachment.Height).
		SetNillableSize(req.Attachment.Size).
		SetNillableType(req.Attachment.Type)

	po, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.convertEntToProto(po), err
}

func (r *AttachmentRepo) Delete(ctx context.Context, req *v1.DeleteAttachmentRequest) (bool, error) {
	err := r.data.db.Attachment.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	return err != nil, err
}
