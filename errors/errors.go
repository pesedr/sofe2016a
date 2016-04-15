package errors

type ErrorCode string

const (
	// Database Error
	DatabaseError = "DB_ERROR"
)

var errStatusCode = map[ErrorCode]int{
	DatabaseError: 404,
}

type ApiError struct {
	Error       error  `json:"error"`
	Code        int    `json:"code"`
	Description string `json:description`
}

func NewAppError() {

}
