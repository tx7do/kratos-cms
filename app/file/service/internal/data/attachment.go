package data

import (
	"context"
	"gorm.io/gorm"
	"kratos-blog/app/file/service/internal/data/model"
	paging "kratos-blog/pkg/util/pagination"

	"github.com/go-kratos/kratos/v2/log"

	"kratos-blog/app/file/service/internal/biz"

	v1 "kratos-blog/api/file/service/v1"
	"kratos-blog/third_party/pagination"
)

var _ biz.AttachmentRepo = (*AttachmentRepo)(nil)

type AttachmentRepo struct {
	data *Data
	log  *log.Helper
}

func NewAttachmentRepo(data *Data, logger log.Logger) biz.AttachmentRepo {
	l := log.NewHelper(log.With(logger, "module", "attachment/repo/file-service"))
	return &AttachmentRepo{
		data: data,
		log:  l,
	}
}

func (r *AttachmentRepo) List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListAttachmentResponse, error) {
	var whereCond, orderCond string

	var res v1.ListAttachmentResponse
	res.Items = make([]*v1.Attachment, 0)

	var builder *gorm.DB
	if len(orderCond) > 0 {
		builder = r.data.db.Order(orderCond)
	}
	if req.GetPage() > 0 && req.GetPageSize() > 0 && !req.GetNopaging() {
		builder.
			Offset(paging.GetPageOffset(req.GetPage(), req.GetPageSize())).
			Limit(int(req.GetPageSize()))
	}

	if len(whereCond) > 0 {

	}

	builder.Find(&res.Items)

	var countBuilder *gorm.DB
	var total int64
	countBuilder.Count(&total)

	res.Total = int32(total)

	return &res, nil
}

func (r *AttachmentRepo) Get(ctx context.Context, req *v1.GetAttachmentRequest) (*v1.Attachment, error) {
	var attachment v1.Attachment
	result := r.data.db.First(&attachment, req.GetId())
	if result.Error != nil {
		return nil, result.Error
	}

	return &attachment, nil
}

func (r *AttachmentRepo) Create(ctx context.Context, req *v1.CreateAttachmentRequest) (*v1.Attachment, error) {
	result := r.data.db.Create(req.Attachment)
	if result.Error != nil {
		return nil, result.Error
	}

	var attachment v1.Attachment
	result = r.data.db.First(&attachment, req.Attachment.GetId())
	if result.Error != nil {
		return nil, result.Error
	}

	return &attachment, nil
}

func (r *AttachmentRepo) Update(ctx context.Context, req *v1.UpdateAttachmentRequest) (*v1.Attachment, error) {
	result := r.data.db.Model(&model.Attachment{}).Updates(req.Attachment)
	if result.Error != nil {
		return nil, result.Error
	}

	var attachment v1.Attachment
	result = r.data.db.First(&attachment, req.Attachment.GetId())
	if result.Error != nil {
		return nil, result.Error
	}

	return nil, nil
}

func (r *AttachmentRepo) Delete(ctx context.Context, req *v1.DeleteAttachmentRequest) (bool, error) {
	result := r.data.db.Delete(&model.Attachment{}, req.GetId())
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
