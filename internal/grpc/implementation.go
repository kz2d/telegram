package grpc

import (
	"context"
	"github.com/kz2d/telegram-bot/internal/grpc/api"
	"github.com/kz2d/telegram-bot/pkg/database"
	"github.com/kz2d/telegram-bot/pkg/telegram"
)

type GreeterServerImpl struct {
	api.UnimplementedSomeServiceServer
}

func (g *GreeterServerImpl) SendMessage(ctx context.Context, request *api.SendMessageItem) (*api.Empty, error) {
	r, _ := database.Db.Query("select chat_id,show from db.public.users")

	var s int64
	var b bool

	for r.Next() {
		r.Scan(&s, &b)
		if b {
			telegram.SendMessage(s, request.Text)
		}
	}

	return &api.Empty{}, nil
}

func (g *GreeterServerImpl) SendPhoto(ctx context.Context, request *api.SendPhotoItem) (*api.Empty, error) {
	r, _ := database.Db.Query("select chat_id,show from db.public.users")

	var s int64
	var b bool

	for r.Next() {
		r.Scan(&s, &b)
		if b {
			telegram.SendPhoto(s, request.Message, request.PhotoUrl)
		}
	}

	return &api.Empty{}, nil
}
