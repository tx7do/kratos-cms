package data

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"

	"github.com/go-kratos/kratos/v2/log"
	entgo "github.com/tx7do/go-utils/entgo/query"
	util "github.com/tx7do/go-utils/time"

	"kratos-cms/app/core/service/internal/data/ent"
	"kratos-cms/app/core/service/internal/data/ent/attachment"

	pagination "github.com/tx7do/kratos-bootstrap/gen/api/go/pagination/v1"
	v1 "kratos-cms/gen/api/go/file/service/v1"
)

type AttachmentRepo struct {
	data *Data
	log  *log.Helper
}

func NewAttachmentRepo(data *Data, logger log.Logger) *AttachmentRepo {
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

func (r *AttachmentRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().Attachment.Query()
	if len(whereCond) != 0 {
		builder.Modify(whereCond...)
	}

	count, err := builder.Count(ctx)
	if err != nil {
		r.log.Errorf("query count failed: %s", err.Error())
	}

	return count, err
}

func (r *AttachmentRepo) List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListAttachmentResponse, error) {
	builder := r.data.db.Client().Attachment.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), attachment.FieldCreateTime)
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

	items := make([]*v1.Attachment, 0, len(results))
	for _, res := range results {
		item := r.convertEntToProto(res)
		items = append(items, item)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &v1.ListAttachmentResponse{
		Total: int32(count),
		Items: items,
	}, nil
}

func (r *AttachmentRepo) Get(ctx context.Context, req *v1.GetAttachmentRequest) (*v1.Attachment, error) {
	res, err := r.data.db.Client().Attachment.Get(ctx, req.GetId())
	if err != nil && !ent.IsNotFound(err) {
		r.log.Errorf("query one data failed: %s", err.Error())
		return nil, err
	}

	return r.convertEntToProto(res), err
}

func (r *AttachmentRepo) Create(ctx context.Context, req *v1.CreateAttachmentRequest) (*v1.Attachment, error) {
	res, err := r.data.db.Client().Attachment.Create().
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
		r.log.Errorf("insert one data failed: %s", err.Error())
		return nil, err
	}

	return r.convertEntToProto(res), err
}

func (r *AttachmentRepo) Update(ctx context.Context, req *v1.UpdateAttachmentRequest) (*v1.Attachment, error) {
	builder := r.data.db.Client().Attachment.UpdateOneID(req.Id).
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

	res, err := builder.Save(ctx)
	if err != nil {
		r.log.Errorf("update one data failed: %s", err.Error())
		return nil, err
	}

	return r.convertEntToProto(res), err
}

func (r *AttachmentRepo) Delete(ctx context.Context, req *v1.DeleteAttachmentRequest) (bool, error) {
	err := r.data.db.Client().Attachment.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	if err != nil {
		r.log.Errorf("delete one data failed: %s", err.Error())
	}

	return err == nil, err
}
