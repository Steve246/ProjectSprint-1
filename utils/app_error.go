package utils

import (
	"errors"
	"fmt"
	"net/http"
)

type AppError struct {
	ErrorCode    string
	ErrorMessage string
	ErrorType    int
}

func (e AppError) Error() string {
	return fmt.Sprintf("type: %d, code: %s, err: %s", e.ErrorType, e.ErrorCode, e.ErrorMessage)
}

var (

	// REGISTER
	ErrEmailNull           = errors.New("email cannot be null")
	ErrInvalidEmail        = errors.New("invalid email format")
	ErrNameNull            = errors.New("name cannot be null")
	ErrInvalidName         = errors.New("name must be between 5 and 50 characters")
	ErrPasswordNull        = errors.New("password cannot be null")
	ErrInvalidPassword     = errors.New("password must be between 5 and 15 characters")
	ErrDuplicateValueFound = errors.New("duplicate Email is found")

	// LOGIN
	ErrEmailCannotFound    = errors.New("email cannot be found")
	ErrPasswordCannotFound = errors.New("password cannot be found")
	ErrPasswordNotMatch    = errors.New("password not match")
	ErrUserNotFound        = errors.New("user not found")
)

func IsValidationError(err error) bool {
	return errors.Is(err, ErrEmailNull) ||
		errors.Is(err, ErrInvalidEmail) ||
		errors.Is(err, ErrNameNull) ||
		errors.Is(err, ErrInvalidName) ||
		errors.Is(err, ErrPasswordNull) ||
		errors.Is(err, ErrInvalidPassword)
}

func IsErrDuplicateValueFound(err error) bool {
	return errors.Is(err, ErrDuplicateValueFound)
}

// register
func EmailFoundError() error {
	return AppError{
		ErrorCode:    "409",
		ErrorMessage: "Email found inside Database",
		ErrorType:    http.StatusConflict,
	}
}

func ReqBodyNotValidError() error {
	return AppError{
		ErrorCode:    "400",
		ErrorMessage: "Didn't pass Validation",
		ErrorType:    http.StatusBadRequest,
	}
}

func ServerError() error {
	return AppError{
		ErrorCode:    "500",
		ErrorMessage: "Server Error",
		ErrorType:    http.StatusInternalServerError,
	}
}

// login

func PasswordCannotBeEncodeError() error {
	return AppError{
		ErrorCode:    "400",
		ErrorMessage: "Password cannot be encode",
		ErrorType:    http.StatusBadRequest,
	}
}

func UserNotFoundError() error {
	return AppError{
		ErrorCode:    "404",
		ErrorMessage: "User Not Found",
		ErrorType:    http.StatusInternalServerError,
	}
}

func PasswordWrongError() error {
	return AppError{
		ErrorCode:    "400",
		ErrorMessage: "Password Is Wrong",
		ErrorType:    http.StatusInternalServerError,
	}
}

// ini yg lama

func EmailDuplicate() error {
	return AppError{
		ErrorCode:    "409",
		ErrorMessage: "Email found inside Database",
		ErrorType:    http.StatusBadRequest,
	}
}

func ValidationError() error {
	return AppError{
		ErrorCode:    "400",
		ErrorMessage: "Email found inside Database",
		ErrorType:    http.StatusBadRequest,
	}
}

// error code lama

func RequiredError() error {
	return AppError{
		ErrorCode:    "X01",
		ErrorMessage: "Input can't be Empty",
		ErrorType:    http.StatusBadRequest,
	}
}

func DataNotFoundError() error {
	return AppError{
		ErrorCode:    "X02",
		ErrorMessage: "No Data Found",
		ErrorType:    http.StatusBadRequest,
	}
}

func DataDuplicateError() error {
	return AppError{
		ErrorCode:    "X03",
		ErrorMessage: "Duplicate value found",
		ErrorType:    http.StatusBadRequest,
	}
}

func UnauthorizedError() error {
	return AppError{
		ErrorCode:    "X04",
		ErrorMessage: "Unauthorized user",
		ErrorType:    http.StatusUnauthorized,
	}
}

func WrongOtpError() error {
	return AppError{
		ErrorCode:    "X05",
		ErrorMessage: "Wrong OTP",
		ErrorType:    http.StatusUnauthorized,
	}
}

func UnsupportedFileExtensionError() error {
	return AppError{
		ErrorCode:    "X06",
		ErrorMessage: "Unsupported file extensions",
		ErrorType:    http.StatusBadRequest,
	}
}

func ActivationFailed() error {
	return AppError{
		ErrorCode:    "X07",
		ErrorMessage: "Activation Failed",
		ErrorType:    http.StatusUnauthorized,
	}
}

func UnsufficientBalance() error {
	return AppError{
		ErrorCode:    "X08",
		ErrorMessage: "Balance must greater than 0",
		ErrorType:    http.StatusBadRequest,
	}
}

func ParamsIncomplete() error {
	return AppError{
		ErrorCode:    "X09",
		ErrorMessage: "Params incomplete",
		ErrorType:    http.StatusBadRequest,
	}
}

func VariableNotFound(value string) error {
	return AppError{
		ErrorCode:    "X10",
		ErrorMessage: value + " not found",
		ErrorType:    http.StatusBadRequest,
	}
}

func ErrorSelect(value string) error {
	return AppError{
		ErrorCode:    "X11",
		ErrorMessage: "error while get " + value,
		ErrorType:    http.StatusBadRequest,
	}
}

func InvalidDateFormat() error {
	return AppError{
		ErrorCode:    "X12",
		ErrorMessage: "Invalid date format",
		ErrorType:    http.StatusBadRequest,
	}
}

func InvalidTypeFormat() error {
	return AppError{
		ErrorCode:    "X13",
		ErrorMessage: "Invalid type format",
		ErrorType:    http.StatusBadRequest,
	}
}

func PasswordNotMatch() error {
	return AppError{
		ErrorCode:    "X14",
		ErrorMessage: "Wrong Password",
		ErrorType:    http.StatusBadRequest,
	}
}
