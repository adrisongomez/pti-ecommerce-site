package services

import (
	"context"
	"fmt"

	"github.com/adrisongomez/pti-ecommerce-site/backends/databases/db"
	"github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/http/order/server"
	. "github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/order"
	"github.com/adrisongomez/pti-ecommerce-site/backends/internal/utils/auth"
	"go.uber.org/zap"
	"goa.design/goa/v3/http"
)

type OrderService struct {
	client *db.PrismaClient

	*auth.JWTValidator
	*zap.Logger
}

var Connection = []db.OrderRelationWith{
	db.Order.Address.Fetch(),
	db.Order.LineItems.Fetch().With(db.OrderLineItem.Order.Fetch()),
	db.Order.User.Fetch(),
}

func (o *OrderService) Create(ctx context.Context, payload *CreatePayload) (*Order, error) {
	o.Info("Create Order got called with", zap.Any("payload", payload))
	return nil, nil
}

func (o *OrderService) Cancel(ctx context.Context, payload *CancelPayload) (bool, error) {
	o.Info("Cancel Order got called with", zap.Any("payload", payload))
	return false, nil
}

func (o *OrderService) List(ctx context.Context, payload *ListPayload) (*OrderList, error) {
	o.Info("List Order got called with", zap.Any("payload", payload))
	return nil, nil

}

func MapOrderModelToOrder(model *db.OrderModel) *Order {
	output := Order{
		ID:         model.ID,
		Email:      model.UserEmail,
		TotalPrice: model.TotalPrice,
		CreatedAt:  model.CreatedAt.String(),
	}
	lineItems := []*OrderLineItem{}
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
	model.Address()
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
			lineItem.ProductID = value
		}
		lineItems = append(lineItems, &lineItem)
	}
	return &output
}

func (o *OrderService) Show(ctx context.Context, payload *ShowPayload) (*Order, error) {
	o.Info("Show Order got called with", zap.Any("payload", payload))
	order, err := o.client.Order.FindUnique(
		db.Order.ID.Equals(payload.OrderID),
	).With().
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
