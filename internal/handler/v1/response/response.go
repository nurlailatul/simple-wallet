package response

import (
	"net/http"
	"reflect"

	"simple-wallet/internal/core"

	"github.com/gin-gonic/gin"
)

var (
	GenericSuccessResponse = 2002
	GenericSuccessMessage  = "Success"
)

type JsonResponse struct {
	Code       int         `json:"code"`
	Status     string      `json:"status,omitempty"`
	Message    string      `json:"message,omitempty"`
	Pagination interface{} `json:"_meta,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Error      interface{} `json:"errors,omitempty"`
}

type CustomError struct {
	HttpCode  int    `json:"-"`
	ErrorCode int    `json:"code"`
	Status    string `json:"status"`
	Message   string `json:"message"`
}

type ErrorDetail struct {
	Code       int         `json:"code"`
	Message    string      `json:"message"`
	Additional interface{} `json:"additional,omitempty"`
}

type ResponseCounter struct {
	Counter string `json:"counter"`
}

func NewJSON() *JsonResponse {
	return &JsonResponse{
		Code:    GenericSuccessResponse,
		Message: GenericSuccessMessage,
	}
}

func NewError() *JsonResponse {
	return &JsonResponse{}
}

func (j *JsonResponse) SetPagination(pageResponse *core.PageResponse) *JsonResponse {
	j.Pagination = pageResponse
	return j
}

func (j *JsonResponse) SetMessage(message string) *JsonResponse {
	j.Message = message
	return j
}

func (j *JsonResponse) SetErrorMessage(err error) *JsonResponse {
	j.Message = err.Error()
	return j
}

func (j *JsonResponse) SetErrorDetail(err error) *JsonResponse {
	j.Error = err.Error()
	return j
}

// SetData set object data
func (j *JsonResponse) SetData(model interface{}) *JsonResponse {
	j.Data = model
	if model == nil {
		j.Data = []string{}
	}
	return j
}

// SetData set list data to show as array
func (j *JsonResponse) SetListData(model interface{}) *JsonResponse {
	j.Data = model
	if model == nil || (reflect.ValueOf(model).Kind() == reflect.Ptr && reflect.ValueOf(model).Elem().IsNil()) || (reflect.ValueOf(model).Kind() == reflect.Slice && reflect.ValueOf(model).IsZero()) {
		j.Data = []string{}
	}
	return j
}

func SendSuccess(c *gin.Context, data interface{}, pageResponse *core.PageResponse) {
	resp := NewJSON().SetData(data)
	if pageResponse != nil {
		resp.SetPagination(pageResponse)
	}
	c.JSON(http.StatusOK, resp)
}

func SendSuccessListData(c *gin.Context, data interface{}, pageResponse *core.PageResponse) {
	resp := NewJSON().SetListData(data)
	if pageResponse != nil {
		resp.SetPagination(pageResponse)
	}
	c.JSON(http.StatusOK, resp)
}

func SendError(c *gin.Context, err CustomError, errDetail interface{}) {
	resp := &JsonResponse{
		Code:   err.ErrorCode,
		Status: err.Status,
	}
	if errDetail != nil {
		switch v := errDetail.(type) {
		case error:
			resp.Error = v.Error()
		default:
			resp.Error = v
		}
	}
	if err.Message != "" {
		resp.Message = err.Message
	}
	c.JSON(err.HttpCode, resp)
}
