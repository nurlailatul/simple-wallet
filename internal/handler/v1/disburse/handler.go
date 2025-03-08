package disburse

import (
	"errors"
	"simple-wallet/internal/handler/v1/response"
	transactionDomain "simple-wallet/internal/module/transaction/domain"
	transactionService "simple-wallet/internal/module/transaction/service"
	userService "simple-wallet/internal/module/user/service"
	walletService "simple-wallet/internal/module/wallet/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	transactionService transactionService.TransactionServiceInterface
	userService        userService.UserServiceInterface
	walletService      walletService.WalletServiceInterface
}

func NewDisburseHandler(
	transactionService transactionService.TransactionServiceInterface,
	userService userService.UserServiceInterface,
	walletService walletService.WalletServiceInterface,
) *HTTPHandler {
	return &HTTPHandler{
		transactionService: transactionService,
		userService:        userService,
		walletService:      walletService,
	}
}

// Deduct Balance Data
//
//	@Router			/api/v1/wallets/{user_id}/deduct [post]
//	@Summary		Create deduct wallet balance transaction
//	@Description	Create deduct wallet balance transaction
//	@Tags			v1
//	@Accept			json
//	@Produce		json
//	@Param			x-api-key	header		string					true	"Insert your api key"	default(<wallet-api-key>)
//	@Param			user_id		path		int						true	"User ID"
//	@Param			data		body		CreateDisburseRequest	true	"Create deduct balance transaction request data"
//	@Success		200			{object}	CreateDisburseResponse
//	@Success		400			{object}	response.CustomError
//	@Success		404			{object}	response.CustomError
//	@Success		422			{object}	response.CustomError
func (h *HTTPHandler) createDisbursement(c *gin.Context) {
	ctx := c.Request.Context()

	request := CreateDisburseRequest{}
	if err := c.BindJSON(&request); err != nil {
		response.SendError(c, response.ErrInvalidArgument, err)
		return
	}

	userID, _ := strconv.ParseInt(c.Param("user_id"), 10, 64)
	user := h.userService.GetByID(ctx, userID)
	if user == nil {
		response.SendError(c, response.ErrRecordNotFound, errors.New("user not found"))
		return
	}

	wallet := h.walletService.GetByUserID(ctx, userID)
	if wallet == nil {
		response.SendError(c, response.ErrRecordNotFound, errors.New("wallet not found"))
		return
	}

	existingTrx := h.transactionService.GetByReferenceID(ctx, request.ReferenceID.String())
	if existingTrx != nil {
		response.SendError(c, response.ErrRequestUnprocessed, errors.New("reference_id exist"))
		return
	}

	req := transactionDomain.DeductBalanceRequest{
		UserID:                userID,
		WalletID:              wallet.ID,
		Amount:                request.Amount,
		ReceiverBank:          request.ReceiverBank,
		ReceiverAccountNumber: request.ReceiverAccountNumber,
		ReferenceID:           request.ReferenceID.String(),
	}

	err := h.transactionService.DeductBalance(ctx, req)
	if err != nil {
		response.SendError(c, response.ErrRequestUnprocessed, err)
		return
	}

	response.SendSuccess(c, nil, nil)
}
