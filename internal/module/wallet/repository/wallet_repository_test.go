package repository

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"regexp"
	"testing"

	"simple-wallet/internal/module/wallet/domain"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	tableName    = "wallets"
	walletColumn = []string{"id", "user_id", "balance", "created_at", "updated_at"}
)

type WalletRepositorySuite struct {
	suite.Suite
	db   *gorm.DB
	mock sqlmock.Sqlmock
}

func TestWalletRepositorySuite(t *testing.T) {
	suite.Run(t, new(WalletRepositorySuite))
}

func (s *WalletRepositorySuite) SetupSuite(t *testing.T) *WalletRepositorySuite {
	var (
		mockDb *sql.DB
		err    error
	)

	s.Suite.SetT(t)

	mockDb, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	dialector := mysql.New(mysql.Config{
		Conn:                      mockDb,
		SkipInitializeWithVersion: true,
	})
	s.db, err = gorm.Open(dialector, &gorm.Config{})
	assert.NoError(s.T(), err)

	return s
}

func (s *WalletRepositorySuite) TestWalletRepository_GetByUserID() {
	type fields struct {
		gorm *gorm.DB
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
				s := new(WalletRepositorySuite).SetupSuite(t)

				s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `wallets` WHERE user_id = ? ORDER BY `wallets`.`id` LIMIT ?")).
					WithArgs(1, 1).
					WillReturnRows(sqlmock.NewRows(walletColumn).
						AddRow(1, 1, float64(10000), 1282312835, 1282312835))

				return fields{gorm: s.db}
			},
			args: args{
				ctx:    context.Background(),
				userID: 1,
			},
			want: &domain.WalletEntity{
				ID:        1,
				UserID:    1,
				Balance:   float64(10000),
				CreatedAt: 1282312835,
				UpdatedAt: 1282312835,
			},
		},
		{
			name: "failed",
			fields: func(t *testing.T) fields {
				s := new(WalletRepositorySuite).SetupSuite(t)

				s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `wallets` WHERE user_id = ? ORDER BY `wallets`.`id` LIMIT ?")).
					WithArgs(1, 1).
					WillReturnError(sql.ErrNoRows)

				return fields{gorm: s.db}
			},
			args: args{
				ctx:    context.Background(),
				userID: 1,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			conn := tt.fields(t)

			repo := NewWalletRepository(conn.gorm)
			got := repo.GetByUserID(tt.args.ctx, tt.args.userID)
			if got == nil && tt.want != nil {
				t.Errorf("WalletRepository.GetByUserID() got = %v, want %v", got, tt.want)
				return
			}
			if got != nil && !reflect.DeepEqual(*got, *tt.want) {
				t.Errorf("WalletRepository.GetByUserID() got = %v, want %v", got, tt.want)
			}

			sqlDB, _ := conn.gorm.DB()
			sqlDB.Close()
		})
	}
}

func (s *WalletRepositorySuite) TestWalletRepository_GetByUserIDForLocking() {
	type fields struct {
		gorm *gorm.DB
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
				s := new(WalletRepositorySuite).SetupSuite(t)

				s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `wallets` WHERE id = ? ORDER BY `wallets`.`id` LIMIT ? FOR UPDATE")).
					WithArgs(1, 1).
					WillReturnRows(sqlmock.NewRows(walletColumn).
						AddRow(1, 1, float64(10000), 1282312835, 1282312835))

				return fields{gorm: s.db}
			},
			args: args{
				ctx:    context.Background(),
				userID: 1,
			},
			want: &domain.WalletEntity{
				ID:        1,
				UserID:    1,
				Balance:   float64(10000),
				CreatedAt: 1282312835,
				UpdatedAt: 1282312835,
			},
		},
		{
			name: "failed",
			fields: func(t *testing.T) fields {
				s := new(WalletRepositorySuite).SetupSuite(t)

				s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `wallets` WHERE id = ? ORDER BY `wallets`.`id` LIMIT ? FOR UPDATE")).
					WithArgs(1, 1).
					WillReturnError(sql.ErrNoRows)

				return fields{gorm: s.db}
			},
			args: args{
				ctx:    context.Background(),
				userID: 1,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			conn := tt.fields(t)

			repo := NewWalletRepository(conn.gorm)
			got := repo.GetByUserIDForLocking(tt.args.ctx, tt.args.userID)
			if got == nil && tt.want != nil {
				t.Errorf("WalletRepository.GetByUserIDForLocking() got = %v, want %v", got, tt.want)
				return
			}
			if got != nil && !reflect.DeepEqual(*got, *tt.want) {
				t.Errorf("WalletRepository.GetByUserIDForLocking() got = %v, want %v", got, tt.want)
			}

			sqlDB, _ := conn.gorm.DB()
			sqlDB.Close()
		})
	}
}

func (s *WalletRepositorySuite) TestWalletRepository_Update() {
	type fields struct {
		gorm *gorm.DB
	}
	type args struct {
		ctx    context.Context
		entity *domain.WalletEntity
	}
	tests := []struct {
		name   string
		fields func(t *testing.T) fields
		args   args
		want   error
	}{
		{
			name: "success",
			fields: func(t *testing.T) fields {
				s := new(WalletRepositorySuite).SetupSuite(t)

				s.mock.ExpectBegin()
				s.mock.ExpectExec(fmt.Sprintf("UPDATE `%s` SET", tableName)).
					WillReturnResult(sqlmock.NewResult(1, 1))
				s.mock.ExpectCommit()

				return fields{gorm: s.db}
			},
			args: args{
				ctx: context.Background(),
				entity: &domain.WalletEntity{
					ID:        1,
					UserID:    1,
					Balance:   float64(10000),
					CreatedAt: 1282312835,
					UpdatedAt: 1282312835,
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			conn := tt.fields(t)

			repo := NewWalletRepository(conn.gorm)
			got := repo.Update(tt.args.ctx, tt.args.entity)
			if got != tt.want {
				t.Errorf("WalletRepository.Update() got = %v, want %v", got, tt.want)
				return
			}

			sqlDB, _ := conn.gorm.DB()
			sqlDB.Close()
		})
	}
}
