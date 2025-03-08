package disburse

import (
	"simple-wallet/internal/handler/v1/response"
	userService "simple-wallet/internal/module/user/service"

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

// Ping godoc
//
//	@Router			/api/v1/ping [get]
//	@Summary		Ping
//	@Description	Ping
//	@Accept			json
//	@Produce		json
//	@Success		200
func (h *HTTPHandler) testHTTP(c *gin.Context) {
	response.SendSuccess(c, "Success", nil)
}

func (h *HTTPHandler) createDisbursement(c *gin.Context) {
	response.SendSuccess(c, "Success", nil)
}
