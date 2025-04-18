package services

import (
	"context"
	"fmt"

	"github.com/adrisongomez/pti-ecommerce-site/backends/databases/db"
	"github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/http/svc_media/server"
	media "github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/svc_media"
	"github.com/adrisongomez/pti-ecommerce-site/backends/internal/utils"
	"github.com/adrisongomez/pti-ecommerce-site/backends/internal/utils/auth"
	mediaUtils "github.com/adrisongomez/pti-ecommerce-site/backends/pkg/utils"
	"go.uber.org/zap"
	"goa.design/clue/log"

	goaHttp "goa.design/goa/v3/http"
)

type MediaService struct {
	client *db.PrismaClient

	*zap.Logger
	*auth.JWTValidator
}

func MapMediaDBToOutput(model *db.MediaModel) *media.Media {
	size := int64(model.Size)
	output := media.Media{
		ID:        model.ID,
		MediaType: media.MediaType(model.Type),
		URL:       mediaUtils.GetResourceURL(model.Bucket, "us-east-1", model.Key),
		Filename:  model.Filename,
		Size:      size,
		MimeType:  model.MimeType,
		Bucket:    model.Bucket,
		Key:       model.Key,
		CreatedAt: model.CreatedAt.String(),
		UpdatedAt: nil,
	}
	if value, ok := model.UpdatedAt(); ok {
		output.UpdatedAt = utils.StringRef(value.String())
	}
	return &output
}

func (m *MediaService) count(ctx context.Context, payload *media.ListPayload) (int, error) {
	var rows []struct {
		Count db.BigInt `json:"count"`
	}
	err := m.client.Prisma.QueryRaw(
		"SELECT count(*) FROM project.medias WHERE bucket like CONCAT('%', $1, '%')",
		payload.Bucket,
	).Exec(ctx, &rows)

	if err != nil {
		return 0, err
	}

	if len(rows) == 0 {
		return 0, nil
	}

	count := int(rows[0].Count)
	return count, nil
}

func (m *MediaService) List(ctx context.Context, payload *media.ListPayload) (*media.MediaList, error) {
	m.Info("Media listed endpoint got called with", zap.Any("payload", payload))
	records, err := m.client.Media.FindMany(
		db.Media.Bucket.Contains(payload.Bucket),
	).Take(payload.PageSize).Skip(payload.After).Exec(ctx)
	if err != nil {
		return nil, err
	}

	mediaList := []*media.Media{}

	for _, record := range records {
		mediaList = append(mediaList, MapMediaDBToOutput(&record))
	}

	count, err := m.count(ctx, payload)

	if err != nil {
		return nil, err
	}

	nextPageCursor := utils.MinInt(count, payload.After+payload.PageSize)
	pageInfo := &media.PageInfo{
		StartCursor:   payload.After,
		EndCursor:     nextPageCursor,
		TotalResource: count,
		HasMore:       nextPageCursor < count,
	}

	response := &media.MediaList{
		Data:     mediaList,
		PageInfo: pageInfo,
	}
	return response, nil
}

func (m *MediaService) GetByID(ctx context.Context, payload *media.GetByIDPayload) (*media.Media, error) {
	mediaDB, err := m.client.Media.FindFirst(db.Media.ID.Equals(payload.MediaID)).Exec(ctx)
	if err != nil {
		if db.IsErrNotFound(err) {
			return nil, media.MakeNotFound(err)
		}
		return nil, err
	}
	return MapMediaDBToOutput(mediaDB), nil
}

func (m *MediaService) Create(ctx context.Context, input *media.CreatePayload) (*media.CreateMediaResponse, error) {
	payload := input.Input
	m.Info("Media create endpoint got called with", zap.Any("payload", payload))
	url, err := mediaUtils.CreateObjectOnS3(ctx, payload.Bucket, payload.Key, payload.Size)
	if err != nil {
		return nil, err
	}
	createdMedia, err := m.client.Media.CreateOne(
		db.Media.Filename.Set(payload.Filename),
		db.Media.Size.Set(db.BigInt(payload.Size)),
		db.Media.Type.Set(db.MediaType(mediaUtils.GetMediaTypeByMimeType(payload.MimeType))),
		db.Media.MimeType.Set(payload.MimeType),
		db.Media.Bucket.Set(payload.Bucket),
		db.Media.Key.Set(payload.Key),
	).Exec(ctx)

	if err != nil {
		if err, ok := db.IsErrUniqueConstraint(err); ok {
			return nil, media.MakeBadRequest(fmt.Errorf("Error on field %v", err))
		}
		return nil, err
	}
	response := media.CreateMediaResponse{
		Media:     MapMediaDBToOutput(createdMedia),
		UploadURL: url,
	}
	return &response, nil
}

func (m *MediaService) DeleteByID(ctx context.Context, payload *media.DeleteByIDPayload) (bool, error) {
	record, err := m.client.Media.FindUnique(db.Media.ID.Equals(payload.MediaID)).Delete().Exec(ctx)
	m.Info(fmt.Sprintf("Delete media record: %v", record))

	if err != nil {
		return false, err
	}

	return true, nil
}

func MountMediaSVC(mux goaHttp.Muxer, svc media.Service) {
	endpoints := media.NewEndpoints(svc)
	endpoints.Use(log.Endpoint)
	req := goaHttp.RequestDecoder
	res := goaHttp.ResponseEncoder
	handler := server.New(endpoints, mux, req, res, nil, nil)
	server.Mount(mux, handler)

	go func() {
		for _, mount := range handler.Mounts {
			zap.L().Info(fmt.Sprintf("%q mounted on %s %s\n", mount.Method, mount.Verb, mount.Pattern))
		}
	}()
}

func NewMediaService(client *db.PrismaClient, validator *auth.JWTValidator) media.Service {
	return &MediaService{client, zap.L(), validator}
}
