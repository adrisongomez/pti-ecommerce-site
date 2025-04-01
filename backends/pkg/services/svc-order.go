package services

import (
	"context"
	"fmt"
	"time"

	"github.com/adrisongomez/pti-ecommerce-site/backends/databases/db"
	"github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/http/order/server"
	. "github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/order"
	"github.com/adrisongomez/pti-ecommerce-site/backends/internal/utils"
	"github.com/adrisongomez/pti-ecommerce-site/backends/internal/utils/auth"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"goa.design/goa/v3/http"
)

type OrderService struct {
	client *db.PrismaClient

	*auth.JWTValidator
	*zap.Logger
}

const DecimalSize = -2

var Connection = []db.OrderRelationWith{
	db.Order.Address.Fetch(),
	db.Order.LineItems.Fetch(),
	db.Order.User.Fetch(),
}

func (o *OrderService) Create(ctx context.Context, payload *CreatePayload) (*Order, error) {
	o.Info("Create Order got called with", zap.Any("payload", payload))

	orderDB, err := o.client.Order.CreateOne(
		db.Order.UserEmail.Set(payload.Input.Email),
		db.Order.TotalPrice.Set(decimal.New(int64(payload.Input.TotalPrice), DecimalSize)),
		db.Order.User.Link(db.User.ID.Equals(payload.Input.UserID)),
		db.Order.Address.Link(db.Address.ID.Equals(payload.Input.AddressID)),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	tx := []db.PrismaTransaction{}

	for _, lineItem := range payload.Input.LineItems {
		tx = append(tx,
			o.client.OrderLineItem.CreateOne(
				db.OrderLineItem.Price.Set(decimal.New(int64(lineItem.Price), DecimalSize)),
				db.OrderLineItem.Quantity.Set(lineItem.Quantity),
				db.OrderLineItem.Product.Link(db.Product.ID.Equals(lineItem.ProductID)),
				db.OrderLineItem.Order.Link(db.Order.ID.Equals(orderDB.ID)),
			).Tx(),
		)

	}

	if err = o.client.Prisma.Transaction(tx...).Exec(ctx); err != nil {
		o.client.Order.FindUnique(db.Order.ID.Equals(orderDB.ID)).Delete().Exec(ctx)
		return nil, err
	}
	return o.Show(ctx, &ShowPayload{OrderID: orderDB.ID})
}

func (o *OrderService) Cancel(ctx context.Context, payload *CancelPayload) (bool, error) {
	o.Info("Cancel Order got called with", zap.Any("payload", payload))
	_, err := o.client.Order.
		FindUnique(db.Order.ID.Equals(payload.OrderID)).
		Update(
			db.Order.CancelledAt.Set(time.Now()),
		).Exec(ctx)
	o.Info("Cancel Order response", zap.Error(err))
	return err != nil, err
}

func (o *OrderService) List(ctx context.Context, payload *ListPayload) (*OrderList, error) {
	o.Info("List Order got called with", zap.Any("payload", payload))

	filters := []db.OrderWhereParam{}
	if user, ok := ctx.Value(auth.UserCtxKey).(*db.UserModel); ok {
		filters = append(filters, db.Order.UserID.Equals(user.ID))
	}

	ordersDB, err := o.client.Order.FindMany(filters...).
		With(Connection...).
		Take(payload.PageSize).
		Skip(payload.After).
		Exec(ctx)
	if err != nil {
		o.Error("Error List Order findMany", zap.Error(err))
		return nil, err
	}
	orders := []*Order{}
	for _, orderDB := range ordersDB {
		orders = append(orders, MapOrderModelToOrder(&orderDB))
	}

	var resp []struct {
		Count db.BigInt `json:"count"`
	}
	err = o.client.Prisma.QueryRaw("SELECT count(*) FROM project.orders").Exec(ctx, &resp)

	if err != nil {
		o.Error("Error count order", zap.Error(err))
		return nil, err
	}
	count := 0
	if len(resp) == 0 {
		count = int(resp[0].Count)
	}
	nextPageCursor := utils.MinInt(count, payload.After+payload.PageSize)
	response := &OrderList{
		Data: orders,
		PageInfo: &PageInfo{
			StartCursor:   payload.After,
			EndCursor:     nextPageCursor,
			TotalResource: count,
			HasMore:       nextPageCursor < count,
		},
	}
	return response, nil

}

func MapOrderModelToOrder(model *db.OrderModel) *Order {
	output := Order{
		ID:         model.ID,
		Email:      model.UserEmail,
		TotalPrice: model.TotalPrice.String(),
		CreatedAt:  model.CreatedAt.String(),
	}
	if value := model.User(); value != nil {
		user := *MapUserDBToOutput(*value)
		output.User = &User{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Role:      UserRole(user.Role),
			UpdatedAt: user.UpdatedAt,
		}
	}
	if value := model.Address(); value != nil {
		addressOut := &Address{
			ID:           value.ID,
			AddressLine1: value.AddressLine1,
			City:         value.City,
			Country:      value.Country,
			State:        value.Province,
			CreatedAt:    value.CreatedAt.String(),
		}

		if data, ok := value.UpdatedAt(); ok {
			addressOut.UpdatedAt = utils.StringRef(data.String())
		}

		if data, ok := value.AddressLine2(); ok {
			addressOut.AddressLine2 = utils.StringRef(data)
		}

		if data, ok := value.ZipCode(); ok {
			addressOut.ZipCode = *utils.StringRef(data)
		}
	}

	lineItems := []*OrderLineItem{}
	lineItemsDB := model.LineItems()
	if len(lineItemsDB) == 0 {
		output.LineItems = lineItems
		return &output
	}

	for _, lineItemDB := range lineItemsDB {
		lineItem := OrderLineItem{
			ID:       lineItemDB.ID,
			Price:    lineItemDB.Price.String(),
			Quantity: lineItemDB.Quantity,
		}
		if value, ok := lineItemDB.ProductID(); ok {
			valueInt := int(value)
			lineItem.ProductID = &valueInt
		}
		lineItems = append(lineItems, &lineItem)
	}
	output.LineItems = lineItems
	return &output
}

func (o *OrderService) Show(ctx context.Context, payload *ShowPayload) (*Order, error) {
	o.Info("Show Order got called with", zap.Any("payload", payload))
	order, err := o.client.Order.FindUnique(
		db.Order.ID.Equals(payload.OrderID),
	).With(Connection...).
		Exec(ctx)
	if err != nil {
		if db.IsErrNotFound(err) {
			return nil, MakeNotFound(err)
		}
		return nil, err
	}
	return MapOrderModelToOrder(order), nil
}

func MountSvcOrder(mux http.Muxer, svc Service) {
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

func NewOrderService(
	client *db.PrismaClient,
	validator *auth.JWTValidator,
) Service {
	logger := zap.L()
	return &OrderService{
		client:       client,
		Logger:       logger,
		JWTValidator: validator,
	}
}

func MountOrderSVC(mux http.Muxer, svc Service) {
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
