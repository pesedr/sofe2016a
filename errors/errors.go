package errors

import (
	"fmt"

	"github.com/labstack/echo"
)

type ErrorCode string

const (
	// DatabaseError contains all errors that have to do with mongoDB
	DatabaseError = "DB_ERROR"
	// GeneralServerError contains errors that have to do with echo framework
	GeneralServerError = "GENERAL_SERVER_ERROR"
	// InvalidID is an error for an object not found in the database
	InvalidID = "INVALID_ID"
	// UserNotFound is an error for when the user was not found in DB
	UserNotFound = "USER_NOT_FOUND"
	// GeneraJSONError is for when a bad request is sent
	GeneraJSONError = "GENERAL_JSON_ERROR"
)

var statusCode = map[ErrorCode]int{
	DatabaseError:      503,
	GeneralServerError: 500,
	InvalidID:          404,
	UserNotFound:       404,
	GeneraJSONError:    400,
}

func NewApiError(errorCode ErrorCode, msg string) *echo.HTTPError {
	return echo.NewHTTPError(statusCode[errorCode], fmt.Sprintln(string(errorCode), "message:", msg))
}
