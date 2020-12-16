package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUp() *gin.Engine {
	engine := gin.Default()
	engine.GET("/hello", func(context *gin.Context) {
		context.String(http.StatusOK, "hello world")
	})
	return engine
}