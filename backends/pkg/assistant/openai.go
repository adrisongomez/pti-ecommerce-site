package assistant

import (
	"context"
	"encoding/json"
	"os"

	"github.com/adrisongomez/pti-ecommerce-site/backends/databases/db"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type AssitantService struct {
	dbClient *db.PrismaClient
	client   *openai.Client
}

func NewAssitantService(dbClient *db.PrismaClient) *AssitantService {
	client := openai.NewClient(
		option.WithAPIKey(os.Getenv("OPEN_AI_API_KEY")), // defaults to os.LookupEnv("OPENAI_API_KEY")
	)
	return &AssitantService{client: &client, dbClient: dbClient}
}

const (
	getProductFunctionName = "get_products"
)

var GetProductFunction = openai.FunctionDefinitionParam{
	Name:        getProductFunctionName,
	Description: openai.String("Get a list of products that combine with orange"),
	Parameters:  openai.FunctionParameters{},
}

func (a *AssitantService) GetOpenAPIMessage(userMessage string, ctx context.Context) (string, error) {
	params := openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(userMessage),
		},
		Tools: []openai.ChatCompletionToolParam{
			{
				Function: GetProductFunction,
			},
		},
		Model: openai.ChatModelGPT4o,
	}
	completions, err := a.client.Chat.Completions.New(ctx, params)
	if err != nil {
		return "", err
	}
	toolCalls := completions.Choices[0].Message.ToolCalls

	if len(toolCalls) == 0 {
		return completions.Choices[0].Message.Content, nil
	}
	params.Messages = append(params.Messages, completions.Choices[0].Message.ToParam())
	for _, toolCall := range toolCalls {
		if toolCall.Function.Name == getProductFunctionName {
			dbProduct, err := a.dbClient.Product.FindMany().Take(100).Exec(ctx)
			if err != nil {
				return "", err
			}
			jsonProduct, err := json.Marshal(dbProduct)
			if err != nil {
				return "", err
			}
			params.Messages = append(params.Messages, openai.ToolMessage(string(jsonProduct), toolCall.ID))
		}
	}

	completions, err = a.client.Chat.Completions.New(ctx, params)
	if err != nil {
		return "", err
	}
	return completions.Choices[0].Message.Content, nil
}
