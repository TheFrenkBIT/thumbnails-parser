package parser

import (
	"context"
	parser "github.com/evrone/go-clean-template/internal/controller/proto"
	parserService "github.com/evrone/go-clean-template/internal/usecase/services/parser"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	parser.UnimplementedPreviewServer
	service parserService.Interface
}

func Register(gRPC *grpc.Server, service parserService.Interface) {
	parser.RegisterPreviewServer(gRPC, &Handler{service: service})
}

func (h *Handler) Parse(ctx context.Context, req *parser.ParseRequest) (*parser.ParseResponse, error) {
	images, err := h.service.Parse(ctx, req.Urls)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &parser.ParseResponse{
		Images: images,
	}, nil
}
