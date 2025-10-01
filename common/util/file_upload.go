package util

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

type MinoObject struct {
	Context    context.Context
	Client     *minio.Client
	Bucket     string
	Path       string
	FileHeader *multipart.FileHeader
}

func (m *MinoObject) UploadToS3() (*minio.UploadInfo, error) {
	file, err := m.FileHeader.Open()
	if err != nil {
		return nil, err
	}

	defer file.Close()

	id := uuid.New()
	ext := filepath.Ext(m.FileHeader.Filename)
	objectName := fmt.Sprintf("%s/%s%s", m.Path, id, ext)

	info, err := m.Client.PutObject(
		m.Context,
		m.Bucket,
		objectName,
		file,
		m.FileHeader.Size,
		minio.PutObjectOptions{
			ContentType: m.FileHeader.Header.Get("Content-Type"),
		},
	)

	if err != nil {
		return nil, err
	}
	return &info, nil
}
