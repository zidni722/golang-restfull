package response

import (
	"github.com/kataras/iris"
	"github.com/pkg/errors"
	"github.com/zidni722/golang-restfull/app/dto/response"
)

var validations = make([]string, 0)

const (
	OK         = 200
	CREATED    = 201
	ACCEPTED   = 202
	NO_CONTENT = 204

	BAD_REQUEST          = 400
	UN_AUTHORIZED        = 401
	PAYMENT_REQUIRED     = 402
	FORBIDDEN            = 403
	NOT_FOUND            = 404
	METHOD_NOT_ALLOWED   = 405
	NOT_ACCEPTABLE       = 406
	REQUEST_TIMEOUT      = 408
	CONFLICT             = 409
	UNPROCESSABLE_ENTITY = 422

	INTERNAL_SERVER_ERROR = 500
	BAD_GATEWAY           = 502
	SERVICE_UNAVAILABLE   = 503
	GATEWAY_TIMEOUT       = 504
)

const (
	OK_MESSAGE         = "Ok"
	CREATED_MESSAGE    = "Created"
	ACCEPTED_MESSAGE   = "Accepted"
	NO_CONTENT_MESSAGE = "No Content"

	BAD_REQUEST_MESSAGE          = "Bad Request"
	UN_AUTHORIZED_MESSAGE        = "Unauthorized"
	PAYMENT_REQUIRED_MESSAGE     = "Payment Required"
	FORBIDDEN_MESSAGE            = "Forbidden"
	NOT_FOUND_MESSAGE            = "Not Found"
	METHOD_NOT_ALLOWED_MESSAGE   = "Method Not Allowed"
	NOT_ACCEPTABLE_MESSAGE       = "Not Acceptable"
	REQUEST_TIMEOUT_MESSAGE      = "Request Timeout"
	CONFLICT_MESSAGE             = "Conflict"
	UNPROCESSABLE_ENTITY_MESSAGE = "Unprocessable Entity"

	INTERNAL_SERVER_ERROR_MESSAGE = "Internal Server Error"
	BAD_GATEWAY_MESSAGE           = "Bad Gateway"
	SERVICE_UNAVAILABLE_MESSAGE   = "Service Unavailable"
	GATEWAY_TIMEOUT_MESSAGE       = "Gateway Timeout"
	SUCCESS_SAVE_USER             = "User has been Created"
)

func SuccessResponse(ctx iris.Context, status int, message string, data interface{}) {
	response := response.BaseResponse{
		Status:           status,
		Message:          message,
		Data:             data,
		ValidationErrors: validations,
	}

	ctx.StatusCode(status)
	ctx.JSON(response)
}

func ValidationResponse(ctx iris.Context, message string, validationErrors []string) {
	response := response.BaseResponse{
		Status:           BAD_REQUEST,
		Message:          message,
		Data:             nil,
		ValidationErrors: validationErrors,
	}

	ctx.StatusCode(BAD_REQUEST)
	ctx.JSON(response)
}

func InternalServerErrorResponse(ctx iris.Context, err interface{}) {
	response := response.BaseResponse{
		Status:           INTERNAL_SERVER_ERROR,
		Message:          INTERNAL_SERVER_ERROR_MESSAGE,
		Data:             nil,
		ValidationErrors: validations,
	}

	if err != nil {
		errs := errors.Wrap(err.(error), "")
		ctx.Application().Logger().Errorf("%+v", errs)
	}

	ctx.StatusCode(INTERNAL_SERVER_ERROR)
	ctx.JSON(response)
}

func NotFoundResponse(ctx iris.Context, message string) {
	response := response.BaseResponse{
		Status:           NOT_FOUND,
		Message:          message,
		Data:             nil,
		ValidationErrors: validations,
	}

	ctx.StatusCode(NOT_FOUND)
	ctx.JSON(response)
}

func ErrorResponse(ctx iris.Context, status int, message string) {
	response := response.BaseResponse{
		Status:           status,
		Message:          message,
		Data:             nil,
		ValidationErrors: validations,
	}

	ctx.StatusCode(status)
	ctx.JSON(response)
}

func UnAuthorizedResponse(ctx iris.Context) {
	response := response.BaseResponse{
		Status:           UN_AUTHORIZED,
		Message:          UN_AUTHORIZED_MESSAGE,
		Data:             nil,
		ValidationErrors: validations,
	}

	ctx.StatusCode(UN_AUTHORIZED)
	ctx.JSON(response)
}
