package service

import (
	"context"

	"github.com/K-Kizuku/Somniosus_back/apps/account/app/domain/model/account"
	"github.com/K-Kizuku/Somniosus_back/apps/account/app/usecase"
	"github.com/K-Kizuku/Somniosus_back/proto.pb/twitter/account/resources"
	"github.com/K-Kizuku/Somniosus_back/proto.pb/twitter/account/rpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Service) CreateTempAccount(ctx context.Context, in *rpc.CreateTempAccountRequest) (*rpc.CreateTempAccountResponse, error) {
	jwt, err := s.au.CreateTempAccount(ctx, &usecase.CreateTempAccountInput{
		Name:     in.Name,
		TelNum:   in.TelNum,
		BirthDay: in.BirthDay.AsTime(),
	})
	if err != nil {
		return nil, err
	}
	return &rpc.CreateTempAccountResponse{
		TempToken: jwt,
	}, nil
}

func (s *Service) CreateAccount(ctx context.Context, in *rpc.CreateAccountRequest) (*rpc.CreateAccountResponse, error) {
	jwt, err := s.au.CreateAccount(ctx, in.TempToken, in.DisplayId, in.Password)
	if err != nil {
		return nil, err
	}
	return &rpc.CreateAccountResponse{
		Jwt: jwt,
	}, nil
}

func (s *Service) VerifyToken(ctx context.Context, in *rpc.VerifyTokenRequest) (*rpc.VerifyTokenResponse, error) {
	id, err := s.au.VerifyToken(ctx, in.Jwt)
	if err != nil {
		return nil, err
	}
	return &rpc.VerifyTokenResponse{
		AccountId: id,
	}, nil
}

func (s *Service) GenerateToken(ctx context.Context, in *rpc.GenerateTokenRequest) (*rpc.GenerateTokenResponse, error) {
	jwt, err := s.au.GenerateToken(ctx, in.AccountId, in.Password, account.AccountStatus(in.Status))
	if err != nil {
		return nil, err
	}
	return &rpc.GenerateTokenResponse{
		Token: jwt,
	}, nil
}

func (s *Service) GetAccount(ctx context.Context, in *rpc.GetAccountRequest) (*rpc.GetAccountResponse, error) {
	ac, err := s.au.GetAccount(ctx, in.AccountId)
	if err != nil {
		return nil, err
	}
	return &rpc.GetAccountResponse{
		Account: &resources.Account{
			AccountId:   ac.ID(),
			DisplayId:   ac.DisplayID(),
			Name:        ac.Name(),
			Description: ac.Description(),
			WebsiteUrl:  ac.WebsiteURL(),
			ImageUrl:    ac.ImageURL(),
			TelNum:      ac.TelNum(),
			BirthDay:    timestamppb.New(ac.BirthDay()),
			Status:      resources.AccountStatus(ac.Status()),
		},
	}, nil
}

func (s *Service) UpdateAccount(ctx context.Context, in *rpc.UpdateAccountRequest) (*emptypb.Empty, error) {
	if err := s.au.UpdateAccount(ctx, account.Assign(account.General, in.AccountId, *in.DisplayId, *in.Name, *in.Description, *in.ImageUrl, *in.TelNum, in.BirthDay.AsTime())); err != nil {
		return nil, err
	}
	return nil, nil
}
