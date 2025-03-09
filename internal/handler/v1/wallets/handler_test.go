package wallets

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	transactionDomain "simple-wallet/internal/module/transaction/domain"
	transactionMocks "simple-wallet/internal/module/transaction/mocks"
	userDomain "simple-wallet/internal/module/user/domain"
	userMocks "simple-wallet/internal/module/user/mocks"
	walletDomain "simple-wallet/internal/module/wallet/domain"
	walletMocks "simple-wallet/internal/module/wallet/mocks"
)

type WalletTestSuite struct {
	suite.Suite
	userService        *userMocks.UserRepositoryInterface
	walletService      *walletMocks.WalletServiceInterface
	transactionService *transactionMocks.TransactionServiceInterface
}

func TestWalletHandler(t *testing.T) {
	suite.Run(t, new(WalletTestSuite))
}

func (s *WalletTestSuite) SetupSuite(t *testing.T) *WalletTestSuite {
	s.Suite.SetT(t)
	s.userService = new(userMocks.UserRepositoryInterface)
	s.walletService = new(walletMocks.WalletServiceInterface)
	s.transactionService = new(transactionMocks.TransactionServiceInterface)

	return s
}

func (s *WalletTestSuite) TestWalletHandler_createWallet() {
	requestBody := CreateDisburseRequest{
		Amount:                1000,
		ReceiverBank:          "BCA",
		ReceiverAccountNumber: "1234567890",
		ReferenceID:           "ref123",
	}

	userID := int64(1)
	walletID := int64(1)

	type fields struct {
		userService        *userMocks.UserRepositoryInterface
		walletService      *walletMocks.WalletServiceInterface
		transactionService *transactionMocks.TransactionServiceInterface
	}
	type args struct {
		ctx     context.Context
		userID  int64
		request CreateDisburseRequest
	}
	tests := []struct {
		name   string
		fields func(t *testing.T) fields
		args   args
		want   int
	}{
		{
			name: "Success",
			fields: func(t *testing.T) fields {
				s := new(WalletTestSuite).SetupSuite(t)

				s.userService.On("GetByID", mock.Anything, userID).Return(&userDomain.UserEntity{ID: userID})
				s.walletService.On("GetByUserID", mock.Anything, userID).Return(&walletDomain.WalletEntity{ID: walletID})
				s.transactionService.On("GetByReferenceID", mock.Anything, requestBody.ReferenceID).Return(nil)
				s.transactionService.On("DeductBalance", mock.Anything, mock.Anything).Return(nil)

				return fields{
					userService:        s.userService,
					walletService:      s.walletService,
					transactionService: s.transactionService,
				}
			},
			args: args{
				ctx:     context.Background(),
				userID:  userID,
				request: requestBody,
			},
			want: http.StatusOK,
		},
		{
			name: "DeductBalanceError",
			fields: func(t *testing.T) fields {
				s := new(WalletTestSuite).SetupSuite(t)

				s.userService.On("GetByID", mock.Anything, userID).Return(&userDomain.UserEntity{ID: userID})
				s.walletService.On("GetByUserID", mock.Anything, userID).Return(&walletDomain.WalletEntity{ID: walletID})
				s.transactionService.On("GetByReferenceID", mock.Anything, requestBody.ReferenceID).Return(nil)
				s.transactionService.On("DeductBalance", mock.Anything, mock.Anything).Return(errors.New("error deduct wallet balance"))

				return fields{
					userService:        s.userService,
					walletService:      s.walletService,
					transactionService: s.transactionService,
				}
			},
			args: args{
				ctx:     context.Background(),
				userID:  userID,
				request: requestBody,
			},
			want: http.StatusBadRequest,
		},
		{
			name: "ReferenceIDExist",
			fields: func(t *testing.T) fields {
				s := new(WalletTestSuite).SetupSuite(t)

				s.userService.On("GetByID", mock.Anything, userID).Return(&userDomain.UserEntity{ID: userID})
				s.walletService.On("GetByUserID", mock.Anything, userID).Return(&walletDomain.WalletEntity{ID: walletID})
				s.transactionService.On("GetByReferenceID", mock.Anything, requestBody.ReferenceID).Return(&transactionDomain.TransactionEntity{ID: 1})

				return fields{
					userService:        s.userService,
					walletService:      s.walletService,
					transactionService: s.transactionService,
				}
			},
			args: args{
				ctx:     context.Background(),
				userID:  userID,
				request: requestBody,
			},
			want: http.StatusBadRequest,
		},
		{
			name: "WalletNotFound",
			fields: func(t *testing.T) fields {
				s := new(WalletTestSuite).SetupSuite(t)

				s.userService.On("GetByID", mock.Anything, userID).Return(&userDomain.UserEntity{ID: userID})
				s.walletService.On("GetByUserID", mock.Anything, userID).Return(nil)

				return fields{
					userService:        s.userService,
					walletService:      s.walletService,
					transactionService: s.transactionService,
				}
			},
			args: args{
				ctx:     context.Background(),
				userID:  userID,
				request: requestBody,
			},
			want: http.StatusNotFound,
		},
		{
			name: "UserNotFound",
			fields: func(t *testing.T) fields {
				s := new(WalletTestSuite).SetupSuite(t)

				s.userService.On("GetByID", mock.Anything, userID).Return(nil)

				return fields{
					userService:        s.userService,
					walletService:      s.walletService,
					transactionService: s.transactionService,
				}
			},
			args: args{
				ctx:     context.Background(),
				userID:  userID,
				request: requestBody,
			},
			want: http.StatusNotFound,
		},
		{
			name: "Error Params",
			fields: func(t *testing.T) fields {
				s := new(WalletTestSuite).SetupSuite(t)

				return fields{
					userService:        s.userService,
					walletService:      s.walletService,
					transactionService: s.transactionService,
				}
			},
			args: args{
				ctx:    context.Background(),
				userID: userID,
				request: CreateDisburseRequest{
					ReceiverBank:          "BCA",
					ReceiverAccountNumber: "1234567890",
					ReferenceID:           "ref123",
				},
			},
			want: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			mocks := tt.fields(t)
			h := NewWalletHandler(mocks.transactionService, mocks.userService, mocks.walletService)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			payload, _ := json.Marshal(tt.args.request)
			c.Request, _ = http.NewRequest(http.MethodPost, "/wallets/"+strconv.FormatInt(tt.args.userID, 10), bytes.NewBuffer(payload))
			c.Params = gin.Params{{Key: "user_id", Value: strconv.FormatInt(tt.args.userID, 10)}}

			h.createDisbursement(c)

			assert.Equal(t, tt.want, w.Code)
		})
	}
}
