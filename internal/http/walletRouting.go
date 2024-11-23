package http

import (
	"net/http"
	"sync"

	db "wallet/internal/db"
	"wallet/internal/errors"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type WalletHandler struct {
	mutex           sync.Mutex
	Wallet          db.WalletModel
	WalletOperation db.WalletOperationModel
}

func InitWalletRoutes(routes *gin.RouterGroup, walletHandler *WalletHandler) {
	tenderRoutes := routes.Group("/wallet") // для единообразия сделала wallet(в задании отличается эта часть у методов)
	//GET
	tenderRoutes.GET("/:walletId", walletHandler.GetWallet)
	//POST
	tenderRoutes.POST("/", walletHandler.AddWalletOperation)
}

func (w *WalletHandler) GetWallet(c *gin.Context) {
	log.Info("Чтение параметров")

	walletId := c.Param("walletId")

	log.Info("Валидация")
	errHttp := CheckWallet(w.Wallet, walletId)
	if !errHttp.IsEmpty() {
		c.AbortWithStatusJSON(errHttp.SeparateCode())
		return
	}

	log.Info("Чтение данных")

	wallet, err := w.Wallet.Get(walletId)
	if err != nil {
		c.AbortWithStatusJSON(errors.GetInternalServerError(err).SeparateCode())
		return
	}

	c.JSON(http.StatusOK, wallet)
}

func (w *WalletHandler) AddWalletOperation(c *gin.Context) {
	log.Info("Чтение параметров")
	walletOperation := db.WalletOperation{}
	err := c.BindJSON(&walletOperation)
	if err != nil {
		c.AbortWithStatusJSON(errors.GetInvalidRequestFormatOrParametersError(err).SeparateCode())
		return
	}

	log.Info("Валидация")
	errHttp := CheckWallet(w.Wallet, walletOperation.WalletId)
	if !errHttp.IsEmpty() {
		c.AbortWithStatusJSON(errHttp.SeparateCode())
		return
	}

	w.mutex.Lock()

	log.Info("Добавление операции")
	err = w.WalletOperation.Add(&walletOperation)
	if err != nil {
		c.AbortWithStatusJSON(errors.GetInternalServerError(err).SeparateCode())
		return
	}

	amount := walletOperation.Amount

	if walletOperation.OperationType == "WITHDRAW" {
		amount = -amount
	}

	log.Info("Изменение баланса")

	err = w.Wallet.ChangeBalance(walletOperation.WalletId, amount)
	if err != nil {
		c.AbortWithStatusJSON(errors.GetInternalServerError(err).SeparateCode())
		return
	}
	w.mutex.Unlock()

	log.Info("Чтение данных кошелька")

	wallet, err := w.Wallet.Get(walletOperation.WalletId)
	if err != nil {
		c.AbortWithStatusJSON(errors.GetInternalServerError(err).SeparateCode())
		return
	}

	c.JSON(http.StatusOK, wallet)
}
