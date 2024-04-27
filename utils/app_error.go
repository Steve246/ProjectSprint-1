package utils

import (
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
