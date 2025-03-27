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

type UserService struct {
	client *db.PrismaClient
	logger *zap.Logger
}

func MapUserDBToOutput(model db.UserModel) *User {
	user := User{
		ID:        model.ID,
		FirstName: model.FirstName,
		LastName:  "",
		Email:     model.Email,
		Role:      UserRole(model.Role),
		CreatedAt: model.CreatedAt.String(),
	}
	if value, ok := model.LastName(); ok {
		user.LastName = value
	}

	if value, ok := model.UpdatedAt(); ok {
		user.UpdatedAt = utils.StringRef(value.String())
	}
	return &user
}

func (u *UserService) List(ctx context.Context, payload *ListPayload) (*UserList, error) {
	u.logger.Info("List got called With", zap.Any("payload", payload))
	usersDB, err := u.client.User.FindMany(
		db.User.Not(db.User.DeletedAt.IsNull()),
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
		"SELECT count(*) FROM project.users WHERE deleted_at IS NOT NULL",
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
			TotalResource: payload.After,
			EndCursor:     nextCursor,
			HasMore:       count != nextCursor,
		},
	}
	return &response, nil
}

func NewUserService(client *db.PrismaClient) *UserService {
	logger := zap.L()
	return &UserService{client, logger}
}

func MountUserServiceSVC(mux http.Muxer, svc *UserService) {
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
