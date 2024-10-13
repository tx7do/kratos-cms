package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	fileV1 "kratos-cms/api/gen/go/file/service/v1"
	frontV1 "kratos-cms/api/gen/go/front/service/v1"
)

type FileService struct {
	frontV1.FileServiceHTTPServer

	sc fileV1.FileServiceClient

	log *log.Helper
}

func NewFileService(logger log.Logger, sc fileV1.FileServiceClient) *FileService {
	l := log.NewHelper(log.With(logger, "module", "file/service/front-service"))
	return &FileService{
		log: l,
		sc:  sc,
	}
}

func (s *FileService) OssUploadUrl(ctx context.Context, req *fileV1.OssUploadUrlRequest) (*fileV1.OssUploadUrlResponse, error) {
	return s.sc.OssUploadUrl(ctx, req)
}
