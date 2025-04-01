package services

import (
	"context"
	"fmt"

	"github.com/adrisongomez/pti-ecommerce-site/backends/databases/db"
	"go.uber.org/zap"
	"goa.design/goa/v3/http"

	. "github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/address"
	"github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/http/address/server"
	"github.com/adrisongomez/pti-ecommerce-site/backends/internal/utils"
)

type AddressService struct {
	client *db.PrismaClient

	*zap.Logger
}

func MapAddressDBToOutput(model *db.AddressModel) *Address {
	output := &Address{
		ID:           model.ID,
		AddressLine1: model.AddressLine1,
		City:         model.City,
		Country:      model.Country,
		State:        model.Province,
		CreatedAt:    model.CreatedAt.String(),
	}

	if data, ok := model.UpdatedAt(); ok {
		output.UpdatedAt = utils.StringRef(data.String())
	}

	if data, ok := model.AddressLine2(); ok {
		output.AddressLine2 = utils.StringRef(data)
	}

	if data, ok := model.ZipCode(); ok {
		output.ZipCode = *utils.StringRef(data)
	}
	return output
}

func (a AddressService) Show(ctx context.Context, payload *ShowPayload) (*Address, error) {
	addressDB, err := a.client.Address.FindUnique(db.Address.ID.Equals(payload.AddressID)).Exec(ctx)
	if err != nil {
		if db.IsErrNotFound(err) {
			return nil, MakeNotFound(err)
		}
		return nil, err
	}
	return MapAddressDBToOutput(addressDB), nil
}

func (a AddressService) Create(ctx context.Context, payload *CreatePayload) (*Address, error) {
	changes := []db.AddressSetParam{
		db.Address.User.Link(db.User.ID.Equals(payload.Input.UserID)),
	}

	if payload.Input.AddressLine2 != nil {
		changes = append(changes, db.Address.AddressLine2.Set(*payload.Input.AddressLine2))
	}

	if payload.Input.ZipCode != nil {
		changes = append(changes, db.Address.ZipCode.Set(*payload.Input.ZipCode))
	}
	addressDB, err := a.client.Address.CreateOne(
		db.Address.AddressLine1.Set(payload.Input.AddressLine1),
		db.Address.City.Set(payload.Input.City),
		db.Address.Province.Set(payload.Input.State),
		db.Address.Country.Set(payload.Input.Country),
		changes...,
	).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return MapAddressDBToOutput(addressDB), nil
}

func (a AddressService) Delete(ctx context.Context, payload *DeletePayload) (bool, error) {
	_, err := a.client.Address.FindUnique(db.Address.ID.Equals(payload.AddressID)).Delete().Exec(ctx)
	return err != nil, nil
}

func (a AddressService) List(ctx context.Context, payload *ListPayload) (*AddressList, error) {
	addressDbs, err := a.client.Address.FindMany().Take(payload.PageSize).Skip(payload.After).Exec(ctx)
	if err != nil {
		return nil, err
	}
	addresses := []*Address{}
	for _, addressDB := range addressDbs {
		addresses = append(addresses, MapAddressDBToOutput(&addressDB))
	}
	var resp []struct {
		Count db.BigInt `json:"count"`
	}
	err = a.client.Prisma.QueryRaw("SELECT count(*) FROM project.orders").Exec(ctx, &resp)

	if err != nil {
		a.Error("Error count order", zap.Error(err))
		return nil, err
	}
	count := 0
	if len(resp) == 0 {
		count = int(resp[0].Count)
	}
	nextPageCursor := utils.MinInt(count, payload.After+payload.PageSize)
	response := &AddressList{
		Data: addresses,
		PageInfo: &PageInfo{
			StartCursor:   payload.After,
			EndCursor:     nextPageCursor,
			TotalResource: count,
			HasMore:       nextPageCursor < count,
		},
	}
	return response, nil
}

func NewAddressService(client *db.PrismaClient) Service {
	return &AddressService{client: client, Logger: zap.L()}
}

func MountAddressSVC(mux http.Muxer, svc Service) {
	endpoints := NewEndpoints(svc)
	req := http.RequestDecoder
	res := http.ResponseEncoder

	handler := server.New(endpoints, mux, req, res, nil, nil)
	server.Mount(mux, handler)

	go func() {
		for _, mount := range handler.Mounts {
			zap.L().Info(fmt.Sprintf("%q mounted on %s %s\n", mount.Method, mount.Verb, mount.Pattern))
		}
	}()
}
