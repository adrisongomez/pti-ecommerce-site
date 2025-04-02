package services

import (
	"context"
	"fmt"

	"github.com/adrisongomez/pti-ecommerce-site/backends/databases/db"
	. "github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/chat"
	"github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/http/chat/server"
	"github.com/adrisongomez/pti-ecommerce-site/backends/internal/utils"
	"github.com/adrisongomez/pti-ecommerce-site/backends/internal/utils/auth"
	"github.com/adrisongomez/pti-ecommerce-site/backends/pkg/assistant"
	"go.uber.org/zap"
	"goa.design/goa/v3/http"
)

type ChatService struct {
	client          *db.PrismaClient
	assistantClient *assistant.AssitantService
	*auth.JWTValidator
}

func NewChatService(client *db.PrismaClient, validator *auth.JWTValidator, assistantClient *assistant.AssitantService) Service {
	return &ChatService{client, assistantClient, validator}
}

func (c *ChatService) CreateChatSession(ctx context.Context, payload *CreateChatSessionPayload) (int, error) {
	user, ok := ctx.Value(auth.UserCtxKey).(*db.UserModel)
	if !ok {
		return 0, fmt.Errorf("Not user in the token")
	}

	chatSession, err := c.client.ChatSession.CreateOne(db.ChatSession.User.Link(db.User.ID.Equals(user.ID))).Exec(ctx)
	if err != nil {
		return 0, nil
	}
	return chatSession.ID, nil
}

func MapChatSessionDBtoAPIOutput(model *db.ChatSessionModel) *ChatSession {
	output := ChatSession{
		ID:        model.ID,
		CreatedAt: model.CreatedAt.String(),
	}
	messages := []*ChatMessage{}
	for _, value := range model.Messages() {
		message := ChatMessage{
			ID:        value.ID,
			Message:   value.Message,
			Source:    ChatSource(value.Source),
			CreatedAt: value.CreatedAt.String(),
		}
		messages = append(messages, &message)
	}
	output.Messages = messages
	if value, ok := model.UpdatedAt(); ok {
		output.UpdatedAt = utils.StringRef(value.String())
	}
	return &output
}

func (c *ChatService) GetSessionByID(ctx context.Context, payload *GetSessionByIDPayload) (*ChatSession, error) {
	user, ok := ctx.Value(auth.UserCtxKey).(*db.UserModel)
	if !ok {
		return nil, fmt.Errorf("Not user in the token")
	}
	chatSession, err := c.client.ChatSession.FindFirst(
		db.ChatSession.ID.Equals(*payload.SessionID),
		db.ChatSession.UserID.Equals(user.ID),
	).With(
		db.ChatSession.Messages.Fetch(),
	).Exec(ctx)
	if err != nil {
		return nil, nil
	}
	return MapChatSessionDBtoAPIOutput(chatSession), nil
}

func (c *ChatService) SubmitMessageToSession(ctx context.Context, payload *SubmitMessageToSessionPayload) (string, error) {
	_, ok := ctx.Value(auth.UserCtxKey).(*db.UserModel)
	if !ok {
		return "", fmt.Errorf("Not user in the token")
	}
	_, err := c.client.ChatMessage.CreateOne(
		db.ChatMessage.Source.Set(db.ChatSourceUser),
		db.ChatMessage.Message.Set(*payload.Message),
		db.ChatMessage.Chat.Link(db.ChatSession.ID.Equals(*payload.SessionID)),
	).Exec(ctx)
	if err != nil {
		zap.L().Error("Error saving user message", zap.Error(err))
		return "", err
	}
	openApiMessage := ""

	_, err = c.client.ChatMessage.CreateOne(
		db.ChatMessage.Source.Set(db.ChatSourceAssistant),
		db.ChatMessage.Message.Set(openApiMessage),
		db.ChatMessage.Chat.Link(db.ChatSession.ID.Equals(*payload.SessionID)),
	).Exec(ctx)

	if err != nil {
		zap.L().Error("Error saving assitant message", zap.String("message", openApiMessage), zap.Error(err))
		return "", err
	}
	return openApiMessage, nil
}

func MountChatSVC(mux http.Muxer, svc Service) {
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
