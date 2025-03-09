package disburse

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	transactionDomain "simple-wallet/internal/module/transaction/domain"
	transactionMocks "simple-wallet/internal/module/transaction/mocks"
	userDomain "simple-wallet/internal/module/user/domain"
	userMocks "simple-wallet/internal/module/user/mocks"
	walletDomain "simple-wallet/internal/module/wallet/domain"
	walletMocks "simple-wallet/internal/module/wallet/mocks"
)

func TestNewDisburseHandler(t *testing.T) {
	mockUserService := new(userMocks.UserRepositoryInterface)
	mockWalletService := new(walletMocks.WalletRepositoryInterface)
	mockTransactionService := new(transactionMocks.TransactionServiceInterface)

	h := NewDisburseHandler(mockTransactionService, mockUserService, mockWalletService)

	assert.NotNil(t, h)
	assert.Equal(t, mockTransactionService, h.transactionService)
	assert.Equal(t, mockUserService, h.userService)
	assert.Equal(t, mockWalletService, h.walletService)
}

func setupHandler() HTTPHandler {
	mockUserService := new(userMocks.UserRepositoryInterface)
	mockWalletService := new(walletMocks.WalletRepositoryInterface)
	mockTransactionService := new(transactionMocks.TransactionServiceInterface)

	return HTTPHandler{
		userService:        mockUserService,
		walletService:      mockWalletService,
		transactionService: mockTransactionService,
	}
}

func TestCreateDisbursement(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	mockUserService := new(userMocks.UserRepositoryInterface)
	mockWalletService := new(walletMocks.WalletRepositoryInterface)
	mockTransactionService := new(transactionMocks.TransactionServiceInterface)

	h := HTTPHandler{
		userService:        mockUserService,
		walletService:      mockWalletService,
		transactionService: mockTransactionService,
	}

	r.POST("/disburse/:user_id", h.createDisbursement)

	// Sample request
	requestBody := CreateDisburseRequest{
		Amount:                1000,
		ReceiverBank:          "BCA",
		ReceiverAccountNumber: "1234567890",
		ReferenceID:           "ref123",
	}
	jsonValue, _ := json.Marshal(requestBody)

	t.Run("Success", func(t *testing.T) {
		userID := int64(1)
		walletID := int64(1)

		mockUserService.On("GetByID", mock.Anything, userID).Return(&userDomain.UserEntity{ID: userID})
		mockWalletService.On("GetByUserID", mock.Anything, userID).Return(&walletDomain.WalletEntity{ID: walletID})
		mockTransactionService.On("GetByReferenceID", mock.Anything, requestBody.ReferenceID).Return(nil)
		mockTransactionService.On("DeductBalance", mock.Anything, mock.Anything).Return(nil)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest(http.MethodPost, "/disburse/1", bytes.NewBuffer(jsonValue))
		c.Params = gin.Params{{Key: "user_id", Value: strconv.FormatInt(userID, 10)}}

		h.createDisbursement(c)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("User_Not_Found", func(t *testing.T) {
		userID := int64(2)
		mockUserService.On("GetByID", mock.Anything, userID).Return(nil)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest(http.MethodPost, "/disburse/2", bytes.NewBuffer(jsonValue))
		c.Params = gin.Params{{Key: "user_id", Value: strconv.FormatInt(userID, 10)}}

		h.createDisbursement(c)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("Wallet_Not_Found", func(t *testing.T) {
		userID := int64(4)

		mockUserService.On("GetByID", mock.Anything, userID).Return(&userDomain.UserEntity{ID: userID})
		mockWalletService.On("GetByUserID", mock.Anything, userID).Return(nil)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest(http.MethodPost, "/disburse/4", bytes.NewBuffer(jsonValue))
		c.Params = gin.Params{{Key: "user_id", Value: strconv.FormatInt(userID, 10)}}

		h.createDisbursement(c)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("Reference_id_exist", func(t *testing.T) {
		userID := int64(4)
		walletID := int64(4)

		mockUserService.On("GetByID", mock.Anything, userID).Return(&userDomain.UserEntity{ID: userID})
		mockWalletService.On("GetByUserID", mock.Anything, userID).Return(&walletDomain.WalletEntity{ID: walletID})
		mockTransactionService.On("GetByReferenceID", mock.Anything, requestBody.ReferenceID).Return(&transactionDomain.TransactionEntity{ID: 1})

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest(http.MethodPost, "/disburse/4", bytes.NewBuffer(jsonValue))
		c.Params = gin.Params{{Key: "user_id", Value: strconv.FormatInt(userID, 10)}}

		h.createDisbursement(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("DeductBalance_Error", func(t *testing.T) {
		userID := int64(4)
		walletID := int64(4)

		mockUserService.On("GetByID", mock.Anything, userID).Return(&userDomain.UserEntity{ID: userID})
		mockWalletService.On("GetByUserID", mock.Anything, userID).Return(&walletDomain.WalletEntity{ID: walletID})
		mockTransactionService.On("GetByReferenceID", mock.Anything, requestBody.ReferenceID).Return(nil)
		mockTransactionService.On("DeductBalance", mock.Anything, mock.Anything).Return(errors.New("deduction failed"))

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest(http.MethodPost, "/disburse/4", bytes.NewBuffer(jsonValue))
		c.Params = gin.Params{{Key: "user_id", Value: strconv.FormatInt(userID, 10)}}

		h.createDisbursement(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
