package data

import (
	"context"
	"github.com/tx7do/kratos-utils/entgo"
	util "github.com/tx7do/kratos-utils/time"
	"kratos-cms/app/core/service/internal/data/ent"
	"kratos-cms/app/core/service/internal/data/ent/comment"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"kratos-cms/app/core/service/internal/biz"
	"kratos-cms/gen/api/go/common/pagination"
	v1 "kratos-cms/gen/api/go/file/service/v1"

	paging "github.com/tx7do/kratos-utils/pagination"
)

var _ biz.AttachmentRepo = (*AttachmentRepo)(nil)

type AttachmentRepo struct {
	data *Data
	log  *log.Helper
}

func NewAttachmentRepo(data *Data, logger log.Logger) biz.AttachmentRepo {
	l := log.NewHelper(log.With(logger, "module", "attachment/repo/core-service"))
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
		ThumbPath:  in.Thumbnail,
		MediaType:  in.MediaType,
		Suffix:     in.Suffix,
		Width:      in.Width,
		Height:     in.Height,
		Size:       in.Size,
		Type:       in.Type,
		CreateTime: util.UnixMilliToStringPtr(in.CreateTime),
		UpdateTime: util.UnixMilliToStringPtr(in.UpdateTime),
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
		builder1.Order(ent.Desc(comment.FieldCreateTime))
	}
	if req.GetPage() > 0 && req.GetPageSize() > 0 && !req.GetNopaging() {
		builder1.
			Offset(paging.GetPageOffset(req.GetPage(), req.GetPageSize())).
			Limit(int(req.GetPageSize()))
	}
	comments, err := builder1.All(ctx)
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
		Select(comment.FieldID).
		Count(ctx)
	if err != nil {
		return nil, err
	}

	items := make([]*v1.Attachment, 0, len(comments))
	for _, po := range comments {
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
		SetNillableThumbnail(req.Attachment.ThumbPath).
		SetNillableMediaType(req.Attachment.MediaType).
		SetNillableSuffix(req.Attachment.Suffix).
		SetNillableWidth(req.Attachment.Width).
		SetNillableHeight(req.Attachment.Height).
		SetNillableSize(req.Attachment.Size).
		SetNillableType(req.Attachment.Type).
		SetCreateTime(time.Now().UnixMilli()).
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
		SetNillableThumbnail(req.Attachment.ThumbPath).
		SetNillableMediaType(req.Attachment.MediaType).
		SetNillableSuffix(req.Attachment.Suffix).
		SetNillableWidth(req.Attachment.Width).
		SetNillableHeight(req.Attachment.Height).
		SetNillableSize(req.Attachment.Size).
		SetNillableType(req.Attachment.Type).
		SetUpdateTime(time.Now().UnixMilli())

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
