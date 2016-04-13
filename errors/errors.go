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
	// Unauthorized error is for when Login fails
	Unauthorized = "UNAUTHORIZED"
	// DuplicateUsername is for when a username is taken and can't create a new one
	DuplicateUsername = "DUPLICATE_USERNAME"
)

var statusCode = map[ErrorCode]int{
	DatabaseError:      503,
	GeneralServerError: 500,
	InvalidID:          404,
	UserNotFound:       401,
	GeneraJSONError:    400,
	Unauthorized:       401,
	DuplicateUsername:  409,
}

func NewApiError(errorCode ErrorCode, msg string) *echo.HTTPError {
	return echo.NewHTTPError(statusCode[errorCode], fmt.Sprintln(string(errorCode), "description", msg))
}
