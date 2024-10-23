package auth

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/algol-84/chat-server/internal/converter"
	desc "github.com/algol-84/chat-server/pkg/chat_v1"
)

// Create обрабатывает GRPC запросы на создание нового юзера
func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Println(req.Chat)
	userID, err := i.chatService.Create(ctx, converter.FromChatToService(req.Chat))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "user creation in DB returned with error: %s", err)
	}

	return &desc.CreateResponse{
		Id: userID,
	}, nil
}