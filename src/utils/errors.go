package utils

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

type ValidationError struct {
	FailedField string      `json:"failed_field"`
	Tag         string      `json:"tag"`
	Value       interface{} `json:"value"`
}

type AppError struct {
	Code             int               `json:"code"`
	Message          string            `json:"message"`
	Details          string            `json:"details", omitempty`
	ValidationErrors []ValidationError `json:"validation_errors", omitempty`
}

func getMessageFromCode(code int) string {
	switch code {
	case 400:
		return "Bad Request"
	case 401:
		return "Unauthorized"
	case 403:
		return "Forbidden"
	case 404:
		return "Not Found"
	case 500:
		return "Internal Server Error"
	}
	return "Unknown Error"
}

func NewAppError(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: getMessageFromCode(code),
		Details: message,
	}
}

func NewAppValidationError(code int, errors []ValidationError) *AppError {
	return &AppError{
		ValidationErrors: errors,
		Code:             code,
		Message:          getMessageFromCode(code),
	}
}

func (me *AppError) ToFiberError() *fiber.Error {
	data, err := json.Marshal(me)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}
	return fiber.NewError(me.Code, string(data))
}
