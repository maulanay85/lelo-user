package util

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type ResponseEntity struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var (
	ErrorUnauthenticated     = errors.New("unauthicated")
	ErrorUnauthorized        = errors.New("unathorized")
	ErrorErrorBadRequest     = errors.New("bad request")
	ErrorNotFound            = errors.New("not found")
	ErrorPreCondition        = errors.New("pre condition")
	ErrorUnprocessableEntity = errors.New("un processable")
	ErrorInternalServer      = errors.New("internal server")
	ErrorCustom              = errors.New("custom error")
)

func SendSuccess(c *gin.Context, data interface{}) {
	res := ResponseEntity{Code: 200, Message: "", Data: data}
	c.JSON(200, res)
}

func SendSuccessWithMessage(c *gin.Context, message string, data interface{}) {
	res := ResponseEntity{Code: 200, Message: message, Data: data}
	c.JSON(200, res)
}

func SendErrorResponse(c *gin.Context, err error) {
	unwrap := errors.Unwrap(err)

	if errors.Is(unwrap, ErrorUnauthenticated) {
		res := ResponseEntity{Code: 401, Message: err.Error()}
		c.JSON(401, res)
	} else if errors.Is(unwrap, ErrorUnauthorized) {
		res := ResponseEntity{Code: 403, Message: err.Error()}
		c.JSON(403, res)
	} else if errors.Is(unwrap, ErrorNotFound) {
		res := ResponseEntity{Code: 404, Message: err.Error()}
		c.JSON(404, res)
	} else if errors.Is(unwrap, ErrorPreCondition) {
		res := ResponseEntity{Code: 412, Message: err.Error()}
		c.JSON(412, res)
	} else if errors.Is(unwrap, ErrorUnprocessableEntity) {
		res := ResponseEntity{Code: 422, Message: err.Error()}
		c.JSON(422, res)
	} else {
		res := ResponseEntity{Code: 500, Message: err.Error()}
		c.JSON(500, res)
	}
}

// type ErrorResponse struct {
// 	message string
// }

// func (e *ErrorResponse) Unauthorized() error {
// 	return fmt.Errorf(e.message, ErrorUnauthorized)
// }

// func (e *ErrorResponse) Unauthenticated() error {
// 	return fmt.Errorf(e.message, ErrorUnauthenticated)
// }

// func (e *ErrorResponse) BadRequest() error {
// 	return fmt.Errorf(e.message, ErrorErrorBadRequest)
// }

// func (e *ErrorResponse) NotFound() error {
// 	return fmt.Errorf(e.message, ErrorNotFound)
// }

// func (e *ErrorResponse) PreCondition() error {
// 	return fmt.Errorf(e.message, ErrorPreCondition)
// }

// func (e *ErrorResponse) UnprocessableEntity() error {
// 	return fmt.Errorf(e.message, ErrorUnprocessableEntity)
// }

// func (e *ErrorResponse) InternalServer() error {
// 	return fmt.Errorf(e.message, ErrorInternalServer)
// }

// func (e *ErrorResponse) Custom(code int16, message string, data interface{}) error {
// 	return fmt.Errorf(e.message, ErrorCustom)
// }

// func SendErrorResponse(c *gin.Context) {
// 	c.JSON()
// }
