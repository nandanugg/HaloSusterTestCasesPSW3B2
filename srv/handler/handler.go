package handler

import (
	"github.com/nandanugg/HaloSusterTestCasesPSW3B2/entity/pb"
	"github.com/nandanugg/HaloSusterTestCasesPSW3B2/service"
)

type Handler struct {
	srv *service.NipService
	pb.UnimplementedNIPServiceServer
}

func NewRequestHandler(srv *service.NipService) *Handler {
	return &Handler{srv: srv}
}
