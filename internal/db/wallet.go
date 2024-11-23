package db

import (
	_ "embed"

	"github.com/jmoiron/sqlx"
)

type Wallet struct {
	Id      string  `json:"id" db:"id" binding:"max=36"`
	UserId  string  `json:"user_id" db:"user_id" binding:"max=36"`
	Balance float64 `json:"balance" db:"balance" binding:"required"`
}

type WalletDb struct {
	Db *sqlx.DB
}

type WalletModel interface {
	Get(walletId string) (*Wallet, error)
}

//go:embed queries/wallet/checkExists.sql
var checkExistsQuery string

//go:embed queries/wallet/get.sql
var getQuery string

//go:embed queries/wallet/changeBalance.sql
var changeBalanceQuery string

func (w *WalletDb) CheckExists(walletId string) error {
	var walletExists bool
	return w.Db.Get(&walletExists, checkExistsQuery, walletId)
}

func (w *WalletDb) Get(walletId string) (*Wallet, error) {
	wallet := &Wallet{}
	err := w.Db.Get(wallet, getQuery, walletId)
	return wallet, err
}

func (w *WalletDb) ChangeBalance(walletId string, amount float64) error {
	tx, err := w.Db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	_, err = tx.Exec(changeBalanceQuery, walletId, amount)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}
