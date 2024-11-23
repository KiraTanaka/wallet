package http

import (
	_ "database/sql"
	"net/http"

	db "wallet/internal/db"

	"github.com/gin-gonic/gin"
)

type RouteHandler struct {
	WalletHandler *WalletHandler
	Routes        *gin.Engine
}

func (route *RouteHandler) InitRoutes(dbModels *db.DbModels) {
	route.Routes = gin.Default()
	route.WalletHandler = &WalletHandler{Wallet: dbModels.WalletModel, WalletOperation: dbModels.WalletOperationModel}

	route.Routes.GET("/", hello)
	routeGroup := route.Routes.Group("/api/v1")

	InitWalletRoutes(routeGroup, route.WalletHandler)
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello!")
}

func (route *RouteHandler) Run(serverAddress string) {
	route.Routes.Run(serverAddress)
}
