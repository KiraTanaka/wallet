package db

import (
	_ "embed"

	"github.com/jmoiron/sqlx"
)

type WalletOperation struct {
	Id            string  `json:"id" db:"id" binding:"max=36"`
	WalletId      string  `json:"wallet_id" db:"wallet_id" binding:"required,max=36"`
	OperationType string  `json:"operation_type" db:"operation_type" binding:"required,oneof=DEPOSIT WITHDRAW"`
	Amount        float64 `json:"amount" db:"amount" building:"required"`
}

type WalletOperationDb struct {
	Db *sqlx.DB
}

type WalletOperationModel interface {
	Add(*WalletOperation) error
}

//go:embed queries/walletOperation/add.sql
var addQuery string

func (wo *WalletOperationDb) Add(operation *WalletOperation) error {
	var lastInsertId string
	tx, err := wo.Db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	err = tx.QueryRow(addQuery, operation.WalletId, operation.OperationType, operation.Amount).Scan(&lastInsertId)
	if err != nil {
		return err
	}
	tx.Commit()
	operation.Id = lastInsertId
	return nil
}
