package data

import (
	"context"
	"mime"
	"net/url"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/tx7do/go-utils/trans"

	conf "github.com/tx7do/kratos-bootstrap/api/gen/go/conf/v1"
	ossMinio "github.com/tx7do/kratos-bootstrap/oss/minio"

	fileV1 "kratos-cms/api/gen/go/file/service/v1"
)

const (
	defaultExpiryTime = time.Minute * 60 // 1小时
)

type MinIOClient struct {
	mc   *minio.Client
	conf *conf.OSS
	log  *log.Helper
}

func NewMinIoClient(cfg *conf.Bootstrap, logger log.Logger) *MinIOClient {
	l := log.NewHelper(log.With(logger, "module", "minio/data/core-service"))
	return &MinIOClient{
		log:  l,
		conf: cfg.Oss,
		mc:   ossMinio.NewClient(cfg.Oss),
	}
}

func (c *MinIOClient) contentTypeToBucketName(contentType string) string {
	h := strings.Split(contentType, "/")
	if len(h) != 2 {
		return "images"
	}

	var bucketName string
	switch h[0] {
	case "image":
		bucketName = "images"
		break
	case "video":
		bucketName = "videos"
		break
	case "audio":
		bucketName = "audios"
		break
	case "application", "text":
		bucketName = "docs"
		break
	default:
		bucketName = "images"
		break
	}

	return bucketName
}

func (c *MinIOClient) contentTypeToFileSuffix(contentType string) string {
	fileSuffix := ""

	switch contentType {
	case "text/plain":
		fileSuffix = ".txt"
		break

	case "image/jpeg":
		fileSuffix = ".jpg"
		break

	case "image/png":
		fileSuffix = ".png"
		break

	default:
		extensions, _ := mime.ExtensionsByType(contentType)
		if len(extensions) > 0 {
			fileSuffix = extensions[0]
		}
	}

	return fileSuffix
}

func (c *MinIOClient) jointObjectName(contentType string, filePath, fileName *string) (string, string) {
	fileSuffix := c.contentTypeToFileSuffix(contentType)

	var _fileName string
	if fileName == nil {
		_fileName = uuid.New().String() + fileSuffix
	} else {
		_fileName = *fileName
	}

	var objectName string
	if filePath != nil {
		objectName = *filePath + "/" + _fileName
	} else {
		objectName = _fileName
	}

	return objectName, _fileName
}

func (c *MinIOClient) OssUploadUrl(ctx context.Context, req *fileV1.OssUploadUrlRequest) (*fileV1.OssUploadUrlResponse, error) {
	var bucketName string
	if req.BucketName != nil {
		bucketName = req.GetBucketName()
	} else {
		bucketName = c.contentTypeToBucketName(req.GetContentType())
	}

	objectName, _ := c.jointObjectName(req.GetContentType(), req.FilePath, req.FileName)

	expiry := defaultExpiryTime

	var uploadUrl string
	var downloadUrl string
	var formData map[string]string

	var err error
	var presignedURL *url.URL

	switch req.GetMethod() {
	case fileV1.UploadMethod_Put:
		presignedURL, err = c.mc.PresignedPutObject(ctx, bucketName, objectName, expiry)
		if err != nil {
			return nil, err
		}

		uploadUrl = presignedURL.String()
		uploadUrl = strings.Replace(uploadUrl, c.conf.Minio.Endpoint, c.conf.Minio.UploadHost, -1)

		downloadUrl = presignedURL.Scheme + "://" + presignedURL.Host + "/" + bucketName + "/" + objectName
		downloadUrl = strings.Replace(downloadUrl, c.conf.Minio.Endpoint, c.conf.Minio.DownloadHost, -1)

	case fileV1.UploadMethod_Post:
		policy := minio.NewPostPolicy()
		_ = policy.SetBucket(bucketName)
		_ = policy.SetKey(objectName)
		_ = policy.SetExpires(time.Now().UTC().Add(expiry))
		_ = policy.SetContentType(req.GetContentType())

		presignedURL, formData, err = c.mc.PresignedPostPolicy(ctx, policy)
		if err != nil {
			return nil, err
		}

		uploadUrl = presignedURL.String()
		uploadUrl = strings.Replace(uploadUrl, c.conf.Minio.Endpoint, c.conf.Minio.UploadHost, -1)

		downloadUrl = presignedURL.Scheme + "://" + presignedURL.Host + "/" + bucketName + "/" + objectName
		downloadUrl = strings.Replace(downloadUrl, c.conf.Minio.Endpoint, c.conf.Minio.DownloadHost, -1)
	}

	return &fileV1.OssUploadUrlResponse{
		UploadUrl:   uploadUrl,
		DownloadUrl: downloadUrl,
		ObjectName:  objectName,
		BucketName:  trans.Ptr(bucketName),
		FormData:    formData,
	}, nil
}
