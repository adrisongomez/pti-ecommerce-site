package services

import (
	"context"
	"fmt"

	"github.com/adrisongomez/pti-ecommerce-site/backends/databases/db"
	svcuserhttp "github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/http/svcuser/server"
	. "github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/svcuser"
	"github.com/adrisongomez/pti-ecommerce-site/backends/internal/utils"
	"go.uber.org/zap"
	"goa.design/goa/v3/http"
)

type UserController struct {
	client *db.PrismaClient
	logger *zap.Logger
}

func MapUserDBToOutput(model db.UserModel) *User {
	user := User{
		ID:        model.ID,
		FirstName: model.FirstName,
		LastName:  nil,
		Email:     model.Email,
		Role:      UserRole(model.Role),
		CreatedAt: model.CreatedAt.String(),
	}
	if value, ok := model.LastName(); ok {
		user.LastName = &value
	}

	if value, ok := model.UpdatedAt(); ok {
		user.UpdatedAt = utils.StringRef(value.String())
	}
	return &user
}

func (u *UserController) List(ctx context.Context, payload *ListPayload) (*UserList, error) {
	u.logger.Info("List got called With", zap.Any("payload", payload))
	usersDB, err := u.client.User.FindMany(
		db.User.DeletedAt.IsNull(),
	).
		Take(payload.PageSize).
		Skip(payload.After).Exec(ctx)
	users := []*User{}
	if err != nil {
		u.logger.Error("Error on requesting users", zap.Error(err))
		return nil, err
	}

	u.logger.Debug("DB response", zap.Any("usersDB", usersDB))
	for _, userDB := range usersDB {
		users = append(users, MapUserDBToOutput(userDB))
	}

	var rows []struct {
		Count db.BigInt `json:"count"`
	}
	err = u.client.Prisma.QueryRaw(
		"SELECT count(*) FROM project.users WHERE deleted_at IS NULL",
	).Exec(ctx, &rows)

	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, fmt.Errorf("Not result from count")
	}
	count := int(rows[0].Count)

	nextCursor := utils.MinInt(count, payload.PageSize+payload.After)

	response := UserList{
		Data: users,
		PageInfo: &PageInfo{
			StartCursor:   payload.After,
			TotalResource: count,
			EndCursor:     nextCursor,
			HasMore:       count != nextCursor,
		},
	}
	return &response, nil
}

func (u *UserController) Create(ctx context.Context, payload *UserCreateInput) (*User, error) {
	u.logger.Info("User#create got called", zap.Any("payload", payload))
	changes := []db.UserSetParam{
		db.User.Role.Set(db.UserRole(payload.Role)),
	}
	if payload.LastName != nil {
		changes = append(changes, db.User.LastName.Set(*payload.LastName))
	}
	userModel, err := u.client.User.CreateOne(
		db.User.FirstName.Set(payload.FirstName),
		db.User.Email.Set(payload.Email),
		changes...,
	).Exec(ctx)

	if err != nil {
		u.logger.Error("User#create got error", zap.Error(err))
		return nil, err
	}

	return MapUserDBToOutput(*userModel), nil
}

func (u *UserController) Show(ctx context.Context, payload *ShowPayload) (*User, error) {
	u.logger.Info("User#Show got called with", zap.Any("payload", payload))
	userDB, err := u.client.User.FindUnique(db.User.ID.Equals(payload.UserID)).Exec(ctx)
	if err != nil {
		u.logger.Error("User#Show got error", zap.Error(err))
		return nil, err
	}
	return MapUserDBToOutput(*userDB), nil
}

func (u *UserController) Update(ctx context.Context, input *UpdatePayload) (*User, error) {
	payload := input.Payload
	u.logger.Info("User#create got called", zap.Any("payload", payload))
	changes := []db.UserSetParam{
		db.User.FirstName.Set(payload.FirstName),
		db.User.Email.Set(payload.Email),
		db.User.Role.Set(db.UserRole(payload.Role)),
	}
	if payload.LastName != nil {
		changes = append(changes, db.User.LastName.Set(*payload.LastName))
	}
	userModel, err := u.client.User.FindUnique(
		db.User.ID.Equals(input.UserID),
	).Update(
		changes...,
	).Exec(ctx)

	if err != nil {
		u.logger.Error("User#create got error", zap.Error(err))
		return nil, err
	}
	return MapUserDBToOutput(*userModel), nil
}

func (u *UserController) Delete(ctx context.Context, input *DeletePayload) (bool, error) {
	u.logger.Info("User#create got called", zap.Any("payload", input))

	_, err := u.client.User.FindUnique(db.User.ID.Equals(input.UserID)).Delete().Exec(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewUserService(client *db.PrismaClient) *UserController {
	logger := zap.L()
	return &UserController{client, logger}
}

func MountUserServiceSVC(mux http.Muxer, svc *UserController) {
	endpoints := NewEndpoints(svc)
	req := http.RequestDecoder
	res := http.ResponseEncoder

	handler := svcuserhttp.New(endpoints, mux, req, res, nil, nil)
	svcuserhttp.Mount(mux, handler)

	go func() {
		for _, mount := range handler.Mounts {
			zap.L().Info(fmt.Sprintf("%q mounted on %s %s\n", mount.Method, mount.Verb, mount.Pattern))
		}
	}()
}
