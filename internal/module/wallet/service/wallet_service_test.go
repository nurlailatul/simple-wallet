package service

import (
	"context"
	"reflect"
	"testing"

	"simple-wallet/internal/module/wallet/domain"
	"simple-wallet/internal/module/wallet/mocks"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var (
	dummyEntity = domain.WalletEntity{
		ID:        1,
		UserID:    1,
		Balance:   float64(10000),
		CreatedAt: 1282312835,
		UpdatedAt: 1282312835,
	}
)

type WalletServiceSuite struct {
	suite.Suite
	repo *mocks.WalletRepositoryInterface
}

func TestWalletService(t *testing.T) {
	suite.Run(t, new(WalletServiceSuite))
}

func (s *WalletServiceSuite) SetupSuite(t *testing.T) *WalletServiceSuite {
	s.Suite.SetT(t)
	s.repo = new(mocks.WalletRepositoryInterface)

	return s
}

func (s *WalletServiceSuite) TestWalletService_GetByCompanyId() {
	type fields struct {
		repo *mocks.WalletRepositoryInterface
	}
	type args struct {
		ctx    context.Context
		userID int64
	}
	tests := []struct {
		name   string
		fields func(t *testing.T) fields
		args   args
		want   *domain.WalletEntity
	}{
		{
			name: "success",
			fields: func(t *testing.T) fields {
				s := new(WalletServiceSuite).SetupSuite(t)
				s.repo.On("GetByUserID", mock.Anything, mock.Anything).Return(&dummyEntity, nil)

				return fields{repo: s.repo}
			},
			args: args{
				ctx:    context.Background(),
				userID: 1,
			},
			want: &dummyEntity,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			mock := tt.fields(t)
			service := NewWalletService(mock.repo)

			got := service.GetByUserID(tt.args.ctx, tt.args.userID)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WalletService.GetByCompanyId() = %v, want %v", got, tt.want)
			}
		})
	}
}
