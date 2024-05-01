package response

import (
	"7Zero4/utils"
	"errors"

	"net/http"
)

type Status struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
}

type Response struct {
	Status
	Data interface{} `json:"data,omitempty"`
}

type ResponseSuccess struct {
	Data interface{} `json:"data,omitempty"`
}

func NewSuccessMessage(data interface{}) (httpStatusCode int, apiResponse ResponseSuccess) {
	// status := Status{
	// 	ResponseCode:    SuccessCode,
	// 	ResponseMessage: SuccessMessage,
	// }
	httpStatusCode = http.StatusOK
	apiResponse = ResponseSuccess{
		// Status: status,
		Data: data,
	}
	return
}

func NewErrorMessage(err error) (httpStatusCode int, apiResponse Response) {
	var userError utils.AppError
	var status Status
	if errors.As(err, &userError) {
		status = Status{
			ResponseCode:    userError.ErrorCode,
			ResponseMessage: userError.ErrorMessage,
		}
		httpStatusCode = userError.ErrorType
	} else {
		status = Status{
			ResponseCode:    DefaultErrorCode,
			ResponseMessage: DefaultErrorMessage,
		}
		httpStatusCode = http.StatusInternalServerError
	}
	apiResponse = Response{
		Status: status,
		Data:   nil,
	}
	return
}
