package handler

import (
	"context"

	"github.com/nandanugg/HaloSusterTestCasesPSW3B2/entity"
	"github.com/nandanugg/HaloSusterTestCasesPSW3B2/entity/pb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func (s *Handler) GetUsedIt(ctx context.Context, req *emptypb.Empty) (*pb.PostUsedAcc, error) {
	usr := s.srv.GetItUsedAccount()
	return &pb.PostUsedAcc{
		Nip:      usr.Nip,
		Password: usr.Password,
	}, nil
}

func (s *Handler) PostUsedIT(ctx context.Context, req *pb.PostUsedAcc) (*emptypb.Empty, error) {
	s.srv.AddItUsedAccount(entity.UsedUser{
		Nip:      req.Nip,
		Password: req.Password,
	})
	return nil, nil
}

func (s *Handler) GetItNip(ctx context.Context, req *emptypb.Empty) (*pb.GetNipResponse, error) {
	return &pb.GetNipResponse{
		Nip: s.srv.GetItNip(),
	}, nil
}
