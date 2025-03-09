package repository

import (
	"context"
	"database/sql"
	"reflect"
	"regexp"
	"testing"

	"simple-wallet/internal/module/user/domain"
	"simple-wallet/pkg/db"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	userColumn = []string{"id", "phone", "email", "name", "status", "created_at", "updated_at"}
)

type UserRepositorySuite struct {
	suite.Suite
	db   *gorm.DB
	mock sqlmock.Sqlmock
}

func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserRepositorySuite))
}

func (s *UserRepositorySuite) SetupSuite(t *testing.T) *UserRepositorySuite {
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

func (s *UserRepositorySuite) TestUserRepository_GetByID() {
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
		want   *domain.UserEntity
	}{
		{
			name: "success",
			fields: func(t *testing.T) fields {
				s := new(UserRepositorySuite).SetupSuite(t)

				s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE id = ? ORDER BY `users`.`id` LIMIT ?")).
					WithArgs(1, 1).
					WillReturnRows(sqlmock.NewRows(userColumn).
						AddRow(1, "+6281234567890", "nurlailatul@gmail.com", "Nur Lailatul", 1, 1282312835, 1282312835))

				return fields{gorm: s.db}
			},
			args: args{
				ctx:    context.Background(),
				userID: 1,
			},
			want: &domain.UserEntity{
				ID:        1,
				Phone:     "+6281234567890",
				Name:      "Nur Lailatul",
				Email:     "nurlailatul@gmail.com",
				Status:    1,
				CreatedAt: 1282312835,
				UpdatedAt: 1282312835,
			},
		},
		{
			name: "failed",
			fields: func(t *testing.T) fields {
				s := new(UserRepositorySuite).SetupSuite(t)

				s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE id = ? ORDER BY `users`.`id` LIMIT ?")).
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

			repo := NewUserRepository(db.NewGormDBWrapper(conn.gorm, "", 10))
			got := repo.GetByID(tt.args.ctx, tt.args.userID)
			if got == nil && tt.want != nil {
				t.Errorf("UserRepository.GetByID() got = %v, want %v", got, tt.want)
				return
			}
			if got != nil && !reflect.DeepEqual(*got, *tt.want) {
				t.Errorf("UserRepository.GetByID() got = %v, want %v", got, tt.want)
			}

			sqlDB, _ := conn.gorm.DB()
			sqlDB.Close()
		})
	}
}
