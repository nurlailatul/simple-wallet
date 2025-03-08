package disburse

import (
	"errors"
	"simple-wallet/internal/handler/v1/response"
	userService "simple-wallet/internal/module/user/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	userService userService.UserServiceInterface
}

func NewDisburseHandler(userService userService.UserServiceInterface) *HTTPHandler {
	return &HTTPHandler{
		userService: userService,
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
//	@Param			X-API-KEY	header		string					true	"Insert your api key"	default(<wallet-api-key>)
//	@Param			user_id		path		int						true	"User ID"
//	@Param			data		body		CreateDisburseRequest	true	"Create deduct balance transaction request data"
//	@Success		200			{object}	CreateDisburseResponse
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
	user := h.userService.GetById(ctx, userID)
	if user == nil {
		response.SendError(c, response.ErrRecordNotFound, errors.New("user not found"))
		return
	}

	response.SendSuccess(c, nil, nil)
	response.SendSuccess(c, "Success", nil)
}
