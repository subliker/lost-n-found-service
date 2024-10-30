package storage

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/subliker/server/internal/config"
	"github.com/subliker/server/internal/logger"
)

var publicBusketName = "public"

type Storage struct {
	client *minio.Client
}

func New(cfg config.StorageConfig) *Storage {
	minioClient, err := minio.New("minio:9000", &minio.Options{
		Creds: credentials.NewStaticV4(cfg.Access, cfg.Secret, ""),
	})
	if err != nil {
		logger.Zap.Fatal(err)
	}

	s := &Storage{
		client: minioClient,
	}

	s.initBuckets()
	return s
}

func (s *Storage) initBuckets() {
	exists, err := s.client.BucketExists(context.Background(), publicBusketName)
	if err != nil {
		logger.Zap.Fatal(err)
	}

	if !exists {
		err := s.client.MakeBucket(context.Background(), publicBusketName, minio.MakeBucketOptions{})
		if err != nil {
			logger.Zap.Fatal(err)
		}
	}
}

func (s *Storage) PutPublicPhoto(photoContent string) (string, error) {
	reader := strings.NewReader(photoContent)
	name := fmt.Sprintf("%d.jpg", time.Now().UnixNano())

	if _, err := s.client.PutObject(context.Background(), publicBusketName, name,
		reader, int64(len(photoContent)), minio.PutObjectOptions{}); err != nil {
		logger.Zap.Error(err)
		return "", err
	}

	return name, nil
}

func (s *Storage) GetPublicPhoto(name string) (string, error) {
	photo, err := s.client.GetObject(context.Background(), publicBusketName, name, minio.GetObjectOptions{})
	if err != nil {
		return "", err
	}
	defer photo.Close()

	_, err = photo.Stat()
	if err != nil {
		return "", err
	}

	reader := bufio.NewReader(photo)
	content, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
