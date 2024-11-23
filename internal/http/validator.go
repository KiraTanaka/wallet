package http

import (
	"wallet/internal/db"
	errors "wallet/internal/errors"

	"github.com/google/uuid"
)

func CheckWallet(model db.WalletModel, walletId string) errors.HttpError {
	if walletId == "" {
		return errors.GetWalletIdNotPassedError()
	}
	err := uuid.Validate(walletId)
	if err != nil {
		return errors.GetInvalidWalletIdError()
	}
	err = model.CheckExists(walletId)
	if err == db.ErrorNoRows {
		return errors.GetWalletNotFoundError()
	} else if err != nil {
		return errors.GetInternalServerError(err)
	}
	return errors.HttpError{}
}
