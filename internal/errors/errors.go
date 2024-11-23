package errors

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

type HttpError struct {
	httpCode int
	Reason   string `json:"reason"`
}

func (e HttpError) SeparateCode() (int, HttpError) {
	return e.httpCode, e
}

func (e HttpError) IsEmpty() bool {
	return e.Reason == ""
}

var (
	//400
	WalletIdNotPassedError    = HttpError{http.StatusBadRequest, "Идентификатор кошелька должен быть указан."}
	InvalidWalletIdError      = HttpError{http.StatusBadRequest, "Недопустимый идентификатор кошелька."}
	InvalidOperationTypeError = HttpError{http.StatusBadRequest, "Недопустимый тип операции"}
	//404
	WalletNotFoundError = HttpError{http.StatusNotFound, "Указанный кошелек не существует."}
)

// 400 (StatusBadRequest)
func GetInvalidRequestFormatOrParametersError(err error) HttpError {
	log.Error(err)
	return HttpError{http.StatusBadRequest, err.Error()}
}

func GetWalletIdNotPassedError() HttpError {
	log.Error(WalletIdNotPassedError)
	return WalletIdNotPassedError
}

func GetInvalidWalletIdError() HttpError {
	log.Error(InvalidWalletIdError)
	return InvalidWalletIdError
}

func GetInvalidOperationTypeError() HttpError {
	log.Error(InvalidOperationTypeError)
	return InvalidOperationTypeError
}

// 404 (StatusNotFound)

func GetWalletNotFoundError() HttpError {
	log.Error(WalletNotFoundError.Reason)
	return WalletNotFoundError
}

// 500 (StatusInternalServerError)

func GetInternalServerError(err error) HttpError {
	log.Error(err)
	return HttpError{http.StatusInternalServerError, err.Error()}
}
