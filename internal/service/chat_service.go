package service

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/CRobinDev/Gemastik/entity"
	"github.com/CRobinDev/Gemastik/internal/repository"
	"github.com/CRobinDev/Gemastik/model"
	"github.com/CRobinDev/Gemastik/pkg/errors"
	"github.com/CRobinDev/Gemastik/pkg/gocron"
	"github.com/sashabaranov/go-openai"
)

type IChatService interface {
	GenerateResponse(req model.ChatRequest) (model.ServiceResponse, error)
	GenerateTextResponse(req model.ChatRequest) (model.ServiceResponse, error)
	GenerateImageResponse(req model.ChatRequest) (model.ServiceResponse, error)
}

type ChatService struct {
	client         *openai.Client
	UserRepository repository.IUserRepository
	ChatRepository repository.IChatRepository
	chatHistory    []string
}

func NewChatService(userRepository repository.IUserRepository, chatRepository repository.IChatRepository) IChatService {
	apiKey := os.Getenv("OPENAI_API")
	client := openai.NewClient(apiKey)
	return &ChatService{
		client:         client,
		UserRepository: userRepository,
		ChatRepository: chatRepository,
		chatHistory:    make([]string, 0),
	}
}
func (cs *ChatService) GenerateResponse(req model.ChatRequest) (model.ServiceResponse, error) {
	if strings.Contains(strings.ToLower(req.Chat), "gambar") {
		return cs.GenerateImageResponse(req)
	}
	return cs.GenerateTextResponse(req)
}

func (cs *ChatService) GenerateTextResponse(req model.ChatRequest) (model.ServiceResponse, error) {
	chatHistory, err := cs.ChatRepository.GetHistory(req.UserID)
	if err != nil {
		return model.ServiceResponse{
			Code:    http.StatusInternalServerError,
			Error:   true,
			Message: errors.ErrInternalServer.Error(),
			Data:    err,
		}, err
	}

	request := openai.ChatCompletionRequest{
		Model: openai.GPT4Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "respond as an expert from Indonesia Corruption Watch (ICW)!",
			},
		},
		MaxTokens: 512,
	}

	for _, chat := range chatHistory {
		request.Messages = append(request.Messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: chat,
		})
	}
	log.Println("chat history", chatHistory)

	request.Messages = append(request.Messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: req.Chat,
	})

	resp, err := cs.client.CreateChatCompletion(context.Background(), request)
	if err != nil {
		return model.ServiceResponse{
			Code:    http.StatusInternalServerError,
			Error:   true,
			Message: errors.ErrInternalServer.Error(),
			Data:    err,
		}, err
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")
	createdAt, _ := time.Parse("2006-01-02 15:04:05", time.Now().In(loc).Format("2006-01-02 15:04:05"))

	chat := entity.Chat{
		UserID:    req.UserID,
		Input:     req.Chat,
		Output:    resp.Choices[0].Message.Content,
		CreatedAt: createdAt,
	}
	
	if err := cs.ChatRepository.InsertChat(chat); err != nil {
		return model.ServiceResponse{
			Code:    http.StatusInternalServerError,
			Error:   true,
			Message: errors.ErrInternalServer.Error(),
			Data:    err,
		}, err
	}

	gocron.StartChatCleanupScheduler(cs.ChatRepository)
	
	return model.ServiceResponse{
		Code:    http.StatusOK,
		Error:   false,
		Message: "success",
		Data:    resp.Choices[0].Message.Content,
	}, nil
}

func (cs *ChatService) GenerateImageResponse(req model.ChatRequest) (model.ServiceResponse, error) {
	respUrl, err := cs.client.CreateImage(
		context.Background(),
		openai.ImageRequest{
			Prompt:         req.Chat,
			Size:           openai.CreateImageSize256x256,
			ResponseFormat: openai.CreateImageResponseFormatURL,
			N:              1,
		},
	)
	if err != nil {
		return model.ServiceResponse{}, err
	}

	imageURL := respUrl.Data[0].URL
	return model.ServiceResponse{
		Code:    http.StatusOK,
		Error:   false,
		Message: "Image generated successfully",
		Data:    imageURL,
	}, nil
}
