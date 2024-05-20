package handler

import (
	"context"

	"github.com/nandanugg/HaloSusterTestCasesPSW3B2/entity"
	"github.com/nandanugg/HaloSusterTestCasesPSW3B2/entity/pb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func (s *Handler) GetUsedNurse(ctx context.Context, req *emptypb.Empty) (*pb.PostUsedAcc, error) {
	usr := s.srv.GetNurseUsedAccount()
	return &pb.PostUsedAcc{
		Nip:      usr.Nip,
		Password: usr.Password,
	}, nil
}

func (s *Handler) PostUsedNurse(ctx context.Context, req *pb.PostUsedAcc) (*emptypb.Empty, error) {
	s.srv.AddNurseUsedAccount(entity.UsedUser{
		Nip:      req.Nip,
		Password: req.Password,
	})
	return nil, nil
}

func (s *Handler) GetNurseNip(ctx context.Context, req *emptypb.Empty) (*pb.GetNipResponse, error) {
	return &pb.GetNipResponse{
		Nip: s.srv.GetNurseNip(),
	}, nil
}
