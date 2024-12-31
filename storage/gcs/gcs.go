package gcs

import (
	"cloud.google.com/go/storage"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/oliverrsebastian/common-utils-go/logger"
	"google.golang.org/api/option"
	"io"
)

type client struct {
	gcs    *storage.Client
	config *Config
}

func NewClient(config *Config) *client {
	var (
		gcs *storage.Client
		err error
	)
	if config.IsEnabled {
		projectKey, err := base64.StdEncoding.DecodeString(config.ProjectKey)
		if err != nil {
			panic(err)
		}
		gcs, err = storage.NewClient(context.Background(), option.WithCredentialsJSON(projectKey))
	} else {
		gcs, err = storage.NewClient(context.Background(), option.WithoutAuthentication())
	}

	if err != nil {
		panic(err)
	}

	return &client{gcs: gcs, config: config}
}

func (s *client) Close() error {
	return s.gcs.Close()
}

func (s *client) Upload(ctx context.Context, folder, file string, r io.Reader) (url string, err error) {

	if !s.config.IsEnabled {
		url = s.config.DefaultUrl
		return
	}

	bucket := s.gcs.Bucket(s.config.ProjectBucketName)
	_, err = bucket.Attrs(ctx)
	if err != nil {
		if err.Error() != NotExistError {
			return
		}

		if err = bucket.Create(ctx, s.config.ProjectID, &storage.BucketAttrs{
			Name:          s.config.ProjectBucketName,
			PredefinedACL: "publicRead",
		}); err != nil {
			return
		}
	}

	obj := bucket.Object(fmt.Sprintf("%v/%v", folder, file))
	w := obj.NewWriter(ctx)

	if _, err = io.Copy(w, r); err != nil {
		return
	}

	if err = w.Close(); err != nil {
		logger.Error(ctx, "got error when closing file %v, err: %v", file, err)
		return
	}

	url = fmt.Sprintf(PublicUrlFormat, s.config.ProjectBucketName, folder, file)

	return
}
