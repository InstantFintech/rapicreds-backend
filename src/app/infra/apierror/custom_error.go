package apierror

import (
	"fmt"
)

const emptyString = ""

type ApiError struct {
	Msg    string
	Status int
	Err    error
}

func NewCustomError(msg string, err error) ApiError {
	errMsg := getErrorMessage(err)
	finalErrMsg := fmt.Sprint(msg + ": " + errMsg)
	return ApiError{
		Msg: finalErrMsg,
		Err: err,
	}
}

func NewCustomErrorWithStatus(msg string, err error, status int) *ApiError {
	errMsg := getErrorMessage(err)
	finalErrMsg := fmt.Sprint(msg + ": " + errMsg)
	return &ApiError{
		Msg:    finalErrMsg,
		Status: status,
		Err:    err,
	}
}

func LogCustomError(err *ApiError) {
	fmt.Print(err.Msg)
}

func getErrorMessage(err error) string {
	if err == nil || err.Error() == emptyString {
		return emptyString
	}
	return err.Error()
}
