package response

import (
	"net/http"
)

var (
	GeneralErrorCode            = 1000
	RecordNotFoundErrorCode     = 1001
	InvalidArgumentErrorCode    = 1002
	RequestUnprocessedErrorCode = 1003
)

var (
	// unexpected error
	ErrGeneral = CustomError{
		HttpCode:  http.StatusInternalServerError,
		ErrorCode: GeneralErrorCode,
		Status:    "GENERAL_ERROR",
	}

	ErrRecordNotFound = CustomError{
		HttpCode:  http.StatusNotFound,
		ErrorCode: RecordNotFoundErrorCode,
		Status:    "RECORD_NOT_FOUND",
	}

	ErrInvalidArgument = CustomError{
		HttpCode:  http.StatusUnprocessableEntity,
		ErrorCode: InvalidArgumentErrorCode,
		Status:    "INVALID_ARGUMENTS",
		Message:   "Error processing request with message: Invalid request body, parse error",
	}

	ErrInvalidParameter = CustomError{
		HttpCode:  http.StatusBadRequest,
		ErrorCode: InvalidParameterCode,
		Status:    "INVALID_PARAMETER",
	}

	ErrRequestUnprocessed = CustomError{
		HttpCode:  http.StatusBadRequest,
		ErrorCode: RequestUnprocessedErrorCode,
		Status:    "UNPROCESSED_REQUEST",
		Message:   "Failed processing request",
	}
)

const (
	NoCode                 = 0
	RateLimitExceededCode  = 4000
	InvalidParameterCode   = 4001
	InvalidImageSizeCode   = 4007
	InvalidPhoneNumberCode = 4030
	InvalidNIKCode         = 4031
	InvalidDOBCode         = 4035
)

var (
	RateLimitExceeded  = ErrorDetail{4000, "Rate limit exceeded", nil}
	InvalidParameter   = ErrorDetail{4001, "Invalid parameter", nil}
	InvalidImageSize   = ErrorDetail{4007, "Image size must less than 1MB", nil}
	InvalidPhoneNumber = ErrorDetail{4030, "Invalid phone number", nil}
	InvalidNIK         = ErrorDetail{4031, "Invalid NIK", nil}
	InvalidDOB         = ErrorDetail{4035, "Invalid date of birth format, it must be yyyy-mm-dd", nil}
)
