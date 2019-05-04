package app

import (
	"github.com/librarios/librc/minio"
)

type Upload struct {
	opt *UploadOption
}

type UploadOption struct {
	MinioConf         *minio.Conf
	DefaultBucketName string
}

func NewUpload(opt *UploadOption) *Upload {
	return &Upload{
		opt: opt,
	}
}

func (c *Upload) Upload(filePath, bucketName, objectName string) (int64, error) {
	if len(bucketName) == 0 {
		bucketName = c.opt.DefaultBucketName
	}
	return minio.Upload(c.opt.MinioConf, bucketName, filePath, objectName)
}
