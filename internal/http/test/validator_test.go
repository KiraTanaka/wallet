package http_test

import (
	"testing"

	"wallet/internal/errors"
	my_http "wallet/internal/http"

	"github.com/stretchr/testify/assert"
)

func TestCheckWallet(t *testing.T) {
	mw := &MockWalletDB{}

	testCases := []struct {
		name     string
		walletId string
		expected errors.HttpError
	}{
		{
			name:     "ok",
			walletId: wallet.Id,
			expected: errors.HttpError{},
		},
		{
			name:     "empty walletId",
			walletId: "",
			expected: errors.GetWalletIdNotPassedError(),
		},
		{
			name:     "invalid type uuid",
			walletId: "b0e82dfa-f37e-4da9-8c38-11",
			expected: errors.GetInvalidWalletIdError(),
		},
		{
			name:     "walletId not exists",
			walletId: "b0e82dfa-1111-1111-8c38-f3fdc9ae8811",
			expected: errors.GetWalletNotFoundError(),
		},
	}

	for _, testCase := range testCases {
		err := my_http.CheckWallet(mw, testCase.walletId)
		assert.Equal(t, testCase.expected, err)
	}
}
