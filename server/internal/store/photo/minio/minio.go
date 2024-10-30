package minio

import (
	"bufio"
	"context"
	"io"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/subliker/server/internal/config"
	"github.com/subliker/server/internal/logger"
	"github.com/subliker/server/internal/store/photo"

	"github.com/google/uuid"
)

var publicBusketName = "public"

type minioStore struct {
	client *minio.Client
}

func New(cfg config.PhotoStore) photo.Store {
	minioClient, err := minio.New("minio:9000", &minio.Options{
		Creds: credentials.NewStaticV4(cfg.Access, cfg.Secret, ""),
	})
	if err != nil {
		logger.Zap.Fatal(err)
	}

	s := &minioStore{
		client: minioClient,
	}

	s.initBuckets()
	return s
}

func (s *minioStore) initBuckets() {
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

func (s *minioStore) Put(photoReader io.Reader, photoName string, photoSize int64) (string, error) {
	// fix photo name
	photoNameRunes := []rune(photoName)
	if len(photoNameRunes) > 20 {
		photoName = string(photoNameRunes[len(photoNameRunes)-20:])
	}

	// making uniq file name
	objectName := uuid.New().String() + photoName

	// put object into storage
	if _, err := s.client.PutObject(context.Background(), publicBusketName, objectName,
		photoReader, photoSize, minio.PutObjectOptions{}); err != nil {
		logger.Zap.Error(err)
		return "", err
	}

	return objectName, nil
}

func (s *minioStore) Get(name string) (string, error) {
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
