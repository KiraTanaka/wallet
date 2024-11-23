package http_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"wallet/internal/db"
	my_http "wallet/internal/http"

	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
)

type MockWalletDB struct{}
type MockWalletOperationDB struct{}

var wallet db.Wallet = db.Wallet{Id: "b0e82dfa-f37e-4da9-8c38-f3fdc9ae881d",
	UserId:  "0c99719c-4146-4377-9e9b-12663c673173",
	Balance: 10}

func (w *MockWalletDB) CheckExists(walletId string) error {
	if wallet.Id == walletId {
		return nil
	}
	return db.ErrorNoRows
}

func (w *MockWalletDB) Get(walletId string) (*db.Wallet, error) {
	return &wallet, nil
}

func (w *MockWalletDB) ChangeBalance(walletId string, amount float64) error {
	wallet.Balance += amount
	return nil
}

func (wo *MockWalletOperationDB) Add(operation *db.WalletOperation) error {
	return nil
}

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetWallet(t *testing.T) {
	router := SetUpRouter()

	mw := &MockWalletDB{}
	mwo := &MockWalletOperationDB{}

	walletId := "b0e82dfa-f37e-4da9-8c38-f3fdc9ae881d"

	walletHandler := my_http.WalletHandler{Wallet: mw, WalletOperation: mwo}
	router.GET("/:walletId", walletHandler.GetWallet)

	req, _ := http.NewRequest("GET", "/"+walletId, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var wallet db.Wallet
	json.Unmarshal(w.Body.Bytes(), &wallet)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, walletId, wallet.Id)
	assert.NotEmpty(t, wallet)
}

func TestAddWalletOperation(t *testing.T) {
	router := SetUpRouter()

	mw := &MockWalletDB{}
	mwo := &MockWalletOperationDB{}
	testCases := []struct {
		name            string
		walletOperation db.WalletOperation
		expected        float64
	}{
		{
			name: "deposit operation",
			walletOperation: db.WalletOperation{
				WalletId:      "b0e82dfa-f37e-4da9-8c38-f3fdc9ae881d",
				OperationType: "DEPOSIT",
				Amount:        10,
			},
			expected: 20.0,
		},
		{
			name: "withraw operation",
			walletOperation: db.WalletOperation{
				WalletId:      "b0e82dfa-f37e-4da9-8c38-f3fdc9ae881d",
				OperationType: "WITHDRAW",
				Amount:        10,
			},
			expected: 10.0,
		},
	}

	walletHandler := my_http.WalletHandler{Wallet: mw, WalletOperation: mwo}
	router.POST("/", walletHandler.AddWalletOperation)

	for _, testCase := range testCases {

		jsonValue, _ := json.Marshal(testCase.walletOperation)

		req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		var wallet db.Wallet
		json.Unmarshal(w.Body.Bytes(), &wallet)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, testCase.walletOperation.WalletId, wallet.Id)
		assert.Equal(t, testCase.expected, wallet.Balance)
	}
}
