package main

import (
	"net/http"
	"time"

	"github.com/JrFarias/go-clean/person"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	person.Router(router)

	router.GET("/healthcheck", func(context *gin.Context) {
		context.JSON(http.StatusOK, time.Now())
	})

	router.Run(":3000")
}
