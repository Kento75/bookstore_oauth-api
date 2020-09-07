package app

import (
	"github.com/Kento75/bookstore_oauth-api/src/domain/access_token"
	"github.com/Kento75/bookstore_oauth-api/src/http"
	"github.com/Kento75/bookstore_oauth-api/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atHandler := http.NewHandler(access_token.NewService(db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.Run("127.0.0.1:8080")
}