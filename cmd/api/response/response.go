package response

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AppErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type AppSuccessResponse struct {
	Success  bool        `json:"success"`
	Response interface{} `json:"response"`
}

func NewSuccessResponse(data interface{}) *AppSuccessResponse {
	return &AppSuccessResponse{Success: true, Response: data}
}

func NewErrorResponse(message string) *AppErrorResponse {
	return &AppErrorResponse{Success: false, Message: message}
}

func SuccessResponseHandler(c *fiber.Ctx, data interface{}) error {
	c.Status(http.StatusOK)
	return c.JSON(NewSuccessResponse(data))
}

func BadRequestHandler(c *fiber.Ctx, message string) error {
	c.Status(http.StatusBadRequest)
	return c.JSON(NewErrorResponse(message))

}

func NotFoundRequestHandler(c *fiber.Ctx, message string) error {
	c.Status(http.StatusNotFound)
	return c.JSON(NewErrorResponse(message))

}

func UnAuthorizedRequestHandler(c *fiber.Ctx, message string) error {
	c.Status(http.StatusUnauthorized)
	return c.JSON(NewErrorResponse(message))

}

func UnexpectedErrorHandler(c *fiber.Ctx, message string) error {
	c.Status(http.StatusInternalServerError)
	return c.JSON(NewErrorResponse(message))

}
