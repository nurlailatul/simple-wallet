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

// Get Identity Image URL godoc
//
//	@Router			/bkk/v1/users/{user_id}/identity [get]
//	@Summary		Get Identity Image URL Data
//	@Description	Get Identity Image URL Data
//	@Tags			Bangkokok
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"Insert your access token"	default(Bearer <Add access token here>)
//	@Param			user_id			path		int		true	"User ID"
//	@Param			type			query		string	true	"Identity Type"	Enums(identity, face_and_identity)
//	@Success		200				{object}	user.GetIdentityImageResponseSingleResponse
func (h *HTTPHandler) getIdentityImageData(c *gin.Context) {
	// ctx := c.Request.Context()

	// userID, _ := strconv.ParseInt(fmt.Sprint(c.Param("id")), 10, 64)
	// docType := c.Query("type")

	response.SendSuccess(c, "Success", nil)
}

func (h *HTTPHandler) testHTTP(c *gin.Context) {
	// ctx := c.Request.Context()

	// userID, _ := strconv.ParseInt(fmt.Sprint(c.Param("id")), 10, 64)
	// docType := c.Query("type")

	response.SendSuccess(c, "Success", nil)
}
