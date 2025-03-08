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

func (h *HTTPHandler) testHTTP(c *gin.Context) {
	response.SendSuccess(c, "Success", nil)
}

func (h *HTTPHandler) createDisbursement(c *gin.Context) {
	response.SendSuccess(c, "Success", nil)
}
