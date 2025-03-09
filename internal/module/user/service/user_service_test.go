package service

import (
	"context"
	"reflect"
	"testing"

	"simple-wallet/internal/module/user/domain"
	"simple-wallet/internal/module/user/mocks"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var (
	dummyEntity = domain.UserEntity{
		ID:        1,
		Phone:     "+6281234567890",
		Name:      "Nur Lailatul",
		Email:     "nurlailatul@gmail.com",
		Status:    1,
		CreatedAt: 1282312835,
		UpdatedAt: 1282312835,
	}
)

type UserServiceSuite struct {
	suite.Suite
	repo *mocks.UserRepositoryInterface
}

func TestUserService(t *testing.T) {
	suite.Run(t, new(UserServiceSuite))
}

func (s *UserServiceSuite) SetupSuite(t *testing.T) *UserServiceSuite {
	s.Suite.SetT(t)
	s.repo = new(mocks.UserRepositoryInterface)

	return s
}

func (s *UserServiceSuite) TestUserService_GetByID() {
	type fields struct {
		repo *mocks.UserRepositoryInterface
	}
	type args struct {
		ctx    context.Context
		userID int64
	}
	tests := []struct {
		name   string
		fields func(t *testing.T) fields
		args   args
		want   *domain.UserEntity
	}{
		{
			name: "success",
			fields: func(t *testing.T) fields {
				s := new(UserServiceSuite).SetupSuite(t)
				s.repo.On("GetByID", mock.Anything, mock.Anything).Return(&dummyEntity, nil)

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
			service := NewUserService(mock.repo)

			got := service.GetByID(tt.args.ctx, tt.args.userID)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.GetByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
