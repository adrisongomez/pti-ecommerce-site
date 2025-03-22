package services

import (
	"context"
	"fmt"

	"github.com/adrisongomez/pti-ecommerce-site/backends/databases/db"
	svcHttp "github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/http/svc_media/server"
	media "github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/svc_media"

	goaHttp "goa.design/goa/v3/http"
)

type MediaService struct {
	client *db.PrismaClient
}

func (m *MediaService) createObject(ctx context.Context, media *db.MediaModel) error {
	return nil
}

func (m *MediaService) mapDBToOutput(model *db.MediaModel) *media.Media {
	return &media.Media{
		ID:        model.ID,
		MediaType: media.MediaType(model.Type),
		URL:       model.MimeType,
		Filename:  &model.Filename,
		Size:      &model.Size,
		MimeType:  &model.MimeType,
		Bucket:    &model.Bucket,
		Key:       &model.Bucket,
		CreatedAt: model.CreatedAt.String(),
		UpdatedAt: nil,
	}
}

func (m *MediaService) List(ctx context.Context, payload *media.ListPayload) (*media.MediaList, error) {
	return nil, nil
}

func (m *MediaService) GetByID(ctx context.Context, payload *media.GetByIDPayload) (*media.Media, error) {
	media, err := m.client.Media.FindFirst(db.Media.ID.Equals(*payload.MediaID)).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return m.mapDBToOutput(media), nil
}

func (m *MediaService) Create(ctx context.Context, payload *media.MediaInput) (*media.Media, error) {
	createdMedia, err := m.client.Media.CreateOne(
		db.Media.Filename.Set(payload.Filename),
		db.Media.Size.Set(payload.Size),
		db.Media.Type.Set(""),
		db.Media.MimeType.Set(payload.MimeType),
		db.Media.Bucket.Set(payload.Bucket),
		db.Media.Key.Set(payload.Key),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}
	return m.mapDBToOutput(createdMedia), nil
}

func (m *MediaService) DeleteByID(ctx context.Context, payload *media.DeleteByIDPayload) (bool, error) {
	return false, nil
}

func MountMediaSVC(mux *goaHttp.Muxer, svc *MediaService) {
	endpoints := media.NewEndpoints(svc)
	req := goaHttp.RequestDecoder
	res := goaHttp.ResponseEncoder
	handler := svcHttp.New(endpoints, *mux, req, res, nil, nil)
	svcHttp.Mount(*mux, handler)

	go func() {
		for _, mount := range handler.Mounts {
			fmt.Printf("%q mounted on %s %s\n", mount.Method, mount.Verb, mount.Pattern)
		}
	}()
}

func NewMediaService(client *db.PrismaClient) *MediaService {
	return &MediaService{client}
}
